# AGENTS.md

## Project Role

This repository is a macOS AI English learning platform. The product is organized around three learning modules:

- Listening and speaking
- Reading
- Writing

AI is a platform capability embedded into each module. It should not be implemented as a generic chatbot layer detached from learning workflows.

## Technical Direction

Follow the technical design in `docs/technical-design/technology-stack.md`.

Primary stack:

- Go
- Wails
- React
- TypeScript
- PostgreSQL
- pgvector
- AI Provider Adapter

## Architecture Rules

- Keep platform utilities in `internal/platform/`.
- Keep shared data structures in `internal/models/`.
- Keep cross-module learning capabilities in `internal/core/`.
- Keep product modules in `internal/modules/`.
- Keep user account and authentication capabilities in `internal/modules/user/`; user is a product module, not part of `internal/core/profile/`.
- Keep third-party or infrastructure adapters in `internal/adapters/`.
- Keep store/database operations in `internal/adapters/storage/`, with PostgreSQL-specific code under `internal/adapters/storage/postgres/`.
- Keep reusable constants in the appropriate `internal/platform/` package instead of scattering constants through services.
- Do not force every module into the same third-level directory shape before it needs it.
- Add deeper folders only when a module becomes complex enough to justify them.
- Keep frontend under `frontend/src/` with `main.tsx`, `App.tsx`, `routes/`, `features/`, `components/`, `api/`, `types/`, `styles/`, and `lib/`.
- Each `features/<name>/` folder is self-contained: page component, page-specific CSS, business-logic hooks, and an `index.ts` re-export.
- Use `@/` path alias instead of relative `../../` imports across directories.
- API modules in `api/` reuse the shared HTTP client from `api/client.ts` and expose domain objects (e.g. `authApi.login()`).
- Type definitions live in `types/` with `Req`/`Resp` suffixes matching backend DTOs.

## Models Layer

Use `internal/models/` for shared structures:

- `dto/`: request and response structures exposed to frontend or external interfaces.
- `db/`: database row models and persistence-facing structures.
- `common/`: shared enums, pagination, IDs, timestamps, and generic response metadata.

- Request and response DTO type names must use `Req` and `Resp` suffixes.
- Put domain-specific DTOs under matching subdirectories, such as `internal/models/dto/user/`.
- Do not define request/response DTOs inside service files.

Avoid scattering duplicate DTOs across modules unless the structure is truly module-specific.

## Error Handling Rules

- Keep shared error codes in `internal/platform/errors/`.
- Use the unified error response wrapper instead of building ad hoc `{code, message}` responses in handlers.
- Service and adapter code should return unified platform errors for expected user-facing failures.
- HTTP handlers should translate errors through the shared platform error mapping.

## AI Integration Rules

- Module code should not directly depend on any AI vendor SDK.
- Use an internal provider abstraction before integrating OpenAI, DeepSeek, Qwen, Doubao, GLM, or other vendors.
- Prefer structured JSON outputs for scoring, review, question generation, and learning reports.
- Log AI usage metadata such as provider, model, purpose, latency, token usage, and cost when available.

## Database Rules

- Use PostgreSQL as the primary database.
- Use pgvector for embeddings and vector search.
- Keep SQL migrations in `migrations/postgres/`.
- Keep PostgreSQL adapter code in `internal/adapters/storage/postgres/`.
- Desktop frontend must not connect to PostgreSQL directly.

## Documentation Rules

- Product requirements live in `docs/prd/`.
- Technical design lives in `docs/technical-design/`.
- Frontend visual style rules live in `docs/design-system/frontend-style-guide.md`.
- Architecture notes live in `docs/architecture/`.
- Architecture decisions live in `docs/decisions/`.
- API/interface notes live in `docs/api/`.

## Frontend Style Rules

- Use the frontend visual style guide before designing or implementing UI.
- Keep the product white, restrained, premium, and learning-workspace oriented.
- Avoid purple-blue AI gradients, robot icons, magic-wand metaphors, glow effects, and chatbot-first layouts.
- AI should appear as feedback, coaching, explanation, and review inside the workflow.

## Third-Party Library Policy

- Before hand-writing a non-trivial utility (JWT, password hashing, WebSocket, rate limiting, etc.), check whether a mature, well-maintained open-source Go library already covers the use case.
- If such a library exists, pause and ask the user: "This feature has an existing open-source library (e.g., `<library-name>`). Would you prefer to use it instead of a custom implementation?"
- Prefer libraries that are widely adopted, have stable APIs, and are actively maintained.
- Document the choice in `docs/decisions/` when the trade-off is non-obvious (e.g., hand-rolled vs library for security-sensitive code).

## Development Rules

- Prefer small, focused files.
- Keep module boundaries clear.
- Do not introduce large abstractions before the use case needs them.
- For Go work, Codex should load the `golang-how-to` skill and any relevant secondary Go skills before coding or reviewing.
- For feature work, Codex should perform static compilation only unless the user asks for automated tests.
- Backend behavior is verified by the user through HTTP requests at the end of the task.
- Keep optional HTTP test notes or request examples in `tests/http/` when useful.
- Use `go test ./...` only when the user asks for Go tests or when working on low-level pure Go utilities.
- If frontend tooling is added later, document the exact install and verification commands.
