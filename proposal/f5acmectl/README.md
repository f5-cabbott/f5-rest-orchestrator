# f5acmectl

Declarative provisioning tool for the BIG-IP native ACMEv2 client + Keyfactor
AnyCA Gateway REST ACME plugin integration.

> **Status: skeleton.** Every subcommand is a stub. Manifest types and client
> interfaces are defined; concrete implementations land during the lab phase.

## Subcommands

| Command | Purpose |
|---|---|
| `validate` | Schema-validate manifests; check secret refs exist (no value reads) |
| `discover` | Read-only snapshot of current gateway + BIG-IP state |
| `plan` | Diff desired manifests against discovered state; emit a change set |
| `apply` | Execute a plan idempotently with state-lock + audit |
| `verify` | Post-apply: certs installed, chains valid, profiles bound |
| `reconcile` | Read-only drift / expiry / orphan detection on a schedule |
| `revoke` | ACME revokeCert + filestore cleanup |
| `rotate-eab` | Mint new EABs, migrate ACME accounts, retire old EABs |
| `state` | Inspect/lock/unlock the state backend |
| `version` | Print build info |

## Build

```sh
go build -o f5acmectl .
```

(Will fail until dependencies are tidied: `go mod tidy`. The skeleton imports
cobra and yaml.v3 only.)

## Layout

```
f5acmectl/
├── main.go            # entry point
├── cmd/               # cobra command tree (one file per subcommand)
├── manifest/          # GatewayConfig / BigIPDevice / ManagedCertificate types + Loader
├── gateway/           # Keyfactor AnyCA Gateway admin-API client interface
├── bigip/             # F5 iControl REST client interface (TMOS 17.5+ ACME)
├── secrets/           # Vault / AWS SM / Azure KV pluggable backend
└── state/             # Locked, versioned snapshot backend (S3 / GCS / Azure Blob)
```

## Design constraints

- **No secrets in manifests, state, or logs.** Manifests carry `secret://` URIs
  and `CredentialsRef`s only. State snapshots carry references, not values.
  Logs redact authorization headers and EAB HMAC keys.
- **Idempotent.** `apply` is safe to re-run; nothing happens if desired ==
  observed.
- **Locked.** Concurrent applies are blocked at the state backend.
- **Auditable.** Every mutation is tagged with actor, git SHA, target,
  and gateway request-id.
- **Single static binary.** Distributable as one file or a container image.
  No runtime plugin loading.
