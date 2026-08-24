# Architecture Overview — hello-word-14

## 1. Scope

hello-word-14 is fullstack proof app: one public page, one backend API, one PostgreSQL row.

Included:

- Next.js frontend renders single centered message.
- Go backend serves health and greeting API.
- PostgreSQL stores greeting text.
- `docker compose up` from repo root boots DB, backend, frontend.

Not included:

- Authentication.
- Editing greeting.
- Multiple pages.
- Admin tools.
- Background jobs.

## 2. Tech stack

| Layer | Choice | Version / constraint | Reason |
|---|---|---|---|
| Frontend | Next.js App Router | 15.x, TypeScript | Matches team default and container contract. |
| Styling | Tailwind CSS | v3 | Existing project convention; tokenized CSS in `app/globals.css`. |
| Backend | Go | 1.22+ | Small HTTP API, fast static binary, team default. |
| Database | PostgreSQL | 16 in local compose | Needed because SRS requires text stored outside frontend. |
| Containers | Existing Dockerfiles + compose contract | Service folders under `code/` | Runtime and CI expect this layout. |
| CI | `.github/workflows/ci.yml` | Read-only | Runs Go build/vet/test, npm lint/build/test, token checks. |

## 3. Repository layout

```text
code/
  backend/
    cmd/api/main.go              # HTTP server entrypoint, exactly one main package
    internal/db/migrate.go       # embedded migrations and runner
    migrations/*.up.sql          # forward migrations, filename order
    migrations/*.down.sql        # rollback reference, not run on boot
    go.mod / go.sum              # Go module
    .env.example                 # backend env keys
    Dockerfile                   # existing service image contract
  frontend/
    app/layout.tsx               # root metadata and document shell
    app/page.tsx                 # story composition root only
    app/globals.css              # shared tokens and base styles
    package.json / package-lock.json
    next.config.js               # standalone output required by Dockerfile
    tailwind.config.ts
    postcss.config.js
    .env.example                 # frontend env keys
docs/
  architecture/
    overview.md                  # this file
    erd.md                       # schema contract
    services.md                  # API contract
  general/SRS.md                 # merged requirements
```

## 4. Runtime data flow

1. Browser requests frontend page.
2. Frontend story component later calls backend using `NEXT_PUBLIC_API_URL`.
3. Backend reads greeting from PostgreSQL.
4. Backend returns JSON response.
5. Frontend renders returned string centered on white background.

`app/page.tsx` must stay thin. Story work adds one import and one element; no story should rewrite scaffold files except its owned mount line.

## 5. Backend contracts

- Module path: `github.com/ThanhNV121097/project-ed9b1c8e/backend`.
- One executable package only: `code/backend/cmd/api`.
- Server reads `DATABASE_URL` before start.
- Server applies every pending migration from `code/backend/migrations/` before listening.
- Migration state lives in `schema_migrations(version text primary key, applied_at timestamptz not null default now())`.
- Migrations run in filename order and are no-op after recorded.
- `/healthz` returns 200 only after migrations succeed and `SELECT 1` succeeds.
- Port order: `PORT`, then `APP_PORT`, then `8080`.
- API routes use `/v1/...`; do not mount `/api/...` because deploy proxy strips `/api` before backend.

## 6. Frontend contracts

- Next.js App Router under `code/frontend/app`.
- `next.config.js` has `output: "standalone"` for container runtime.
- `npm run start` exists because runner uses it after build.
- `npm run lint` runs `next lint` and must pass.
- ESLint errors: unused vars, explicit `any`, max 300 non-comment lines, complexity 12.
- Server Components by default. Any component using browser APIs, event handlers, refs, state, or effects must start with first line `"use client"`.
- Every React component file uses `export default function ComponentName()`.
- `app/globals.css` owns all shared tokens: color, spacing, typography, radius, shadow, motion.
- CSS modules must use tokens, no hardcoded colors or un-tokenized lengths.

