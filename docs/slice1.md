# Slice 1

## Goal

Create the first runnable Passage foundation slice.

Passage is the Haiven registration, identity, account, organization, session, role, and product access platform.

Slice 1 does not attempt to build the full identity surface. It creates the repo foundation that later slices will extend.

## Scope

- Docker-first local runtime
- Go API process
- PostgreSQL service
- Health endpoint
- Readiness endpoint
- OpenAPI contract placeholder
- Runtime config schema placeholder
- CI check path
- Product boundary documentation

## Rules

- Follow the Haiven Constitution.
- Build vertical slices.
- Keep qurl and Heard domain logic out of Passage.
- Keep product integration contract-first.
- Keep local development boring.

## Acceptance criteria

- `docker compose up --build` starts the API and database.
- `GET /healthz` returns healthy service status.
- `GET /readyz` checks database readiness.
- Backend tests run with `go test ./...`.
- Contract and config schema locations exist.
- Product docs explain that Passage owns shared identity and product access, not product-domain facts.
