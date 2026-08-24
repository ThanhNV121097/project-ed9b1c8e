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

## Migration bookkeeping

### `schema_migrations`

Created by backend migrator before applying migrations.

| Column | Type | Null | Default | Notes |
|---|---|---|---|---|
| `version` | `text` | no | none | Migration filename. Primary key. |
| `applied_at` | `timestamptz` | no | `now()` | Apply timestamp. |

## Relationships

No foreign keys. `greetings` is standalone.

## Notes

- No user table because page is public and auth is out of scope.
- No edit history because editing greeting is out of scope.
- No soft delete because missing row has no approved UI state.
