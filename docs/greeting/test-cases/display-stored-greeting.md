# Test Cases — Display stored greeting

Risk level: low. Single read-only page, one stored value, one API-backed loading/error path.

## Coverage map
- GREETING-001 / AC-1: seed value shown on open
- GREETING-001 / AC-2: changed stored value shown after reload
- GREETING-001 / AC-3: loading state shows seed text and loading note
- GREETING-001 / AC-4: read failure or unusable greeting shows seed text and error note
- Services contract: `GET /v1/greeting` success, `404 not_found`, `500 internal_error`

## Scenarios

**Scenario**: Seed greeting shown on first page load
**Given** stored greeting row contains `Hello Word`
**When** Guest opens page
**Then** page shows `Hello Word` as visible greeting text
Check: render_url

**Scenario**: Reload shows changed stored greeting
**Given** stored greeting row contains `Hello Word` and system of record updates it to a different value before reload
**When** Guest reloads page
**Then** page shows changed value exactly as stored
Check: render_url

**Scenario**: Loading state shows seed text and loading note
**Given** greeting read is pending
**When** Guest opens page
**Then** page shows centred greeting area with `Hello Word` and loading note below it, and no other content is visible
Check: render_url

**Scenario**: Read failure shows seed text and error note
**Given** greeting read fails or returns no usable greeting
**When** Guest opens page
**Then** page shows centred greeting area with `Hello Word` and error note below it, and no retry control, empty state, or additional content is shown
Check: render_url

**Scenario**: GET /v1/greeting returns stored greeting payload
**Given** stored greeting row exists
**When** client sends `GET /v1/greeting`
**Then** response is `200 OK` with `application/json; charset=utf-8` and body containing `greeting.text` and `greeting.updatedAt`
Check: fetch_url

**Scenario**: GET /v1/greeting missing row returns not_found
**Given** greeting row with `id = 1` does not exist
**When** client sends `GET /v1/greeting`
**Then** response is `404` with error code `not_found`
Check: fetch_url

**Scenario**: GET /v1/greeting read failure returns internal_error envelope
**Given** database or read path fails
**When** client sends `GET /v1/greeting`
**Then** response is `500` with error envelope code `internal_error` and message `Could not read greeting.`
Check: fetch_url

**Scenario**: Guest access has no permission gate
**Given** any visitor is not signed in
**When** Guest opens page
**Then** page is still shown; no sign-in or role-based denial appears
Check: render_url

**Scenario**: Greeting text is not truncated at normal product length
**Given** stored greeting row contains normal product text length
**When** Guest opens page
**Then** page shows exact stored greeting text with no truncation
Check: render_url
