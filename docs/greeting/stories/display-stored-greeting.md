# Story — Display stored greeting

## User story
As a Guest, I want to see stored greeting text on page load, so that I can read current product content.

## In scope
- Load greeting from product API when page opens.
- Show centred greeting frame in loaded, loading, and error states.
- Render exact stored greeting text when read succeeds.
- Render `Hello Word` seed text while loading or when read fails or returns no usable greeting.
- Reloading page reflects changed stored greeting value from system of record.

## Out of scope
- Navigation.
- Forms.
- Admin UI.
- Extra styling beyond approved centred greeting layout.
- Write flows, edit flows, or conflict handling.
- Storage layout, API contract, and service boundaries; these live in architecture docs.

## UI scope
This story covers `Greeting page` in the approved design, across default, loading, and error states only.
- Loaded: centred greeting on white background.
- Loading: centred greeting with loading note.
- Error: centred greeting with error note.

## Acceptance criteria
- Given stored greeting row contains `Hello Word`, when Guest opens page, then page shows `Hello Word`.
- Given stored greeting row contains a different value, when Guest reloads page, then page shows changed value.
- Given greeting read is pending, when Guest opens page, then page shows loading state with `Hello Word` and loading note.
- Given greeting read fails or returns no usable greeting, when Guest opens page, then page shows error state with `Hello Word` and error note.

## Dependencies
- PostgreSQL with stored greeting row.
- HTTP API service contract for reading current greeting text.
- Seed value set to exactly `Hello Word`.
