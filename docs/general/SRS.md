# SRS — general

Module: `general`
Last updated: 2025-02-14
Design: [View the approved design](http://localhost:8080/design/ed9b1c8e-ce15-4dda-a39f-d39483f80f14)
Design system: `design/design-system.md`

## 1. Purpose

This module delivers the single public landing page for hello-word-14. Page shows stored text from PostgreSQL on plain white background.

## 2. Actors

| Actor | Who they are | What they may do in this module |
|---|---|---|
| Guest | Any visitor of public page | View centered message |
| System | Frontend and backend runtime | Read stored message and render it |

## 3. Scope

**In scope**

- Centered Hello Word page

**Out of scope**

- Authentication, user accounts, and permissions.
- Multiple pages, routing, or navigation.
- Editing the message.

## 4. Functional requirements

### 4.1 Centered Hello Word page

**Requirement GENERAL-001 — show stored greeting on single centered page**

Guest sees one centered message screen with text `Hello Word`.

Behaviour:

1. Guest opens page.
2. System reads one greeting value from backend.
3. System renders that value centered horizontally and vertically on plain white background.
4. Frontend does not contain greeting text as static page copy.

**Acceptance criteria** — each maps one-to-one onto a test case in `docs/general/test-cases/centered-hello-word-page.md`.

| # | Given | When | Then |
|---|---|---|---|
| AC-1 | Stored greeting value is `Hello Word` | Guest opens page | Page shows `Hello Word` centered horizontally and vertically |
| AC-2 | Backend returns stored greeting | Guest inspects rendered page source or network response | Greeting text is not hardcoded in frontend page copy |
| AC-3 | Approved design is loaded | Guest opens page | Background is white, text is black, and no extra chrome or animation appears |

**Failure, boundary and permission behaviour**

| Case | Condition | Expected behaviour |
|---|---|---|
| Invalid input | Stored greeting is empty | Empty greeting is not rendered |
| Boundary | Greeting is a short plain string | Value renders unchanged within single centered text line |
| Not found | Greeting row is missing | Page has no loading, empty, or error state in approved design |
| Not permitted | Guest opens page | No permission gate; page remains public |
| Conflict | Greeting row changes during read | Latest committed stored value is shown on refresh |
| Upstream failure | Backend or database is unavailable | Page has no loading, empty, or error state in approved design |

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
| Performance | Page renders stored greeting after backend response returns, with no extra client-side loading state |
| Accessibility | Centered greeting has contrast of at least 21:1 against white background |
| Responsive | Layout remains centered at 320px width and up with no horizontal scroll |
| Localisation | Copy remains exactly `Hello Word` in English |

## 7. Dependencies and assumptions

- Depends on backend API for greeting retrieval.
- Depends on PostgreSQL for stored greeting row.
- One greeting row exists with text `Hello Word`.

## 8. Traceability

| Plan item | Requirement ids | Test cases |
|---|---|---|
| Centered Hello Word page | GENERAL-001 | `test-cases/centered-hello-word-page.md` |
