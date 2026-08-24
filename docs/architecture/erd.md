# ERD — hello-word-14

## Scope

Schema supports one stored greeting displayed by public landing page.

## Tables

### `greetings`

| Column | Type | Null | Default | Notes |
|---|---|---|---|---|
| `id` | `smallint` | no | none | Primary key. Only row `1` is used. |
| `text` | `text` | no | none | Plain greeting string rendered exactly as stored. |
| `created_at` | `timestamptz` | no | `now()` | Audit timestamp. |
| `updated_at` | `timestamptz` | no | `now()` | Audit timestamp. |

Constraints:

- Primary key: `greetings_pkey` on `id`.
- Check: `id = 1` to enforce one-row table.
- Check: `length(text) > 0` to prevent empty rendered greeting.

Seed data:

```sql
insert into greetings (id, text) values (1, 'Hello Word');
```

Queries served:

- `GET /v1/greeting` reads row `id = 1` through `greetings_pkey`; no extra index needed.

## Story extension — Centered Hello Word page

No new entities. Existing `greetings` table satisfies reviewed UI mock contract:

```ts
{
  greeting: {
    text: "Hello Word",
  },
}
```

Backend API reads `greetings.text` for row `id = 1` and returns it as `greeting.text`.

## Migration bookkeeping

### `schema_migrations`

Created by backend migrator before applying migrations.

| Column | Type | Null | Default | Notes |
|---|---|---|---|---|
| `version` | `text` | no | none | Migration filename. Primary key. |
| `applied_at` | `timestamptz` | no | `now()` | Apply timestamp. |

## Relationships

No foreign keys. `greetings` is standalone.

## Migration plan

Forward:

1. Create `schema_migrations` if missing.
2. Create `greetings` with columns, constraints, and primary key listed above.
3. Seed row `(1, 'Hello Word')` if missing.

Backward:

1. Drop `greetings`.
2. Keep `schema_migrations` unless full database reset is requested.

Safety on populated tables:

- Safe for empty database.
- Safe for populated database when row `id = 1` is absent or already contains approved value.
- Not safe to overwrite changed greeting text; backend migration must not replace existing row text.

## Notes

- No user table because page is public and auth is out of scope.
- No edit history because editing greeting is out of scope.
- No soft delete because missing row has no approved UI state.
