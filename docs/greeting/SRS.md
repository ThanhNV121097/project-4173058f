# SRS — greeting

Module: `greeting`
Last updated: 2025-02-14
Design: [View the approved design](http://localhost:8080/design/4173058f-232f-4eae-9be2-8521d91eb2f7)
Design system: `design/design-system.md`

> One file per module, at `docs/{module}/SRS.md`. It covers only the functions
> that belong to this module. Never write `docs/SRS.md`.

## 1. Purpose

`greeting` shows the product's single stored greeting to any visitor. If this module does not exist, the app becomes a blank shell with no visible product value.

## 2. Actors

| Actor | Who they are | What they may do in this module |
|---|---|---|
| Guest | Any visitor with no sign-in requirement | View stored greeting |
| Operator | Team member who updates stored content outside this module | Change stored greeting data through the system of record, then reload the page to verify the new value appears |

## 3. Scope

**In scope** — the functions specified below, by their plan titles:

- Display stored greeting

**Out of scope** — name what a reader would reasonably expect here and say
where it lives instead. This section prevents the same argument twice.

- Navigation, forms, admin UI, and extra styling — deliberately not built; scope excludes them.
- Storage layout, API contract, and service boundaries — belong to `architecture` docs.

## 4. Functional requirements

### 4.1 Display stored greeting

**Requirement GREETING-001 — show current greeting**

*As a* Guest, *I want to* see the stored greeting text on page load, *so that* I can read current product content.

Behaviour:

1. When the page loads, the app reads the stored greeting through the product API.
2. When the API returns a greeting value, the page shows that exact text as the only visible content in the centred greeting area.
3. When the stored greeting value changes in the system of record and the page is reloaded, the new value is shown.

**Acceptance criteria** — each maps one-to-one onto a test case in `docs/greeting/test-cases/display-stored-greeting.md`. Given/When/Then, no compound conditions: one behaviour per criterion.

| # | Given | When | Then |
|---|---|---|---|
| AC-1 | Stored greeting row contains `Hello Word` | Guest opens page | Page shows `Hello Word` |
| AC-2 | Stored greeting row contains a different value | Guest reloads page | Page shows the changed value |

**Failure, boundary and permission behaviour**

| Case | Condition | Expected behaviour |
|---|---|---|
| Permission | Any visitor opens page | Not applicable: no sign-in or role-based permission exists in this module |
| Boundary | Stored greeting length is normal product text length | Accepted and shown exactly as stored, with no truncation defined by this module |
| Upstream failure | Greeting API or backing store is unavailable | Not applicable: no error or empty state is part of the approved design; the API contract's error envelope is specified in the service contract |
| Conflict | Two actors update greeting at same time | Not applicable: this module only displays the stored value and does not resolve writes |

**Data touched** — the fields this function reads and writes, in product terms.

| Field | Type | Required | Rule |
|---|---|---|---|
| Greeting text | text | yes | Must render exact stored value, including `Hello Word` seed |

## 5. Screens

The design is the source of truth for appearance; this section maps functions onto it so nothing in the design is unaccounted for and nothing specified here is missing from the design.

| Screen | Section in the design | Functions it serves | States that must exist |
|---|---|---|---|
| Greeting page | Loaded | GREETING-001 | default |
| Greeting page | Loading | GREETING-001 | default |
| Greeting page | Error | GREETING-001 | default |

## 6. Non-functional requirements

| Area | Requirement |
|---|---|
| Performance | Greeting page responds within 2s at p95 on 1 Mbps connection with cold browser cache |
| Accessibility | Greeting text remains keyboard reachable only as page content, with visible focus not required because no interactive controls exist; contrast between text and background is at least 4.5:1 |
| Responsive | Page shows one centred line without horizontal scroll at 320px width and up |
| Localisation | Copy is in English; no locale-specific formatting is used |
| Privacy | No personal data is stored or displayed by this module |

## 7. Dependencies and assumptions

- **Depends on:** PostgreSQL, for the stored greeting row.
- **Depends on:** HTTP API service contract, for reading current greeting text.
- **Assumption:** Seed value is exactly `Hello Word`; if it changes, AC-1 and the seed-related test case change too.

| Open question | Proposed default | Who decides |
|---|---|---|
| None | Not applicable | Stakeholder |

## 8. Traceability

Every plan item in this module appears exactly once, and every requirement id traces to a test case. A gap in this table is a gap in the build.

| Plan item | Requirement ids | Test cases |
|---|---|---|
| Display stored greeting | GREETING-001 | `docs/greeting/test-cases/display-stored-greeting.md` |
