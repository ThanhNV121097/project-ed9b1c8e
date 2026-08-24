# SRS — general

Module: `general`
Last updated: 2025-02-14
Design: [View the approved design](http://localhost:8080/design/ed9b1c8e-ce15-4dda-a39f-d39483f80f14)
Design system: `design/design-system.md`

## 1. Purpose

This module delivers single-page hello-word-14 landing page. Guest sees centered "Hello Word" text on plain white screen, loaded from backend from PostgreSQL. If module does not exist, project loses its only visible proof that frontend, backend, and database all connect.

## 2. Actors

| Actor | Who they are | What they may do in this module |
|---|---|---|
| Guest | Any visitor of public page | View centered message |
| System | Frontend and backend runtime | Read stored message and render it |

## 3. Scope

**In scope** — functions specified below, by plan title:

- Centered Hello Word page

**Out of scope**

- Authentication, user accounts, and permissions — not part of hello-word-14.
- Multiple pages, routing, or navigation — design shows one screen only.
- Editing the message — not requested in brief.

## 4. Functional requirements

### 4.1 Centered Hello Word page

**Requirement GENERAL-001 — show stored greeting**

As a Guest, I want to see "Hello Word" loaded from stored data, so that page proves frontend, backend, and database are connected.

Behaviour:

1. Guest opens page.
2. System reads one greeting value from backend.
3. System renders that value centered horizontally and vertically on plain white background.
4. System does not hardcode greeting in frontend.

**Acceptance criteria** — each maps one-to-one onto a test case in `docs/general/test-cases/centered-hello-word-page.md`.

| # | Given | When | Then |
|---|---|---|---|
| AC-1 | Stored greeting value is `Hello Word` | Guest opens page | Page shows `Hello Word` centered horizontally and vertically |
| AC-2 | Backend returns stored greeting | Guest inspects frontend source or rendered output | Greeting text is not hardcoded in frontend |
| AC-3 | Approved design is loaded | Guest opens page | Background is white, text is black, and no extra chrome or animation appears |

**Failure, boundary and permission behaviour**

| Case | Condition | Expected behaviour |
|---|---|---|
| Invalid input | Stored greeting is empty | Empty greeting is not rendered; exact handling is specified by backend contract |
| Boundary | Greeting is a short plain string | Value renders unchanged within single centered text line |
| Not found | Greeting row is missing | No error or empty state is part of approved design; API contract's error envelope is specified in service contract |
| Not permitted | Guest opens page | No permission gate; page remains public |
| Conflict | Greeting row changes during read | Latest committed stored value is shown on refresh |
| Upstream failure | Backend or database is unavailable | No failure screen is part of approved design; API contract's error envelope is specified in service contract |

**Data touched**

| Field | Type | Required | Rule |
|---|---|---|---|
| greeting_text | text | yes | One plain string, shown exactly as stored |

## 5. Screens

| Screen | Section in the design | Functions it serves | States that must exist |
|---|---|---|---|
| Single landing page | Approved design preview | GENERAL-001 | default |

## 6. Non-functional requirements

| Area | Requirement |
|---|---|
| Performance | Page shows stored greeting within 1 second on typical connection after API response returns |
| Accessibility | Centered greeting has contrast of at least 21:1 against white background |
| Responsive | Layout remains centered at 320px width and up with no horizontal scroll |
| Localisation | Copy remains exactly `Hello Word` in English |

## 7. Dependencies and assumptions

- **Depends on:** backend API, for greeting retrieval.
- **Depends on:** PostgreSQL, for stored greeting row.
- **Assumption:** one greeting row exists; if absent, backend contract decides error response.

| Open question | Proposed default | Who decides |
|---|---|---|
| Should missing greeting row return empty body or error? | Error response with no UI error state | Stakeholder / TL |

## 8. Traceability

| Plan item | Requirement ids | Test cases |
|---|---|---|
| Centered Hello Word page | GENERAL-001 | `test-cases/centered-hello-word-page.md` |
