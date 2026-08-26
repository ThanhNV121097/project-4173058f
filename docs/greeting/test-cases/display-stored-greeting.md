# Test Cases — Display stored greeting

Risk: low. Single read-only page, one stored value, no auth, no writes.

## Case 1
**Scenario**: Show seeded greeting on first open
**Given**: Stored greeting row contains `Hello Word`
**When**: Guest opens page
**Then**: Page shows `Hello Word` and no extra navigation, forms, admin UI, or other visible content
**Check**: render_url

Trace: GREETING-001 AC-1, SRS 4.1.1-4.1.3, service `GET /v1/greeting` success shape.

## Case 2
**Scenario**: Show changed greeting after reload
**Given**: Stored greeting row contains `Bonjour` instead of seed value
**When**: Guest reloads page
**Then**: Page shows `Bonjour` exactly, not `Hello Word`
**Check**: render_url

Trace: GREETING-001 AC-2, SRS 4.1.5.

## Case 3
**Scenario**: Show loading state while read is pending
**Given**: Greeting read request is still pending
**When**: Guest opens page
**Then**: Page shows centred greeting area with `Hello Word` and loading note below it, with no other content visible
**Check**: render_url

Trace: GREETING-001 AC-3, SRS 4.1.1-4.1.2.

## Case 4
**Scenario**: Show error state when greeting cannot be read
**Given**: Greeting read fails or returns no usable greeting
**When**: Guest opens page
**Then**: Page shows centred greeting area with `Hello Word` and error note below it, with no retry control, empty state, or additional content
**Check**: render_url

Trace: GREETING-001 AC-4, SRS 4.1.4, service `GET /v1/greeting` 404/500 failure envelope.
