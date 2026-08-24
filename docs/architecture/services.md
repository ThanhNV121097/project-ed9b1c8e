# Services — hello-word-14

## Conventions

- Backend serves paths without `/api` prefix.
- Public API version prefix is `/v1`.
- JSON fields use lower camelCase.
- Success responses return endpoint-specific objects.
- Error responses share one envelope.

## Error envelope

```json
{
  "error": {
    "code": "internal_error",
    "message": "Request failed"
  }
}
```

| Field | Type | Notes |
|---|---|---|
| `error.code` | string | Stable machine-readable code. |
| `error.message` | string | Generic human-readable message. Do not expose internals. |

Common errors:

| HTTP | Code | Message | When |
|---|---|---|---|
| `404` | `not_found` | `Greeting not found` | Greeting row missing. |
| `500` | `internal_error` | `Request failed` | Database or unexpected server failure. |
| `503` | `service_unavailable` | `Service unavailable` | Health check dependency unavailable. |

## Endpoints

### Health check

`GET /healthz`

Purpose: container and runtime health probe.

Auth: none.

Request body: none.

Success response: HTTP `200`, plain text body:

```text
ok
```

Failure response: HTTP `503`, error envelope:

```json
{
  "error": {
    "code": "service_unavailable",
    "message": "Service unavailable"
  }
}
```

Rules:

- Return 200 only after migrations succeeded.
- Return 200 only when `SELECT 1` against PostgreSQL succeeds.

### Read greeting

`GET /v1/greeting`

Purpose: frontend reads stored greeting value.

Auth: none. Public by SRS.

Request body: none.

Query parameters: none. Ignore query string.

Success response: HTTP `200`:

```json
{
  "greeting": {
    "text": "Hello Word"
  }
}
```

Response fields:

| Field | Type | Notes |
|---|---|---|
| `greeting` | object | Wrapper matching reviewed UI mock module. |
| `greeting.text` | string | Stored greeting text, rendered exactly by frontend. |

Reviewed UI mock contract:

```ts
{
  greeting: {
    text: "Hello Word",
  },
}
```

Errors:

| HTTP | Code | Message | When |
|---|---|---|---|
| `404` | `not_found` | `Greeting not found` | Row `greetings.id = 1` missing. |
| `500` | `internal_error` | `Request failed` | Database or unexpected server failure. |

Rules:

- Read only row `greetings.id = 1`.
- Use parameterized SQL even for fixed row query.
- Use request context for database call.
- Do not trim, transform, localize, or substitute `greeting.text`.
- Do not add loading, empty, or error UI states in backend contract; errors use envelope only.

## Security and reliability

- No authentication; endpoint is public by SRS.
- No request parameters; ignore query string.
- Use parameterized SQL even for fixed row query.
- Use request context for database calls.
- Do not log secrets or `DATABASE_URL`.
