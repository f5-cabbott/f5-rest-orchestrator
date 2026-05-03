// Package gateway abstracts the Keyfactor AnyCA Gateway REST admin API used
// by f5acmectl to administer EAB credentials and inspect ACME plugin state.
// Issuance itself happens between the BIG-IP and the gateway's ACME directory;
// f5acmectl does not act as an ACME client.
package gateway

import "context"

// Client is the minimal admin-API surface f5acmectl needs.
//
// All methods are expected to be idempotent where the underlying API allows
// it, and to surface request IDs in errors for audit correlation.
type Client interface {
	// ListEABs enumerates configured EAB credentials, optionally filtered by
	// label selectors carried in opaque tags.
	ListEABs(ctx context.Context, selector map[string]string) ([]EAB, error)

	// CreateEAB mints a new EAB credential bound to the given policy mapping.
	// The returned EAB carries the HMAC key; the caller is responsible for
	// writing it to the secret backend immediately.
	CreateEAB(ctx context.Context, req CreateEABRequest) (*EAB, error)

	// RevokeEAB revokes the EAB identified by kid. Existing ACME accounts
	// already created under that EAB continue to function unless explicitly
	// deactivated; callers should rotate accounts first.
	RevokeEAB(ctx context.Context, kid string) error

	// DescribePolicy returns the policy mapping (template, allowed
	// identifiers, key constraints) bound to a kid.
	DescribePolicy(ctx context.Context, kid string) (*Policy, error)

	// HealthCheck verifies admin API reachability and authentication.
	HealthCheck(ctx context.Context) error
}

type EAB struct {
	KID          string            // public account identifier
	HMACKey      string            // base64url; populated only on CreateEAB response
	PolicyName   string            // gateway-side policy mapping name
	Tags         map[string]string // opaque labels (env, device, partition)
	Disabled     bool
}

type CreateEABRequest struct {
	PolicyName string
	Tags       map[string]string
}

type Policy struct {
	Name              string
	Template          string
	AllowedIdentifiers []string
	AllowedKeyTypes   []string
}
