# Domainry Lifecycle SDK

Deployment-neutral Lifecycle contracts, boundary models, host ports, and
binding interfaces. This module never contains Lifecycle domain behavior,
persistence contracts, or Runtime implementation.

## Package layout

- The root package is the stable Lifecycle `Factory` and `Binding` entrypoint.
- `contract` contains cross-owner cleanup, request-idempotent subject execution,
  and artifact protocols.
- `model` and `access` contain serialized boundary values and explicit access scope.
- `modulehost` describes infrastructure borrowed by an embedded Lifecycle module.

Domain policy, application orchestration, repositories, DDL, and DML remain in
the Lifecycle implementation. The SDK exposes business capabilities through
the root `Governance`, `System`, and `LocalWorkers` contracts, never through a
repository or transaction escape hatch.

Run `go test ./...` before publishing an immutable SDK version.
