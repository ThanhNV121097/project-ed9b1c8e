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

## Endpoints

### Health check

`GET /healthz`

Purpose: container and runtime health probe.

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

Request body: none.

Success response: HTTP `200`:

```json
{
  "text": "Hello Word"
}
```

Response fields:

| Field | Type | Notes |
|---|---|---|
| `text` | string | Stored greeting text, rendered exactly by frontend. |

Errors:

| HTTP | Code | Message |
|---|---|---|
| `404` | `not_found` | `Greeting not found` |
| `500` | `internal_error` | `Request failed` |

## Security and reliability

- No authentication; endpoint is public by SRS.
- No request parameters; ignore query string.
- Use parameterized SQL even for fixed row query.
- Use request context for database calls.
- Do not log secrets or `DATABASE_URL`.