## 7. Env vars

### Root compose `.env.example`

| Key | Used by | Meaning |
|---|---|---|
| `POSTGRES_USER` | compose DB/backend | Local database user. |
| `POSTGRES_PASSWORD` | compose DB/backend | Local database password. |
| `POSTGRES_DB` | compose DB/backend | Local database name. |
| `BACKEND_PORT` | compose | Host port for backend. |
| `FRONTEND_PORT` | compose | Host port for frontend. |
| `NEXT_PUBLIC_API_URL` | frontend build/runtime | Browser-visible backend origin. |
| `IMAGE_REPO` | compose | Optional image repository prefix. |
| `IMAGE_TAG` | compose | Optional image tag. |

### Backend `code/backend/.env.example`

| Key | Meaning |
|---|---|
| `DATABASE_URL` | PostgreSQL DSN injected by runtime or compose. |
| `PORT` | HTTP listen port, preferred. |
| `APP_PORT` | Legacy fallback listen port. |

### Frontend `code/frontend/.env.example`

| Key | Meaning |
|---|---|
| `NEXT_PUBLIC_API_URL` | Backend origin visible to browser, for example `http://localhost:8080`. |

## 8. Run and verify

Local run:

```bash
docker compose --profile local up --build
```

Backend checks:

```bash
cd code/backend
go build ./...
go vet ./...
go test ./...
```

Frontend checks:

```bash
cd code/frontend
npm ci
npm run lint
npm run build
npm test --if-present
```

Health check:

```bash
curl http://localhost:8080/healthz
```

## 9. Naming conventions

| Area | Convention |
|---|---|
| Go packages | Lowercase, short, no underscores. |
| Go files | Lowercase words, underscores only when needed by Go convention. |
| Migrations | UTC timestamp prefix: `YYYYMMDDHHMMSS_description.up.sql`. |
| API paths | `/v1/resource`, nouns, no `/api` prefix. |
| JSON fields | lower camelCase. |
| React components | PascalCase filenames and `export default function ComponentName()`. |
| CSS tokens | `--category-name`, sourced from `design/design-system.md`. |
| Env vars | Upper snake case, listed in closest `.env.example`. |

## 10. Decisions

| Decision | Rejected alternative | Tradeoff |
|---|---|---|
| Fullstack shape | Static page with hardcoded text | SRS requires PostgreSQL row and backend API; fullstack adds services but proves pipeline end-to-end. |
| Go stdlib HTTP server | Web framework | Stdlib is enough for two routes; fewer dependencies and smaller review surface. |
| `pgx` database driver | `lib/pq` or ORM | Maintained driver with stdlib adapter; no ORM needed for one query. |
| Self-migrate on boot | Separate migration command | Runtime provides empty DB and no separate migrator; boot cost is small and reliability is higher. |
| Embed migrations in backend binary | Read migration files from working directory | Embed avoids container path drift; must keep embed file beside `migrations/`. |
| Next.js App Router | Pages Router or plain React | Team default, existing CI/container expectations. |
| Tokenized `globals.css` | Story-level design values | CI can catch drift early; scaffold must define all categories now. |
| No API envelope for success beyond data object | Global success wrapper | One response object is enough; shared envelope is only needed for errors. |

## 11. Risks and constraints

| Risk | Mitigation |
|---|---|
| Missing greeting seed blocks UI display | ERD migration seeds one row with `Hello Word`; backend story must read same row. |
| Frontend hardcodes greeting | Story review checks page source and network call against SRS AC-2. |
| Health passes before DB works | `/healthz` performs `SELECT 1` after migrations. |
| `/api` route mismatch | Services doc uses `/v1/...` only. |
| Token drift | CI checks used tokens and fallbacks; `globals.css` mirrors design system. |

## 12. Unknowns

None blocking. Product intentionally has no loading, empty, or error state in UI; backend still returns structured errors for API consumers.
