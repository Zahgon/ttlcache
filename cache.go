package ttlcache

import (
	"container/list"
	"context"
	"sync"
	"time"

	"golang.org/x/sync/singleflight"
)

// Available eviction reasons.
const (
	EvictionReasonDeleted EvictionReason = iota + 1
	EvictionReasonCapacityReached
	EvictionReasonExpired
	EvictionReasonMaxCostExceeded
)

// EvictionReason is used to specify why a certain item was
// evicted/deleted.
type EvictionReason int

// Cache is a synchronised map of items that are automatically removed
// when they expire or the capacity is reached.
type Cache[K comparable, V any] struct {
	items struct {
		mu     sync.RWMutex
		values map[K]*list.Element

		// a generic doubly linked list would be more convenient
		// (and more performant?). It's possible that this
		// will be introduced with/in go1.19+
		lru      *list.List
		expQueue expirationQueue[K, V]

		timerCh chan time.Duration
	}
	cost uint64

	metricsMu sync.RWMutex
	metrics   Metrics

	events struct {
		insertion struct {
			mu     sync.RWMutex
			nextID uint64
			fns    map[uint64]func(*Item[K, V])
		}
		update struct {
			mu     sync.RWMutex
			nextID uint64
			fns    map[uint64]func(*Item[K, V])
		}
		eviction struct {
			mu     sync.RWMutex
			nextID uint64
			fns    map[uint64]func(EvictionReason, *Item[K, V])
		}
	}

	stopMu  sync.Mutex
	stopCh  chan struct{}
	stopped bool

	options options[K, V]
}

// New creates a new instance of cache.
func New[K comparable, V any](opts ...Option[K, V]) *Cache[K, V] {
	_ = "STUB: not implemented"
	return nil
}

// cache cleanup process is stopped by default

// buffer is important

// updateExpirations updates the expiration queue and notifies
// the cache auto cleaner if needed.
// Not safe for concurrent use by multiple goroutines without additional
// locking.
func (c *Cache[K, V]) updateExpirations(fresh bool, elem *list.Element) {
	_ = "STUB: not implemented"
	return
}

// check if the closest/soonest expiration timestamp changed

// It's possible that the auto cleaner isn't active or
// is busy, so we need to drain the channel before
// sending a new value.
// Also, since this method is called after locking the items' mutex,
// we can be sure that there is no other concurrent call of this
// method

// we need to drain this channel in a select with a default
// case because it's possible that the auto cleaner
// read this channel just after we entered this if

// since the channel has a size 1 buffer, we can be sure
// that the line below won't block (we can't overfill the buffer
// because we just drained it)

// set creates a new item, adds it to the cache and then returns it.
// Not safe for concurrent use by multiple goroutines without additional
// locking.
func (c *Cache[K, V]) set(key K, value V, ttl time.Duration) *Item[K, V] {
	_ = "STUB: not implemented"
	return nil
}

// update/overwrite an existing item

// delete the oldest item

// create a new item

// get retrieves an item from the cache and extends its expiration
// time if 'touch' is set to true.
// It returns nil if the item is not found or is expired.
// Not safe for concurrent use by multiple goroutines without additional
// locking.
func (c *Cache[K, V]) get(key K, touch bool, includeExpired bool) *list.Element {
	_ = "STUB: not implemented"
	return nil
}

// getWithOpts wraps the get method, applies the given options, and updates
// the metrics.
// It returns nil if the item is not found or is expired.
// If 'lockAndLoad' is set to true, the mutex is locked before calling the
// get method and unlocked after it returns. It also indicates that the
// loader should be used to load external data when the get method returns
// a nil value and the mutex is unlocked.
// If 'lockAndLoad' is set to false, neither the mutex nor the loader is
// used.
func (c *Cache[K, V]) getWithOpts(key K, lockAndLoad bool, opts ...Option[K, V]) *Item[K, V] {
	_ = "STUB: not implemented"
	return nil
}

// evict deletes items from the cache.
// If no items are provided, all currently present cache items
// are evicted.
// Not safe for concurrent use by multiple goroutines without additional
// locking.
func (c *Cache[K, V]) evict(reason EvictionReason, elems ...*list.Element) {
	_ = "STUB: not implemented"
	return
}

// delete deletes an item by the provided key.
// The method is no-op if the item is not found.
// Not safe for concurrent use by multiple goroutines without additional
// locking.
func (c *Cache[K, V]) delete(key K) { _ = "STUB: not implemented"; return }

