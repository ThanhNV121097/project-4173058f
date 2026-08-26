# Test Cases — Display stored greeting

Risk: low. One read-only page, no auth, no user input. Cover exact display, reload behavior, loading state, error/no-greeting state, and service contract shape.

## Scenarios

**Scenario**: Show seed greeting on first open
**Given**: Stored greeting row contains `Hello Word`
**When**: Guest opens page
**Then**: Page shows `Hello Word` as visible greeting text
**Check**: render_url
**Trace**: GREETING-001 / AC-1

**Scenario**: Show changed greeting after reload
**Given**: Stored greeting row contains a different value
**When**: Guest reloads page
**Then**: Page shows the changed value exactly as stored
**Check**: render_url
**Trace**: GREETING-001 / AC-2

**Scenario**: Show loading state while greeting read pending
**Given**: Greeting read request is still pending
**When**: Guest opens page
**Then**: Page shows centred greeting area with `Hello Word` and loading note below it; no other content is visible
**Check**: render_url
**Trace**: GREETING-001 / Behaviour 1-2 / AC-3

**Scenario**: Show error state when greeting read fails or returns no usable greeting
**Given**: Greeting read fails or returns no usable greeting
**When**: Guest opens page
**Then**: Page shows centred greeting area with `Hello Word` and error note below it; no retry control, empty state, or additional content is shown
**Check**: render_url
**Trace**: GREETING-001 / Behaviour 4 / AC-4

**Scenario**: Return current greeting from API
**Given**: Stored greeting row contains a greeting value
**When**: Client calls `GET /v1/greeting`
**Then**: Response is `200 OK` with JSON body containing `greeting.text` equal to stored value and `greeting.updatedAt`
**Check**: fetch_url
**Trace**: Services — greeting / GET /v1/greeting success

**Scenario**: Return not_found when greeting row missing
**Given**: Greeting row `id = 1` does not exist
**When**: Client calls `GET /v1/greeting`
**Then**: Response is `404` with error code `not_found`
**Check**: fetch_url
**Trace**: Services — greeting / GET /v1/greeting failure

**Scenario**: Return internal_error on greeting read failure
**Given**: Database or read path fails
**When**: Client calls `GET /v1/greeting`
**Then**: Response is `500` with error code `internal_error`
**Check**: fetch_url
**Trace**: Services — greeting / GET /v1/greeting failure

## Notes

- No permission cases: module has no sign-in or role-based access.
- No manual cases: all expectations are observable by HTTP response or rendered page text.
