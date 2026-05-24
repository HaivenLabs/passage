# Passage Product Plan

Version: 1

Passage is the Haiven Labs registration, identity, authentication, authorization, account, organization, and product access platform.

It exists so qurl, Heard, and future Haiven products can share one customer identity foundation without every product rebuilding signup, login, teams, roles, invitations, sessions, product access, or account recovery.

The product goal is simple: one Haiven account, many Haiven products, no extra friction.

## 1. Positioning

Passage is a standalone Haiven product and platform service.

For customers, Passage should feel invisible. They sign up once, join or create an organization once, and move between Haiven products naturally.

For developers, Passage should feel like a drop-in identity layer. Products integrate through clear APIs, SDKs, webhooks, and shared UI primitives.

For self-hosters, Passage should be understandable, runnable, and replaceable. The meaningful core should not require paid vendors.

## 2. Core Promise

Passage owns the shared account layer.

Passage owns:

- User registration
- Login and logout
- Session lifecycle
- Account recovery
- Email verification
- Multi-factor authentication
- Passkeys
- Organizations
- Teams and memberships
- Invitations
- Roles and permissions
- Product access
- Tenant and organization context
- API clients and service credentials
- Audit trail for account and access changes
- Identity provider integrations
- Admin and operator tooling

Products own their product domains.

qurl owns QR projects, payloads, designs, exports, static QR behavior, dynamic QR behavior, tracking choices, and asset delivery.

Heard owns feedback campaigns, locations, guest feedback, recovery cases, restaurant operations data, review prompts, and customer recovery workflows.

Passage must not absorb qurl or Heard domain facts.

## 3. Product Boundaries

A Haiven product should use Passage to answer:

- Who is this user?
- Is this session valid?
- Which organization is active?
- Which products can this user access?
- Which organization roles does this user have?
- Which product roles or entitlements apply?
- Can this user perform this account-level or product-entry action?

A Haiven product should answer its own domain questions.

Examples:

- qurl decides whether a user can edit a specific QR project.
- Heard decides whether a user can manage a specific location or recovery case.
- Passage decides whether the user belongs to the organization and has access to the product.

## 4. qurl Requirements

qurl needs Passage to support:

1. Anonymous QR creation with no account required.
2. Anonymous-to-account handoff when a user wants to save a QR project.
3. Account signup from the save flow without losing unsaved work.
4. Login from the save flow without losing unsaved work.
5. Saved QR dashboard access after authentication.
6. Personal workspace for solo users.
7. Organization workspace for teams later.
8. Product entitlement checks for advanced features.
9. API access for future developer usage.
10. Clear distinction between direct QR codes and hosted dynamic QR codes.

qurl must not require Passage for anonymous direct QR creation and download.

## 5. Heard Requirements

Heard needs Passage to support:

1. Restaurant admin registration.
2. Organization creation during onboarding.
3. Location setup after organization creation.
4. Team invitations for managers and staff.
5. Organization roles such as owner, admin, manager, and viewer.
6. Product-level access to Heard.
7. Secure admin session management.
8. Audit trail for membership, role, and configuration changes.
9. Future SSO for larger restaurant groups.
10. Future multi-location and enterprise governance.

Guest feedback flows should not require Passage unless the guest is intentionally creating an account in a future product surface.

## 6. MVP Scope

The MVP should prove Passage can power both qurl and Heard without slowing either product down.

### MVP must include

- Email and password registration
- Email verification
- Login and logout
- Password reset
- Session cookies for web apps
- Personal account creation
- Organization creation
- Organization membership
- Organization invitation flow
- Organization role assignment
- Active organization context
- Product access grants for qurl and Heard
- Basic admin console
- Product integration API
- Product-facing SDK or client package
- OpenAPI contract
- JSON Schema config contract
- Docker-first local development
- PostgreSQL persistence
- Audit events for account and membership changes
- Webhook events for downstream products
- Seeded local development data

### MVP should defer

- Enterprise SSO
- SCIM provisioning
- Social login
- Passkeys
- MFA
- Fine-grained policy language
- Billing integration
- Marketplace integrations
- Advanced identity brokering
- Device management
- Risk scoring

Deferring these features keeps the first slice shippable.

## 7. ZITADEL-Equivalent Capability Map

Passage should eventually cover the meaningful ZITADEL-style identity platform surface, while staying simpler and more Haiven-specific.

### Identity and login

- Username or email login
- Password authentication
- Passwordless email link login
- Passkeys and WebAuthn
- Multi-factor authentication
- Account recovery
- Email verification
- Phone verification later
- Login policies
- Password policies
- Session lifetime policies
- Device and session visibility
- User lockout controls

### Organizations and projects

- Organizations
- Organization settings
- Organization membership
- Invitations
- Teams
- Product access grants
- Product projects or applications
- Application callback URLs
- Application logout URLs
- Application allowed origins
- Environment-aware app registration