// Set creates a new item from the provided key and value, adds
// it to the cache and then returns it. If an item associated with the
// provided key already exists, the new item overwrites the existing one.
// NoTTL constant or -1 can be used to indicate that the item should never
// expire.
// DefaultTTL constant or 0 can be used to indicate that the item should use
// the default/global TTL that was specified when the cache instance was
// created.
func (c *Cache[K, V]) Set(key K, value V, ttl time.Duration) *Item[K, V] {
	_ = "STUB: not implemented"
	return nil
}

// Get retrieves an item from the cache by the provided key.
// Unless this is disabled, it also extends/touches an item's
// expiration timestamp on successful retrieval.
// If the item is not found, a nil value is returned.
func (c *Cache[K, V]) Get(key K, opts ...Option[K, V]) *Item[K, V] {
	_ = "STUB: not implemented"
	return nil
}

// Delete deletes an item from the cache. If the item associated with
// the key is not found, the method is no-op.
func (c *Cache[K, V]) Delete(key K) { _ = "STUB: not implemented"; return }

// Has checks whether the key exists in the cache.
func (c *Cache[K, V]) Has(key K) bool { _ = "STUB: not implemented"; return false }

// GetOrSet retrieves an item from the cache by the provided key.
// If the item is not found, it is created with the provided options and
// then returned.
// The bool return value is true if the item was found, false if created
// during the execution of the method.
// If the loader is non-nil (i.e., used as an option or specified when
// creating the cache instance), its execution is skipped.
func (c *Cache[K, V]) GetOrSet(key K, value V, opts ...Option[K, V]) (*Item[K, V], bool) {
	_ = "STUB: not implemented"
	return nil, false
}

// GetOrSetFunc retrieves an item from the cache by the provided key.
// If the element is not found, it is created by executing the fn function
// with the provided options and then returned.
// The bool return value is true if the item was found, false if created
// during the execution of the method.
// If the loader is non-nil (i.e., used as an option or specified when
// creating the cache instance), its execution is skipped.
func (c *Cache[K, V]) GetOrSetFunc(key K, fn func() V, opts ...Option[K, V]) (*Item[K, V], bool) {
	_ = "STUB: not implemented"
	return nil, false
}

// used only to update the TTL

// GetAndDelete retrieves an item from the cache by the provided key and
// then deletes it.
// The bool return value is true if the item was found before
// its deletion, false if not.
// If the loader is non-nil (i.e., used as an option or specified when
// creating the cache instance), it is executed normaly, i.e., only when
// the item is not found.
func (c *Cache[K, V]) GetAndDelete(key K, opts ...Option[K, V]) (*Item[K, V], bool) {
	_ = "STUB: not implemented"
	return nil, false
}

// used only to update the loader

// DeleteAll deletes all items from the cache.
func (c *Cache[K, V]) DeleteAll() { _ = "STUB: not implemented"; return }

// DeleteExpired deletes all expired items from the cache.
func (c *Cache[K, V]) DeleteExpired() { _ = "STUB: not implemented"; return }

// expiration queue has a new root

// Touch simulates an item's retrieval without actually returning it.
// Its main purpose is to extend an item's expiration timestamp.
// If the item is not found, the method is no-op.
func (c *Cache[K, V]) Touch(key K) { _ = "STUB: not implemented"; return }

// Len returns the number of unexpired items in the cache.
func (c *Cache[K, V]) Len() int { _ = "STUB: not implemented"; return 0 }

// search the heap-based expQueue by BFS

// Keys returns all unexpired keys in the cache.
func (c *Cache[K, V]) Keys() []K { _ = "STUB: not implemented"; return nil }

// Items returns a copy of all items in the cache.
// It does not update any expiration timestamps.
func (c *Cache[K, V]) Items() map[K]*Item[K, V] { _ = "STUB: not implemented"; return nil }

// Range calls fn for each unexpired item in the cache. If fn returns false,
// Range stops the iteration.
func (c *Cache[K, V]) Range(fn func(item *Item[K, V]) bool) {
	_ = "STUB: not implemented"

	// Check if cache is empty
	return
}

// unlock mutex so fn func can access it (if it needs to)

// RangeBackwards calls fn for each unexpired item in the cache in reverse order.
// If fn returns false, RangeBackwards stops the iteration.
func (c *Cache[K, V]) RangeBackwards(fn func(item *Item[K, V]) bool) {
	_ = "STUB: not implemented"

	// Check if cache is empty
	return
}

