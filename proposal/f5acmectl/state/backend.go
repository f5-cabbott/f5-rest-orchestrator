// Package state defines the locked, versioned state-snapshot backend used by
// `f5acmectl apply` and `reconcile`. Implementations should map to S3, GCS,
// Azure Blob, or a local file (for lab use). State NEVER contains secret
// material; only secret references.
package state

import (
	"context"
	"time"
)

type Backend interface {
	// Lock acquires an exclusive lock for the duration of an apply.
	// Implementations should support a TTL so a crashed apply can be recovered.
	Lock(ctx context.Context, holder string, ttl time.Duration) (Lock, error)

	// LatestSnapshot returns the most recent snapshot, or nil with no error
	// if the backend is empty.
	LatestSnapshot(ctx context.Context) (*Snapshot, error)

	// ListSnapshots returns metadata for the most recent N snapshots,
	// newest first.
	ListSnapshots(ctx context.Context, limit int) ([]SnapshotMeta, error)

	// PutSnapshot writes a new snapshot atomically and returns its ID.
	PutSnapshot(ctx context.Context, snap *Snapshot) (string, error)

	// GetSnapshot retrieves a specific snapshot by ID.
	GetSnapshot(ctx context.Context, id string) (*Snapshot, error)
}

type Lock interface {
	// Refresh extends the lock TTL.
	Refresh(ctx context.Context) error
	// Release releases the lock. Idempotent.
	Release(ctx context.Context) error
}

type SnapshotMeta struct {
	ID        string
	CreatedAt time.Time
	Actor     string
	GitSHA    string
	Summary   string // short human-readable change summary
}

// Snapshot captures what was applied. Concrete shape is defined by the
// applier; we keep it opaque here so the backend stays storage-only.
type Snapshot struct {
	Meta SnapshotMeta
	Body []byte // serialized state document (e.g. JSON)
}
