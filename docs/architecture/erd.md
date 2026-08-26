# ERD — greeting

## Tables

### `greetings`
| Column | Type | Null | Default | Notes |
|---|---|---|---|---|
| `id` | `smallint` | no | none | Primary key; single row uses `1` |
| `text` | `text` | no | none | Exact greeting shown to visitors |
| `created_at` | `timestamptz` | no | `now()` | Row creation time |
| `updated_at` | `timestamptz` | no | `now()` | Last content change time |

Constraints:
- `PRIMARY KEY (id)`
- `CHECK (id = 1)` keeps one product greeting row.
- `CHECK (length(btrim(text)) > 0)` rejects empty greeting.

Seed:
- Migration inserts row `id = 1`, `text = 'Hello Word'`.

## Relationships
None. `greetings` is standalone.

## Migration ownership
Backend applies migrations on boot in filename order and records applied versions in `schema_migrations`.
