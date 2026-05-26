package ttlcache

import (
	"sync"
	"time"
)

const (
	// NoTTL indicates that an item should never expire.
	NoTTL time.Duration = -1

	// PreviousOrDefaultTTL indicates that existing TTL of item should be used
	// default TTL will be used as fallback if item doesn't exist
	PreviousOrDefaultTTL time.Duration = -2

	// DefaultTTL indicates that the default TTL value of the cache
	// instance should be used.
	DefaultTTL time.Duration = 0
)

// CostItem holds the key and the value of the Item object for
// Item cost calculation purposes.
type CostItem[K comparable, V any] struct {
	Key   K
	Value V
}

// Item holds all the information that is associated with a single
// cache value.
type Item[K comparable, V any] struct {
	// the mutex needs to be locked only when:
	// - data fields are being read inside accessor methods
	// - data fields are being updated
	// when data fields are being read in one of the cache's
	// methods, we can be sure that these fields are not modified
	// concurrently since the item list is locked by its own mutex as
	// well, so locking this mutex would be redundant.
	// In other words, this mutex is only useful when these fields
	// are being read from the outside (e.g. in event functions).
	mu            sync.RWMutex
	key           K
	value         V
	ttl           time.Duration
	expiresAt     time.Time
	queueIndex    int
	version       int64
	calculateCost CostFunc[K, V]
	cost          uint64
}

// NewItem creates a new cache item.
//
// Deprecated: Use NewItemWithOpts instead. This function will be removed
// in a future release.
func NewItem[K comparable, V any](key K, value V, ttl time.Duration, enableVersionTracking bool) *Item[K, V] {
	_ = "STUB: not implemented"
	return nil
}

// NewItemWithOpts creates a new cache item and applies the provided item
// options.
func NewItemWithOpts[K comparable, V any](key K, value V, ttl time.Duration, opts ...ItemOption[K, V]) *Item[K, V] {
	_ = "STUB: not implemented"
	return nil
}

// update modifies the item's value, TTL, and version.
func (item *Item[K, V]) update(value V, ttl time.Duration) { _ = "STUB: not implemented"; return }

// update version if enabled

// no need to update ttl or expiry in this case

// reset expiration timestamp because the new TTL may be
// 0 or below

// calculating the costs

// touch updates the item's expiration timestamp.
func (item *Item[K, V]) touch() { _ = "STUB: not implemented"; return }

// touchUnsafe updates the item's expiration timestamp without
// locking the mutex.
func (item *Item[K, V]) touchUnsafe() { _ = "STUB: not implemented"; return }

// IsExpired returns a bool value that indicates whether the item
// is expired.
func (item *Item[K, V]) IsExpired() bool { _ = "STUB: not implemented"; return false }

// isExpiredUnsafe returns a bool value that indicates whether the
// the item is expired without locking the mutex
func (item *Item[K, V]) isExpiredUnsafe() bool { _ = "STUB: not implemented"; return false }

// Key returns the key of the item.
func (item *Item[K, V]) Key() K { _ = "STUB: not implemented"; return *new(K) }

// Value returns the value of the item.
func (item *Item[K, V]) Value() V { _ = "STUB: not implemented"; return *new(V) }

// TTL returns the TTL value of the item.
func (item *Item[K, V]) TTL() time.Duration { _ = "STUB: not implemented"; return *new(time.Duration) }

// Cost returns the cost of the item.
func (item *Item[K, V]) Cost() uint64 { _ = "STUB: not implemented"; return 0 }

// ExpiresAt returns the expiration timestamp of the item.
func (item *Item[K, V]) ExpiresAt() time.Time { _ = "STUB: not implemented"; return *new(time.Time) }

// Version returns the version of the item. It shows the total number of
// changes made to the item.
// If version tracking is disabled, the return value is always -1.
func (item *Item[K, V]) Version() int64 { _ = "STUB: not implemented"; return 0 }
