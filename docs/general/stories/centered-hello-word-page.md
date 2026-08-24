# Story — Centered Hello Word page

As a guest, I want to see a single centered Hello Word page, so that I know the app is working end to end.

## In scope

- Single public landing page for hello-word-14.
- Frontend reads greeting value from backend API.
- Backend reads greeting from PostgreSQL row.
- Page renders `Hello Word` centered horizontally and vertically on plain white background.
- Text is shown exactly as stored and is not hardcoded in frontend page copy.

## Out of scope

- Authentication, accounts, or permissions.
- Multiple pages, routing, or navigation.
- Editing or updating greeting text.
- Loading, empty, or error UI states beyond approved static design.
- Animation, extra chrome, or alternate palettes.

## UI scope

- One screen only: the approved single landing page in `design/index.html`.
- Default state only.
- No interactive controls, no keyboard interaction, no client-side motion.
- Visual result: white background, black centered message, no extra elements.

## Acceptance criteria

1. Guest opens page and sees `Hello Word` centered horizontally and vertically.
2. Page content comes from backend response backed by PostgreSQL, not hardcoded frontend page copy.
3. Rendered page matches approved design: white background, black text, no extra chrome, no animation.
4. Guest can load page without login or other permission gate.
5. Missing, empty, or failing backend data does not introduce an alternate UI state in this story slice.

## Dependencies

- Approved design and design system for single centered message.
- `docs/general/SRS.md` GENERAL-001.
- Backend API for greeting retrieval.
- PostgreSQL row seeded with `Hello Word`.