### Authorization

- System roles
- Organization roles
- Product roles
- Resource-scoped permissions
- Service accounts
- Machine-to-machine clients
- API keys or client credentials
- Token claims customization
- Policy evaluation hooks

### Federation and external identity

- Google login
- GitHub login
- Microsoft login
- Apple login later
- SAML identity provider support
- OIDC identity provider support
- Enterprise SSO mapping
- External identity linking
- Domain-based organization discovery

### Developer platform

- OIDC-compatible flows
- OAuth2-compatible flows
- REST admin API
- Public JWKS endpoint
- Token introspection endpoint later
- SDKs for TypeScript and Go
- Webhooks
- CLI later
- Terraform provider later

### Operations

- Admin console
- Operator console
- Audit log
- Event log
- Health checks
- Structured logs
- Metrics and traces
- Backup and restore documentation
- Migration tooling
- Local fake email provider
- Local seeded data

### Self-hosting

- Docker Compose local runtime
- Production container images
- PostgreSQL as primary store
- Config through environment variables and JSON config
- No mandatory paid vendor dependency
- Optional SMTP provider
- Optional object storage later
- Clear upgrade and migration path

## 8. Architecture

Passage starts as a modular Go monolith with a React or Next.js admin and account UI.

Use a modular monolith first because identity is already complex enough. The product needs clear boundaries, not premature distributed systems.

### Backend modules

- accounts
- users
- credentials
- sessions
- organizations
- memberships
- invitations
- roles
- permissions
- products
- applications
- oauth
- webhooks
- audit
- notifications
- admin

### Frontend surfaces

- hosted signup
- hosted login
- hosted password reset
- hosted account settings
- organization switcher
- organization settings
- member management
- invitation acceptance
- product access screen
- admin console

### Packages

- `packages/contracts` for OpenAPI and JSON Schema
- `packages/passage-client` for TypeScript product integration
- `packages/passage-go` for Go product integration later
- `packages/ui` for shared account UI primitives if this repo hosts them

## 9. Recommended Repository Layout

```text
apps/admin                 Passage admin and account UI
apps/hosted                Hosted auth UI if separated from admin
backend                    Go API and workers
packages/contracts         OpenAPI and JSON Schema contracts
packages/passage-client    TypeScript client for Haiven products
packages/passage-go        Go client for Haiven products, later
docs                       Product and architecture docs
scripts                    Local bootstrap and checks
```

## 10. Core Data Model

Initial entities:

- users
- user_emails
- credentials
- sessions
- organizations
- organization_memberships
- invitations
- roles
- permissions
- products
- product_access_grants
- applications
- oauth_clients
- service_accounts
- api_clients
- audit_events
- webhook_subscriptions
- webhook_deliveries
- outbox_events

IDs should be Passage-owned UUIDs. External provider IDs should be stored as references, not primary keys.

## 11. Registration Flows

### Standard signup

1. User enters email and password.
2. Passage creates a pending user.
3. Passage sends verification email.
4. User verifies email.
5. Passage creates default personal workspace or asks whether to create an organization.
6. User lands in the requesting product with active account context.

### qurl save flow

1. User creates QR anonymously.
2. User clicks save.
3. qurl stores pending QR draft locally or server-side with short-lived handoff token.
4. Passage handles signup or login.
5. Passage redirects back to qurl with verified session.
6. qurl attaches the draft to the user or active organization.
7. User lands on the saved QR project.

### Heard restaurant onboarding

1. User starts Heard onboarding.
2. Passage handles signup or login.
3. Passage creates or selects organization.
4. Heard creates restaurant or location setup flow.
5. Heard owns location data after Passage confirms organization context.

### Invitation acceptance

1. Existing admin invites a user by email.
2. Passage sends invitation.
3. Invitee accepts.
4. Existing users authenticate and accept membership.
5. New users register and accept membership.
6. Passage redirects to the correct product and organization.

## 12. Session and Token Model

For web apps, prefer secure HTTP-only cookies controlled by Passage and validated by products through a lightweight middleware or token verification package.

For APIs, support bearer tokens with signed claims. Products should verify issuer, audience, expiration, organization context, product access, and role claims.

Do not expose long-lived credentials by default. Service credentials should be explicitly created, named, scoped, and auditable.

## 13. Product Integration Model

Products integrate with Passage using:

- Hosted login and registration URLs
- Callback URLs
- Product client IDs
- Signed session or access tokens
- JWKS-based token verification
- SDK helpers
- Admin API for product access checks where needed
- Webhooks for membership and entitlement changes

Products should not read Passage tables directly.

## 14. Product Access Model

Passage should represent Haiven products explicitly.

Initial products:

- qurl
- Heard
- Passage
- StackSmith later

A user can belong to multiple organizations. An organization can have access to multiple products. A user can have different roles per organization and product.

Minimum roles:

