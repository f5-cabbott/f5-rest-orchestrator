// Package secrets defines the pluggable secret-backend interface. Manifests
// and state files only ever carry references; concrete values are fetched
// at apply/verify time and never logged.
package secrets

import "context"

// Backend is implemented by Vault, AWS Secrets Manager, and Azure Key Vault.
type Backend interface {
	// Name returns the backend identifier as it appears in CredentialsRef.Backend.
	Name() string

	// Exists reports whether a secret is reachable at path without reading
	// its value. Used by `f5acmectl validate` to fail fast on missing refs.
	Exists(ctx context.Context, path string) (bool, error)

	// Get reads the secret at path. Implementations must zero the returned
	// buffer when the caller calls Close on the returned Secret.
	Get(ctx context.Context, path string) (*Secret, error)

	// Put writes a secret at path. Used during EAB minting. Implementations
	// should fail if the path already exists unless overwrite is true.
	Put(ctx context.Context, path string, data map[string][]byte, overwrite bool) error

	// Delete removes a secret at path. Used during EAB rotation cleanup.
	Delete(ctx context.Context, path string) error
}

// Secret is a transient, owner-closed wrapper around secret material.
// Callers must call Close as soon as the value is no longer needed.
type Secret struct {
	Path string
	Data map[string][]byte
	close func() error
}

func NewSecret(path string, data map[string][]byte, closer func() error) *Secret {
	return &Secret{Path: path, Data: data, close: closer}
}

func (s *Secret) Close() error {
	if s == nil || s.close == nil {
		return nil
	}
	for k := range s.Data {
		// Best-effort wipe. Callers that need stronger guarantees should
		// handle key material in fixed-size byte slices outside this map.
		for i := range s.Data[k] {
			s.Data[k][i] = 0
		}
	}
	return s.close()
}
