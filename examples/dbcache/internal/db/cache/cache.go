// Package cache provides a mock implementation of a database cache.
package cache

import (
	"context"
	"dbcache/internal/order"
	"time"

	"github.com/jellydator/ttlcache/v3"
)

// Cache is an interface that defines methods for
// interacting with a database.
type Cache struct {
	db          order.DB
	volumeCache *ttlcache.Cache[string, int64]
}

// NewCache creates a new instance of the cache.
func NewCache(
	db order.DB,
	expiration time.Duration,
) *Cache {
	_ = "STUB: not implemented"
	return nil
}

// Close stops the cache and releases resources.
func (c *Cache) Close() error { _ = "STUB: not implemented"; return nil }

// FetchAssetVolume retrieves the volume for a given asset.
func (c *Cache) FetchAssetVolume(ctx context.Context, asset string) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// UpsertAssetVolume updates or inserts the volume for a given asset.
func (c *Cache) UpsertAssetVolume(ctx context.Context, asset string, volume int64) error {
	_ = "STUB: not implemented"
	return nil
}

// We set the volume after the upsert operation to ensure
// that cache is in sync with the database.
