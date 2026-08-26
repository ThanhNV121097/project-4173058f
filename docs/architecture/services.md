# Services — greeting

## Shared rules
- Public read-only API; no auth.
- Paths omit `/api`; deploy proxy strips that prefix before backend receives request.
- JSON responses use `application/json; charset=utf-8`.

## Error envelope
```json
{
  "error": {
    "code": "internal_error",
    "message": "Could not read greeting."
  }
}
```

| Code | HTTP | Meaning |
|---|---:|---|
| `not_found` | 404 | Greeting row missing |
| `internal_error` | 500 | Database or unexpected server failure |

## Endpoints

### `GET /healthz`
Checks service and database readiness.

Request: none.

Success response: `200 OK`
```json
{"status":"ok"}
```

Failure response: `503 Service Unavailable`
```json
{
  "error": {
    "code": "internal_error",
    "message": "Service unavailable."
  }
}
```

### `GET /v1/greeting`
Returns current stored greeting.

Request: none.

Success response: `200 OK`
```json
{
  "greeting": {
    "text": "Hello Word",
    "updatedAt": "2025-02-14T00:00:00Z"
  }
}
```

Failure responses:
- `404 not_found` when row `id = 1` does not exist.
- `500 internal_error` for database/read failures.
