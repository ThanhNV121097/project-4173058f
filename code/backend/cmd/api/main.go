package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/fs"
	"log"
	"net/http"
	"os"
	"sort"
	"strings"
	"time"

	"github.com/ThanhNV121097/project-4173058f/backend/migrations"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type server struct {
	db    *pgxpool.Pool
	ready bool
}

type errorEnvelope struct {
	Error apiError `json:"error"`
}

type apiError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

type greetingResponse struct {
	Greeting greetingDTO `json:"greeting"`
}

type greetingDTO struct {
	Text      string `json:"text"`
	UpdatedAt string `json:"updatedAt"`
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	databaseURL := os.Getenv("DATABASE_URL")
	if strings.TrimSpace(databaseURL) == "" {
		log.Fatal("DATABASE_URL is required")
	}

	db, err := pgxpool.New(ctx, databaseURL)
	if err != nil {
		log.Fatalf("connect database: %v", err)
	}
	defer db.Close()

	if err := migrate(ctx, db); err != nil {
		log.Fatalf("apply migrations: %v", err)
	}

	if err := db.Ping(ctx); err != nil {
		log.Fatalf("ping database: %v", err)
	}

	app := &server{db: db, ready: true}
	mux := http.NewServeMux()
	mux.HandleFunc("GET /healthz", app.healthz)
	mux.HandleFunc("GET /v1/greeting", app.greeting)

	addr := ":" + listenPort()
	log.Printf("listening on %s", addr)
	if err := http.ListenAndServe(addr, mux); err != nil {
		log.Fatal(err)
	}
}

func listenPort() string {
	if port := strings.TrimSpace(os.Getenv("PORT")); port != "" {
		return port
	}
	if port := strings.TrimSpace(os.Getenv("APP_PORT")); port != "" {
		return port
	}
	return "8080"
}

func (s *server) healthz(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 3*time.Second)
	defer cancel()

	if !s.ready || s.db.Ping(ctx) != nil {
		writeError(w, http.StatusServiceUnavailable, "internal_error", "Service unavailable.")
		return
	}
	writeJSON(w, http.StatusOK, map[string]string{"status": "ok"})
}

func (s *server) greeting(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 3*time.Second)
	defer cancel()

	var text string
	var updatedAt time.Time
	err := s.db.QueryRow(ctx, `SELECT text, updated_at FROM greetings WHERE id = 1`).Scan(&text, &updatedAt)
	if errors.Is(err, pgx.ErrNoRows) {
		writeError(w, http.StatusNotFound, "not_found", "Greeting not found.")
		return
	}
	if err != nil {
		writeError(w, http.StatusInternalServerError, "internal_error", "Could not read greeting.")
		return
	}
	if strings.TrimSpace(text) == "" {
		writeError(w, http.StatusNotFound, "not_found", "Greeting not found.")
		return
	}

	writeJSON(w, http.StatusOK, greetingResponse{Greeting: greetingDTO{
		Text:      text,
		UpdatedAt: updatedAt.UTC().Format(time.RFC3339),
	}})
}

func writeJSON(w http.ResponseWriter, status int, body any) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(body); err != nil {
		log.Printf("write response: %v", err)
	}
}

func writeError(w http.ResponseWriter, status int, code string, message string) {
	writeJSON(w, status, errorEnvelope{Error: apiError{Code: code, Message: message}})
}

func migrate(ctx context.Context, db *pgxpool.Pool) error {
	_, err := db.Exec(ctx, `CREATE TABLE IF NOT EXISTS schema_migrations (
		version text PRIMARY KEY,
		applied_at timestamptz NOT NULL DEFAULT now()
	)`)
	if err != nil {
		return fmt.Errorf("create schema_migrations: %w", err)
	}

	entries, err := fs.ReadDir(migrations.Files, ".")
	if err != nil {
		return fmt.Errorf("read migrations: %w", err)
	}

	var names []string
	for _, entry := range entries {
		name := entry.Name()
		if !entry.IsDir() && strings.HasSuffix(name, ".up.sql") {
			names = append(names, name)
		}
	}
	sort.Strings(names)

	for _, name := range names {
		version := strings.TrimSuffix(name, ".up.sql")
		if err := applyMigration(ctx, db, version, name); err != nil {
			return err
		}
	}
	return nil
}

func applyMigration(ctx context.Context, db *pgxpool.Pool, version string, name string) error {
	tx, err := db.Begin(ctx)
	if err != nil {
		return fmt.Errorf("begin migration %s: %w", version, err)
	}
	defer func() {
		if err := tx.Rollback(ctx); err != nil && !errors.Is(err, pgx.ErrTxClosed) {
			log.Printf("rollback migration %s: %v", version, err)
		}
	}()

	var exists bool
	if err := tx.QueryRow(ctx, `SELECT EXISTS (SELECT 1 FROM schema_migrations WHERE version = $1)`, version).Scan(&exists); err != nil {
		return fmt.Errorf("check migration %s: %w", version, err)
	}
	if exists {
		return tx.Commit(ctx)
	}

	sqlBytes, err := migrations.Files.ReadFile(name)
	if err != nil {
		return fmt.Errorf("read migration %s: %w", name, err)
	}
	if _, err := tx.Exec(ctx, string(sqlBytes)); err != nil {
		return fmt.Errorf("run migration %s: %w", version, err)
	}
	if _, err := tx.Exec(ctx, `INSERT INTO schema_migrations (version) VALUES ($1)`, version); err != nil {
		return fmt.Errorf("record migration %s: %w", version, err)
	}
	return tx.Commit(ctx)
}
