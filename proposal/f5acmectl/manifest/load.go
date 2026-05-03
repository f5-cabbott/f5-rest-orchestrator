package manifest

import (
	"context"
	"fmt"
)

// Loader walks a manifest directory tree and returns a resolved Set.
//
// Implementations are expected to:
//   - parse every YAML document in the tree (multi-doc files allowed),
//   - validate each document against the JSON Schema for its kind,
//   - reject any literal secret material (only secret:// URIs and
//     CredentialsRef values are permitted),
//   - cross-reference ManagedCertificate.spec.deviceRef against device
//     metadata.name and partition membership,
//   - return a single Set per GatewayConfig (one environment).
type Loader interface {
	Load(ctx context.Context, root string) (*Set, error)
}

// ValidationError is the structured error returned by Loader.Load when
// validation fails. It carries one entry per offending document.
type ValidationError struct {
	Issues []Issue
}

type Issue struct {
	Path    string // file path relative to the manifest root
	Pointer string // JSON Pointer into the document
	Kind    string // GatewayConfig | BigIPDevice | ManagedCertificate | <unknown>
	Message string
}

func (e *ValidationError) Error() string {
	return fmt.Sprintf("manifest validation failed: %d issue(s)", len(e.Issues))
}
