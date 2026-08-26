# Story — Display stored greeting

## User story
As a Guest, I want to see the stored greeting text on page load, so that I can read current product content.

## In scope
- Show the stored greeting as single centred line on page load.
- Preserve exact stored text, including seed value `Hello Word`.
- Reload page after stored value changes and show new value.
- Match approved design for the loaded/default state only: white background, centred greeting frame, black greeting text.

## Out of scope
- Loading state, loading note, spinner, or retry control.
- Error state, empty state, or any upstream-failure screen.
- Navigation, forms, admin UI, and extra styling.
- Storage layout, API contract, and service boundaries.

## UI scope
- Greeting page, default state only.
- One centred greeting frame on white background, per approved design system.
- No separate loading or error screens are part of this story.

## Acceptance criteria
1. Given stored greeting row contains `Hello Word`, when Guest opens page, then page shows `Hello Word`.
2. Given stored greeting row contains a different value, when Guest reloads page, then page shows the changed value.
3. Given greeting value exists in stored row, when page loads, then only the centred greeting content is visible.
4. Given page is rendered, when Guest views it, then no navigation, form, admin, loading, or error content is visible.

## Dependencies
- PostgreSQL stores the greeting row.
- HTTP API service contract reads current greeting text.
- Seed value starts as exactly `Hello Word`.