// unlock mutex so fn func can access it (if it needs to)

// Metrics returns the metrics of the cache.
func (c *Cache[K, V]) Metrics() Metrics { _ = "STUB: not implemented"; return *new(Metrics) }

// Start starts an automatic cleanup process that periodically deletes
// expired items.
// It blocks until Stop is called.
func (c *Cache[K, V]) Start() { _ = "STUB: not implemented"; return }

// execute immediately

// drain the timer chan

// Stop stops the automatic cleanup process.
// It blocks until the cleanup process exits.
func (c *Cache[K, V]) Stop() { _ = "STUB: not implemented"; return }

// OnInsertion adds the provided function to be executed when
// a new item is inserted into the cache. The function is executed
// on a separate goroutine and does not block the flow of the cache
// manager.
// The returned function may be called to delete the subscription function
// from the list of insertion subscribers.
// When the returned function is called, it blocks until all instances of
// the same subscription function return. A context is used to notify the
// subscription function when the returned/deletion function is called.
func (c *Cache[K, V]) OnInsertion(fn func(context.Context, *Item[K, V])) func() {
	_ = "STUB: not implemented"
	return nil
}

// OnUpdate adds the provided function to be executed when
// an item is updated in the cache. The function is executed
// on a separate goroutine and does not block the flow of the cache
// manager.
// The returned function may be called to delete the subscription function
// from the list of update subscribers.
// When the returned function is called, it blocks until all instances of
// the same subscription function return. A context is used to notify the
// subscription function when the returned/deletion function is called.
func (c *Cache[K, V]) OnUpdate(fn func(context.Context, *Item[K, V])) func() {
	_ = "STUB: not implemented"
	return nil
}

// OnEviction adds the provided function to be executed when
// an item is evicted/deleted from the cache. The function is executed
// on a separate goroutine and does not block the flow of the cache
// manager.
// The returned function may be called to delete the subscription function
// from the list of eviction subscribers.
// When the returned function is called, it blocks until all instances of
// the same subscription function return. A context is used to notify the
// subscription function when the returned/deletion function is called.
func (c *Cache[K, V]) OnEviction(fn func(context.Context, EvictionReason, *Item[K, V])) func() {
	_ = "STUB: not implemented"
	return nil
}

// Loader is an interface that handles missing data loading.
type Loader[K comparable, V any] interface {
	// Load should execute a custom item retrieval logic and
	// return the item that is associated with the key.
	// It should return nil if the item is not found/valid.
	// The method is allowed to fetch data from the cache instance
	// or update it for future use.
	Load(c *Cache[K, V], key K) *Item[K, V]
}

// LoaderFunc type is an adapter that allows the use of ordinary
// functions as data loaders.
type LoaderFunc[K comparable, V any] func(*Cache[K, V], K) *Item[K, V]

// Load executes a custom item retrieval logic and returns the item that
// is associated with the key.
// It returns nil if the item is not found/valid.
func (l LoaderFunc[K, V]) Load(c *Cache[K, V], key K) *Item[K, V] {
	_ = "STUB: not implemented"

	// SuppressedLoader wraps another Loader and suppresses duplicate
	// calls to its Load method.
	return nil
}

type SuppressedLoader[K comparable, V any] struct {
	loader Loader[K, V]
	group  *singleflight.Group
}

// NewSuppressedLoader creates a new instance of suppressed loader.
// If the group parameter is nil, a newly created instance of
// *singleflight.Group is used.
func NewSuppressedLoader[K comparable, V any](loader Loader[K, V], group *singleflight.Group) *SuppressedLoader[K, V] {
	_ = "STUB: not implemented"
	return nil
}

// Load executes a custom item retrieval logic and returns the item that
// is associated with the key.
// It returns nil if the item is not found/valid.
// It also ensures that only one execution of the wrapped Loader's Load
// method is in-flight for a given key at a time.
func (l *SuppressedLoader[K, V]) Load(c *Cache[K, V], key K) *Item[K, V] {
	_ = "STUB: not implemented"
	// there should be a better/generic way to create a
	// singleflight Group's key. It's possible that a generic
	// singleflight.Group will be introduced with/in go1.19+
	return nil
}

// the error can be discarded since the singleflight.Group
// itself does not return any of its errors, it returns
// the error that we return ourselves in the func below, which
// is also nil
