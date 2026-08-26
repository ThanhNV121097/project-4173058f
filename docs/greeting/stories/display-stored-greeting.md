# Story plan — Display stored greeting

As a Guest, I want to see stored greeting text on page load, so that I can read current product content.

## In scope
- One greeting page that reads current stored greeting through product API and shows returned text centered on white background.
- Default state only; no loading note, spinner, retry control, fallback note, or separate error screen.
- Page reload reflects updated stored value from system of record.
- Exact greeting text is shown with no extra navigation, forms, admin UI, or extra styling.

## Out of scope
- Navigation, forms, admin UI, and extra styling; those stay out of module scope.
- Storage layout, API contract, and service boundaries; those belong in architecture docs.
- Loading, error, retry, spinner, and any upstream-failure UI; read failure is treated as defect, not screen.

## UI scope
- Greeting page, default state only, matching approved centered single-line design.
- No additional screens or states beyond the single rendered greeting view.

## Acceptance criteria
1. Given stored greeting row contains `Hello Word`, when Guest opens page, then page shows `Hello Word`.
2. Given stored greeting row contains a different value, when Guest reloads page, then page shows changed value.
3. Given greeting is readable from store, when Guest opens page, then page shows centered greeting only.
4. Given greeting value changes in system of record, when Guest reloads page, then new value appears with same centered layout.

## Dependencies
- PostgreSQL stores greeting row.
- HTTP API contract for reading current greeting text.
- Seed value is exactly `Hello Word`.
- Frontend design and architecture constraints already approved.
