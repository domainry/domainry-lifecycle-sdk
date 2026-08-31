# Domainry Lifecycle SDK

Deployment-neutral Lifecycle contracts, models, application services, host
ports, and binding interfaces. This module never imports the Lifecycle
implementation repository or Runtime implementation.

## Package layout

- The root package is the stable Lifecycle `Factory` and `Binding` entrypoint.
- `contract` contains cross-owner lifecycle execution and artifact protocols.
- `application`, `model`, `policy`, and `access` contain deployment-neutral lifecycle behavior and values.
- `persistence` owns Lifecycle persistence capability contracts.
- `modulehost` describes infrastructure borrowed by an embedded Lifecycle module.

Concrete DDL and DML remain in the Lifecycle implementation and use the host database and migration registrar; the SDK exposes only their contracts through `persistence`.

Run `go test ./...` before publishing an immutable SDK version.
