// Package bigip abstracts the F5 iControl REST surface that f5acmectl uses to
// configure the BIG-IP native ACMEv2 client (TMOS 17.5+) and to bind issued
// certificates to client/serverSSL profiles. Endpoint paths are intentionally
// not encoded into the interface; concrete implementations target a specific
// TMOS version.
package bigip

import (
	"context"
	"crypto/x509"
	"time"
)

// Client is the minimal iControl REST surface f5acmectl needs against one device.
type Client interface {
	// TMOSVersion returns the running TMOS version (e.g. "17.5.0.1").
	TMOSVersion(ctx context.Context) (string, error)

	// ListPartitions returns partition names visible to the API user.
	ListPartitions(ctx context.Context) ([]string, error)

	// ConfigureACMEAccount declares (or updates) an ACME account binding on
	// a partition: directory URL, contact email, EAB kid + HMAC, key type.
	// Idempotent on (partition, directoryURL, kid).
	ConfigureACMEAccount(ctx context.Context, p ACMEAccountParams) error

	// DeleteACMEAccount removes an ACME account binding.
	DeleteACMEAccount(ctx context.Context, partition, accountName string) error

	// DeclareManagedCertificate ensures the device has an ACME-managed cert
	// declaration for the given identifiers, key type, and challenge config.
	// Triggers an order on first creation; subsequent calls are no-ops if
	// nothing changed.
	DeclareManagedCertificate(ctx context.Context, p ManagedCertParams) error

	// BindCertificateToProfile points an SSL profile at the issued cert/key.
	// Uses iControl REST transactions for atomic swap.
	BindCertificateToProfile(ctx context.Context, b BindParams) error

	// ListInstalledCerts returns metadata for every cert in a partition's
	// filestore.
	ListInstalledCerts(ctx context.Context, partition string) ([]InstalledCert, error)

	// FetchCert returns the parsed leaf certificate by filestore name.
	FetchCert(ctx context.Context, partition, name string) (*x509.Certificate, error)

	// DeleteCert removes a cert/key pair from the filestore. Caller must
	// confirm no profiles still reference it.
	DeleteCert(ctx context.Context, partition, name string) error
}

type ACMEAccountParams struct {
	Partition    string
	AccountName  string
	DirectoryURL string
	ContactEmail string
	EABKID       string
	EABHMACKey   string // sourced from secret backend; never logged
	KeyType      string
}

type ManagedCertParams struct {
	Partition       string
	Name            string
	Identifiers     []string
	KeyType         string
	Challenge       string // http-01 | dns-01 | tls-alpn-01
	DNSProvider     string
	RenewBeforeDays int
	AccountName     string
}

type BindParams struct {
	Partition   string
	ProfileKind string // clientSSLProfile | serverSSLProfile
	ProfileName string
	CertName    string
	KeyName     string
}

type InstalledCert struct {
	Partition string
	Name      string
	NotBefore time.Time
	NotAfter  time.Time
	Subject   string
	SAN       []string
	Issuer    string
}
