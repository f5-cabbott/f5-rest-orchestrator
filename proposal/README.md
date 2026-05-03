# Proposal: Replace `f5-rest-orchestrator` with ACME-driven cert lifecycle

This directory holds the design artifacts for retiring the .NET Keyfactor
Universal Orchestrator extension in this repo and replacing it with an
ACMEv2 integration between BIG-IP (TMOS 17.5+) and a Keyfactor AnyCA Gateway
REST 24.2+ instance running the ACME plugin.

The full project scope and functional spec live in the conversation thread
that produced this proposal; this README is the entry point for the
artifacts.

## Layout

```
proposal/
├── README.md
├── schema/                      # JSON Schemas — source of truth for manifests
│   ├── gatewayconfig.schema.json
│   ├── bigipdevice.schema.json
│   └── managedcertificate.schema.json
├── examples/                    # Worked example manifests
│   ├── gateway.yaml
│   ├── device-bigip-dc1-01.yaml
│   └── cert-vip-app1.yaml
└── f5acmectl/                   # Go skeleton for the provisioning tool
    ├── go.mod
    ├── main.go
    ├── cmd/                     # cobra subcommands (stubs)
    ├── manifest/                # in-memory types + loader interface
    ├── gateway/                 # Keyfactor AnyCA Gateway admin-API client interface
    ├── bigip/                   # iControl REST client interface for native ACME
    ├── secrets/                 # pluggable secret-backend interface
    └── state/                   # locked, versioned snapshot backend interface
```

## Status

| Artifact | State |
|---|---|
| JSON Schemas | Draft v1 — covers all three kinds, validates the example manifests |
| Example manifests | Draft v1 — illustrative only, not real device data |
| `f5acmectl` Go skeleton | Skeleton only — every subcommand is a stub returning `not implemented` |
| Gateway / BIG-IP / secrets / state client interfaces | Draft v1 — no concrete implementations yet |

## What's intentionally missing

- **Concrete client implementations.** Gateway and BIG-IP packages define
  interfaces only. Concrete impls are deferred until the lab phase, when
  endpoint paths and request/response shapes can be confirmed against a
  real gateway 24.2 + TMOS 17.5+ pair.
- **JSON Schema → Go struct generator.** For now, types are hand-maintained
  in `manifest/types.go` and must be kept in sync with `schema/*.json`.
  When the schemas stabilize, generate one from the other.
- **Tests.** The skeleton compiles (or will, once `cobra` and `yaml.v3` are
  vendored) but exercises nothing. Test scaffolding lands with the first
  real implementation.
- **`acme-dns` fallback.** The proposal assumes the BIG-IP native ACME
  client supports the chosen DNS provider. Lab phase verifies. If not,
  `acme-dns` delegation is the documented escape hatch (no code change
  required in `f5acmectl` — it manages BIG-IP config either way).
- **NGINX coverage.** Phase 2; out of scope for this proposal.

## Decisions that are still open

These are the same items called out in the project doc. Listed here for
quick reference:

1. EAB scope — defaulted to `per-partition` in the example; confirm.
2. Secret backend — Vault is the reference; confirm vs AWS SM / Azure KV.
3. DNS provider list for DNS-01 — pending lab confirmation against TMOS 17.5.
4. CI platform for the GitOps pipeline.
5. Reconciler runtime location (CI cron / k8s CronJob / systemd timer).
6. CA bundle / trust store automation — separate project or appended.
7. Cert revocation policy on decommission — automatic or manual.

## Relationship to the legacy code

Nothing in this directory replaces or modifies the legacy .NET extension at
the repo root. Cutover is per-partition, parallel-run for one renewal cycle
per cert, then the legacy store-type registration is removed for that
partition. After the last partition migrates, the .NET code is archived and
this repo (or its successor) becomes the single source of truth.
