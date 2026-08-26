# Test Cases — Display stored greeting

Risk level: low. One static page, one API read, no auth, no writes. Cases focus on exact text, reload behavior, and no extra UI.

## Scenario 1: Seed greeting shows on first open
**Given** stored greeting row `id = 1` contains exactly `Hello Word`
**When** Guest opens page
**Then** page displays exactly `Hello Word` as only visible product text in centered greeting area, with no navigation, forms, admin UI, loading note, error note, retry control, spinner, or fallback content
**Check:** render_url
**Trace:** GREETING-001 AC-1; seed assumption; default one-state screen

## Scenario 2: Changed stored greeting shows after reload
**Given** stored greeting row `id = 1` contains `Good morning`
**When** Guest reloads page after system of record save
**Then** page displays exactly `Good morning` and not previous value
**Check:** render_url
**Trace:** GREETING-001 AC-2; reload reflects stored value

## Scenario 3: Page has no loading or error state
**Given** page opens in normal runtime
**When** Guest opens page
**Then** page renders only default state: centered greeting area on white background with fetched greeting text, and no loading note, error note, spinner, retry control, empty state, or upstream-failure screen
**Check:** render_url
**Trace:** revised scope; single-state default screen

## Scenario 4: Greeting read endpoint returns expected success shape
**Given** backend has stored greeting row `id = 1` with `Hello Word`
**When** client requests `GET /v1/greeting`
**Then** response is `200 OK`, `Content-Type: application/json; charset=utf-8`, and body contains `{"greeting":{"text":"Hello Word","updatedAt":...}}`
**Check:** fetch_url
**Trace:** services contract; display source data

## Scenario 5: Greeting read failure is not masked by UI state
**Given** greeting read cannot complete because row `id = 1` is missing or backing store read fails
**When** Guest opens page
**Then** no alternate page state appears; failure is treated as defect and checked through API/logs, not by loading or error UI
**Check:** manual
**Trace:** revised scope; failure is defect, not screen
