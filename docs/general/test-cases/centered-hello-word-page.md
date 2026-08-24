# Test Cases — Centered Hello Word page

Risk level: low. Single public page, one read-only API, one stored row. Cases cover render, contract, failure, and visual criteria.

## Scenario 1: Stored greeting renders centered on page
**Given** PostgreSQL has one greeting row with text `Hello Word`, backend API returns that stored value, and page is opened as guest
**When** guest loads page
**Then** browser displays only `Hello Word` centered horizontally and vertically on plain white background, with black text, no extra chrome, and no animation
**Requirement** GENERAL-001, AC-1, AC-3, screen default, NFR Accessibility, NFR Responsive
**Check**: render_url; measure_styles

## Scenario 2: Frontend does not hardcode greeting text
**Given** backend API returns stored greeting and frontend bundle/source is available for inspection
**When** guest inspects rendered page source or network response
**Then** greeting text appears only from backend response, not as static frontend page copy
**Requirement** GENERAL-001, AC-2
**Check**: fetch_url; render_url

## Scenario 3: Public guest access has no permission gate
**Given** public page is open as guest and no authentication is configured for module
**When** guest loads page
**Then** page is accessible without login or other permission prompt
**Requirement** GENERAL-001, Not permitted
**Check**: render_url

## Scenario 4: Empty stored greeting does not render empty text
**Given** stored greeting row exists with empty text and backend returns that value
**When** guest loads page
**Then** empty greeting is not rendered as visible page text
**Requirement** GENERAL-001, Invalid input
**Check**: render_url

## Scenario 5: Short plain greeting stays unchanged on one line
**Given** stored greeting row exists with short plain string `Hello Word` and backend returns it
**When** guest loads page at 320px width
**Then** visible text matches stored value exactly and stays on single centered text line without horizontal scroll
**Requirement** GENERAL-001, Boundary, NFR Responsive, NFR Localization
**Check**: render_url; measure_styles

## Scenario 6: Missing greeting row yields no approved-design error state
**Given** greeting row is missing and backend returns 404 `not_found` with message `Greeting not found`
**When** guest loads page
**Then** page shows no loading, empty, or error state in approved design
**Requirement** GENERAL-001, Not found
**Check**: fetch_url; render_url

## Scenario 7: Backend or database failure yields no approved-design error state
**Given** backend or database is unavailable and request fails with 500 `internal_error` and message `Request failed`
**When** guest loads page
**Then** page shows no loading, empty, or error state in approved design
**Requirement** GENERAL-001, Upstream failure
**Check**: fetch_url; render_url

## Scenario 8: Service contract success shape for greeting API
**Given** backend has stored greeting row with text `Hello Word`
**When** guest sends `GET /v1/greeting`
**Then** response is HTTP 200 with JSON body `{ "text": "Hello Word" }`
**Requirement** services.md Read greeting success response
**Check**: fetch_url

## Scenario 9: Service contract 404 for missing greeting row
**Given** greeting row is missing
**When** guest sends `GET /v1/greeting`
**Then** response is HTTP 404 with error envelope `{ "error": { "code": "not_found", "message": "Greeting not found" } }`
**Requirement** services.md Read greeting errors
**Check**: fetch_url

## Scenario 10: Service contract 500 for database or unexpected failure
**Given** database or backend fails while handling greeting request
**When** guest sends `GET /v1/greeting`
**Then** response is HTTP 500 with error envelope `{ "error": { "code": "internal_error", "message": "Request failed" } }`
**Requirement** services.md Error envelope, common errors
**Check**: fetch_url

## Scenario 11: Health check returns ok only when PostgreSQL works
**Given** migrations succeeded and `SELECT 1` against PostgreSQL succeeds
**When** guest sends `GET /healthz`
**Then** response is HTTP 200 with plain text body `ok`
**Requirement** services.md Health check success
**Check**: fetch_url

## Scenario 12: Health check fails with service unavailable when runtime or DB is not ready
**Given** migrations not succeeded or `SELECT 1` against PostgreSQL fails
**When** guest sends `GET /healthz`
**Then** response is HTTP 503 with error envelope `{ "error": { "code": "service_unavailable", "message": "Service unavailable" } }`
**Requirement** services.md Health check failure
**Check**: fetch_url