- owner
- admin
- manager
- member
- viewer

Product-specific roles may be mapped or extended inside each product.

## 15. Admin Console

The Passage admin console should include:

- users
- organizations
- memberships
- invitations
- products
- product access grants
- applications
- OAuth clients
- service accounts
- audit log
- webhook deliveries
- system health

Operator-only features should be separated from customer organization admin features.

## 16. API Surface

Initial API groups:

- `/api/v1/auth`
- `/api/v1/users`
- `/api/v1/account`
- `/api/v1/organizations`
- `/api/v1/memberships`
- `/api/v1/invitations`
- `/api/v1/products`
- `/api/v1/applications`
- `/api/v1/webhooks`
- `/api/v1/audit`
- `/api/v1/admin`

All public APIs must have OpenAPI contracts.

Configuration schemas should use JSON Schema.

## 17. Events and Webhooks

Passage should emit durable events for important changes.

Initial events:

- `user.created` version `1`
- `user.email_verified` version `1`
- `organization.created` version `1`
- `membership.created` version `1`
- `membership.role_changed` version `1`
- `membership.removed` version `1`
- `invitation.created` version `1`
- `invitation.accepted` version `1`
- `product_access.granted` version `1`
- `product_access.revoked` version `1`
- `application.created` version `1`
- `api_client.created` version `1`

Use a transactional outbox for reliable event publishing.

## 18. Security Requirements

Passage must be treated as critical infrastructure.

Minimum requirements:

- Secure password hashing
- Email verification before sensitive account operations
- Password reset token expiration
- Session expiration
- Session revocation
- CSRF protection for cookie-based flows
- Rate limiting on sensitive endpoints
- Server-side authorization checks
- Audit trail for account, organization, role, product access, and credential changes
- Secrets managed outside source control
- Safe logging with no passwords, reset tokens, session tokens, or API credentials in logs
- Tenant and organization isolation tests
- Permission tests for account and organization actions

## 19. Observability

Passage should include:

- health endpoint
- readiness endpoint
- structured logs
- request IDs
- audit events
- metrics for auth flow success and failure
- trace support for login, signup, invitation, token, and webhook flows
- worker metrics for outbox and webhook delivery

## 20. First Vertical Slice

The first shippable slice should prove qurl and Heard can both use Passage.

### Slice 1 outcome

A developer can run Passage, qurl, and Heard locally. A user can register once through Passage, save a qurl project, then access Heard admin with the same account and organization context.

### Slice 1 scope

- Docker Compose local runtime
- PostgreSQL database
- Go API
- hosted signup page
- hosted login page
- email verification with local fake provider
- password reset with local fake provider
- session cookie
- organization creation
- organization switcher
- product registry with qurl and Heard
- product access grant
- qurl integration stub
- Heard integration stub
- audit log entries
- OpenAPI contract
- basic tests

### Slice 1 acceptance criteria

- User can create account.
- User can verify email locally.
- User can log in.
- User can create organization.
- User can switch active organization.
- qurl can redirect to Passage and receive a valid authenticated return.
- Heard can redirect to Passage and receive a valid authenticated return.
- Products can verify a Passage-issued session or token.
- Product access can be granted and checked.
- Membership and product access changes write audit events.
- Local development starts with one documented command.

## 21. Build Sequence

### Phase 1: Foundation

- repo setup
- docs
- database migrations
- backend skeleton
- frontend skeleton
- OpenAPI contract
- local runtime
- health checks

### Phase 2: Account basics

- registration
- login
- logout
- email verification
- password reset
- session persistence
- account settings

### Phase 3: Organization basics

- create organization
- active organization context
- membership model
- invitations
- roles
- organization settings

### Phase 4: Product integration

- product registry
- qurl product application
- Heard product application
- callback URLs
- product access grants
- token claims
- SDK helpers

### Phase 5: Admin and operations

- admin console
- audit log viewer
- webhook delivery
- structured logs
- metrics
- traces
- backup and migration docs

### Phase 6: ZITADEL-parity expansion

- passkeys
- MFA
- social login
- OIDC provider support
- SAML provider support
- SCIM
- service accounts
- machine-to-machine auth
- advanced policies
- CLI
- Terraform provider

## 22. Non-goals for the First Release

- Replacing every enterprise identity product immediately
- Supporting every OAuth and SAML edge case
- Building a policy language before product needs prove it
- Creating a generic auth product detached from Haiven needs
- Forcing login into anonymous qurl creation
- Putting guest identity into Heard feedback flows prematurely

## 23. Product Principles Specific to Passage

Passage should be boring where identity must be boring and beautiful where users touch it.

It should be invisible during happy paths, clear during recovery paths, strict during admin paths, and friendly during onboarding paths.

It should reduce friction for customers and reduce repeated work for Haiven products.

If a feature makes qurl or Heard slower to ship without making customers safer or happier, it does not belong in the first slice.
