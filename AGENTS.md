# Development rules

- This repository is the only public Lifecycle dependency for Runtime and other modules.
- Do not import `github.com/domainry/domainry-lifecycle` or Runtime implementation packages.
- Do not expose database handles, SQL, ORM builders, or implementation structs in public contracts.
- Dependencies must use published semantic-version tags; local directory replacements are forbidden.
