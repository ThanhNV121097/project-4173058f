# Architecture overview — hello-word-28

## Scope
Fullstack slice: PostgreSQL stores one greeting row, Go backend exposes read API, Next.js frontend renders centred greeting. No auth, navigation, forms, admin UI, or extra styling.

## Stack
| Part | Choice | Reason | Rejected alternative |
|---|---|---|---|
| Frontend | Next.js 15 App Router, TypeScript, Tailwind v3 | Matches repo Dockerfile and SSR-capable UI baseline | Static HTML rejected because greeting must come from API |
| Backend | Go 1.25 module under `code/backend` | Matches committed Dockerfile and small HTTP service need | Node API rejected to avoid second runtime stack |
| Database | PostgreSQL 16 | Required system of record for greeting row | In-memory seed rejected because reload must reflect stored value |
| CI | Existing `.github/workflows/ci.yml` | Runs build, vet, tests, lint, token checks on PR | New workflow rejected because `.github/` is fixed/read-only |

## Folder structure
```text
code/backend/
  cmd/api/main.go            # one main package, HTTP server
  internal/migrations/       # embed + apply SQL files
  migrations/*.sql           # ordered up/down migrations
  .env.example               # runtime keys, no secrets
code/frontend/
  app/layout.tsx             # App Router root layout
  app/page.tsx               # composition root only
  app/globals.css            # shared design tokens and base styles
  components/                # story-owned components later
  lib/mock/                  # UI story mocks until API lands
docs/architecture/
  overview.md
  erd.md
  services.md
```

## Runtime flow
1. `docker compose --profile local up` starts PostgreSQL, backend, frontend.
2. Backend reads `DATABASE_URL`, applies pending migrations, checks DB, then listens on `PORT`, `APP_PORT`, or `8080`.
3. `/healthz` returns 200 only after migrations succeeded and `SELECT 1` works.
4. Frontend reads API base from `API_ORIGIN` on server and `NEXT_PUBLIC_API_URL` in browser code.

## Naming and conventions
- Backend route paths use `/v1/...`; never `/api/...` because deploy proxy strips `/api` before Go sees request.
- Database tables and columns use `snake_case`.
- Go packages use short lowercase names; one executable lives at `cmd/api`.
- React components use `export default function ComponentName()`.
- `app/page.tsx` stays server component and only composes story components.
- Shared CSS values must be tokens in `app/globals.css`; CSS modules may not hardcode colors or spacing.

## Environment variables
| Service | Key | Purpose |
|---|---|---|
| backend | `DATABASE_URL` | PostgreSQL connection string injected by runtime/compose |
| backend | `PORT` | HTTP listen port, preferred |
| backend | `APP_PORT` | Legacy/fallback listen port |
| frontend | `NEXT_PUBLIC_API_URL` | Browser-visible backend origin |
| frontend | `API_ORIGIN` | Server-side backend origin inside container network |
| root compose | `POSTGRES_USER`, `POSTGRES_PASSWORD`, `POSTGRES_DB` | Local database bootstrap values |
| root compose | `BACKEND_PORT`, `FRONTEND_PORT` | Optional host port overrides |

## How to run
```bash
cp .env.example .env
cp code/backend/.env.example code/backend/.env
cp code/frontend/.env.example code/frontend/.env.local
docker compose --profile local up --build
```

## Local checks
```bash
cd code/backend && go build ./... && go vet ./... && go test ./...
cd code/frontend && npm ci && npm run lint && npm run build && npm test --if-present
```

## Decisions and tradeoffs
| Decision | Tradeoff |
|---|---|
| Self-migrate on backend boot | Startup owns schema readiness; bad migration blocks health instead of failing first request |
| Keep seed in SQL migration | DB is source of truth; changing seed later needs migration instead of code edit |
| Use stdlib HTTP router | Less code and no dependency; upgrade only if route complexity grows |
| Use `pgx` driver only external backend dependency | Needed for PostgreSQL; ORM rejected as overkill for one read model |
| Minimal frontend shell | Story can mount component with one import and one element; scaffold does not implement product UI |

## Risks and unknowns
- Operator update path is outside scope; direct DB/system-of-record change assumed.
- Greeting length has no product limit beyond PostgreSQL `text`; UI must preserve exact value.
- No auth by design; API is public read-only.
