// Package manifest defines the in-memory representation of f5acmectl's three
// manifest kinds (GatewayConfig, BigIPDevice, ManagedCertificate) and their
// shared metadata. JSON Schema files in proposal/schema/ are the source of
// truth; these structs must stay in sync.
package manifest

const APIVersion = "f5acme.example.com/v1"

const (
	KindGatewayConfig       = "GatewayConfig"
	KindBigIPDevice         = "BigIPDevice"
	KindManagedCertificate  = "ManagedCertificate"
)

type ChallengeType string

const (
	ChallengeHTTP01    ChallengeType = "http-01"
	ChallengeDNS01     ChallengeType = "dns-01"
	ChallengeTLSALPN01 ChallengeType = "tls-alpn-01"
)

type KeyType string

const (
	KeyRSA2048    KeyType = "rsa-2048"
	KeyRSA3072    KeyType = "rsa-3072"
	KeyRSA4096    KeyType = "rsa-4096"
	KeyECDSAP256  KeyType = "ecdsa-p256"
	KeyECDSAP384  KeyType = "ecdsa-p384"
)

type SecretBackend string

const (
	BackendVault   SecretBackend = "vault"
	BackendAWSSM   SecretBackend = "aws-sm"
	BackendAzureKV SecretBackend = "azure-kv"
)

type EABScope string

const (
	EABScopePerDevice      EABScope = "per-device"
	EABScopePerPartition   EABScope = "per-partition"
	EABScopePerEnvironment EABScope = "per-environment"
)

type TypeMeta struct {
	APIVersion string `yaml:"apiVersion" json:"apiVersion"`
	Kind       string `yaml:"kind"       json:"kind"`
}

type ObjectMeta struct {
	Name        string            `yaml:"name"                  json:"name"`
	Labels      map[string]string `yaml:"labels,omitempty"      json:"labels,omitempty"`
	Annotations map[string]string `yaml:"annotations,omitempty" json:"annotations,omitempty"`
}

type CredentialsRef struct {
	Backend SecretBackend `yaml:"backend" json:"backend"`
	Path    string        `yaml:"path"    json:"path"`
}

// GatewayConfig

type GatewayConfig struct {
	TypeMeta   `yaml:",inline"     json:",inline"`
	Metadata   ObjectMeta          `yaml:"metadata" json:"metadata"`
	Spec       GatewayConfigSpec   `yaml:"spec"     json:"spec"`
}

type GatewayConfigSpec struct {
	Gateway  GatewayEndpoint   `yaml:"gateway"  json:"gateway"`
	Defaults GatewayDefaults   `yaml:"defaults" json:"defaults"`
	EAB      EABPolicy         `yaml:"eab"      json:"eab"`
}

type GatewayEndpoint struct {
	DirectoryURL          string          `yaml:"directoryURL"                    json:"directoryURL"`
	AdminAPI              string          `yaml:"adminAPI"                        json:"adminAPI"`
	CABundle              string          `yaml:"caBundle,omitempty"              json:"caBundle,omitempty"`
	AdminCredentialsRef   *CredentialsRef `yaml:"adminCredentialsRef,omitempty"   json:"adminCredentialsRef,omitempty"`
	RequestTimeoutSeconds int             `yaml:"requestTimeoutSeconds,omitempty" json:"requestTimeoutSeconds,omitempty"`
}

type GatewayDefaults struct {
	KeyType         KeyType       `yaml:"keyType,omitempty"         json:"keyType,omitempty"`
	RenewBeforeDays int           `yaml:"renewBeforeDays,omitempty" json:"renewBeforeDays,omitempty"`
	Challenge       ChallengeType `yaml:"challenge,omitempty"       json:"challenge,omitempty"`
	ContactEmail    string        `yaml:"contactEmail,omitempty"    json:"contactEmail,omitempty"`
}

type EABPolicy struct {
	Scope              EABScope      `yaml:"scope"              json:"scope"`
	RotationDays       int           `yaml:"rotationDays"       json:"rotationDays"`
	SecretBackend      SecretBackend `yaml:"secretBackend"      json:"secretBackend"`
	SecretPathTemplate string        `yaml:"secretPathTemplate" json:"secretPathTemplate"`
}

// BigIPDevice

type BigIPDevice struct {
	TypeMeta `yaml:",inline"  json:",inline"`
	Metadata ObjectMeta       `yaml:"metadata" json:"metadata"`
	Spec     BigIPDeviceSpec  `yaml:"spec"     json:"spec"`
}

type BigIPDeviceSpec struct {
	Host               string              `yaml:"host"                         json:"host"`
	Port               int                 `yaml:"port,omitempty"               json:"port,omitempty"`
	MinTMOSVersion     string              `yaml:"minTmosVersion,omitempty"     json:"minTmosVersion,omitempty"`
	CredentialsRef     CredentialsRef      `yaml:"credentialsRef"               json:"credentialsRef"`
	Partitions         []string            `yaml:"partitions"                   json:"partitions"`
	ChallengeResponder *ChallengeResponder `yaml:"challengeResponder,omitempty" json:"challengeResponder,omitempty"`
	HAPeer             string              `yaml:"haPeer,omitempty"             json:"haPeer,omitempty"`
	Tags               map[string]string   `yaml:"tags,omitempty"               json:"tags,omitempty"`
}

type ChallengeResponder struct {
	Type        ChallengeType `yaml:"type"                  json:"type"`
	ListenerVS  string        `yaml:"listenerVS,omitempty"  json:"listenerVS,omitempty"`
	DNSProvider string        `yaml:"dnsProvider,omitempty" json:"dnsProvider,omitempty"`
}

// ManagedCertificate

type ManagedCertificate struct {
	TypeMeta `yaml:",inline"      json:",inline"`
	Metadata ObjectMeta            `yaml:"metadata" json:"metadata"`
	Spec     ManagedCertificateSpec `yaml:"spec"     json:"spec"`
}

type ManagedCertificateSpec struct {
	DeviceRef         string        `yaml:"deviceRef"                   json:"deviceRef"`
	Partition         string        `yaml:"partition"                   json:"partition"`
	Identifiers       []string      `yaml:"identifiers"                 json:"identifiers"`
	KeyType           KeyType       `yaml:"keyType,omitempty"           json:"keyType,omitempty"`
	Challenge         ChallengeType `yaml:"challenge,omitempty"         json:"challenge,omitempty"`
	DNSProvider       string        `yaml:"dnsProvider,omitempty"       json:"dnsProvider,omitempty"`
	RenewBeforeDays   int           `yaml:"renewBeforeDays,omitempty"   json:"renewBeforeDays,omitempty"`
	PolicyTemplate    string        `yaml:"policyTemplate,omitempty"    json:"policyTemplate,omitempty"`
	Bindings          []Binding     `yaml:"bindings"                    json:"bindings"`
	RetainOldCertDays int           `yaml:"retainOldCertDays,omitempty" json:"retainOldCertDays,omitempty"`
}

type Binding struct {
	Kind string `yaml:"kind" json:"kind"` // clientSSLProfile | serverSSLProfile
	Name string `yaml:"name" json:"name"` // /<partition>/<name>
}

// Set is the resolved, cross-referenced manifest collection for one environment.
type Set struct {
	Gateway      *GatewayConfig
	Devices      map[string]*BigIPDevice         // by metadata.name
	Certificates map[string]*ManagedCertificate  // by metadata.name
}
