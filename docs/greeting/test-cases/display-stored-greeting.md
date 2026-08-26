# Test cases — Display stored greeting

Risk level: low. Single read-only page, but covers startup loading, API failure, seeded content, reload-after-update, and service contract shapes.

## Cases

**Scenario**: Seeded greeting shows on page load
**Given** Stored greeting row contains `Hello Word`
**When** Guest opens page
**Then** Page shows `Hello Word` as visible greeting content and no extra navigation, forms, or admin UI
Check: render_url
Trace: GREETING-001 AC-1

**Scenario**: Updated greeting shows after reload
**Given** Stored greeting row contains a different value
**When** Guest reloads page
**Then** Page shows changed value exactly as stored
Check: render_url
Trace: GREETING-001 AC-2

**Scenario**: Loading state shows while greeting read is pending
**Given** Greeting read request is in flight and response has not arrived yet
**When** Guest opens page
**Then** Page shows centred greeting area with `Hello Word` and loading note below it; no other content is visible
Check: render_url
Trace: GREETING-001 AC-3

**Scenario**: Error state shows when greeting read fails
**Given** Greeting read request fails or returns no usable greeting
**When** Guest opens page
**Then** Page shows centred greeting area with `Hello Word` and error note below it; no retry control, empty state, or additional content is shown
Check: render_url
Trace: GREETING-001 AC-4

**Scenario**: Service returns greeting payload shape
**Given** Stored greeting row exists
**When** Client fetches `GET /v1/greeting`
**Then** Response is `200 OK` with `application/json; charset=utf-8` and body object containing `greeting.text` and `greeting.updatedAt`
Check: fetch_url
Trace: services.md GET /v1/greeting success response

**Scenario**: Missing greeting row returns not_found
**Given** Greeting row `id = 1` does not exist
**When** Client fetches `GET /v1/greeting`
**Then** Response is `404` with error code `not_found`
Check: fetch_url
Trace: services.md GET /v1/greeting failure response

**Scenario**: Read failure returns internal_error envelope
**Given** Database or read path fails
**When** Client fetches `GET /v1/greeting`
**Then** Response is `500` with error code `internal_error` and message `Could not read greeting.`
Check: fetch_url
Trace: services.md error envelope and GET /v1/greeting failure response

**Scenario**: Public read-only access requires no sign-in
**Given** Any visitor with no sign-in
**When** Guest opens page
**Then** Page is viewable without auth prompt or permission gate
Check: manual
Trace: GREETING-001 Permission

**Scenario**: Centre alignment and colors match approved design
**Given** Page is loaded
**When** Guest opens page
**Then** Main greeting area is centred on white background with black text and no extra styling beyond centring
Check: measure_styles
Trace: GREETING-001 screens and design spec

**Scenario**: No horizontal scroll at 320px width
**Given** Browser viewport is 320px wide
**When** Guest opens page
**Then** Page shows one centred line without horizontal scroll
Check: measure_styles
Trace: GREETING-001 Responsive

**Scenario**: Copy is English only
**Given** Page is loaded
**When** Guest opens page
**Then** Visible copy uses English text only and no locale-specific formatting appears
Check: render_url
Trace: GREETING-001 Localisation
