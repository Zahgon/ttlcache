// Package order provides functionality for managing orders.
package order

import (
	"context"
	"io"
	"log/slog"
)

// Manager is responsible for managing orders,
// consuming them from a stream, and updating the
// asset volumes in the database.
type Manager struct {
	log      *slog.Logger
	db       DB
	streamer Streamer
}

// NewManager creates a new order manager instance.
func NewManager(streamer Streamer, db DB) *Manager { _ = "STUB: not implemented"; return nil }

// Run starts the order manager, consuming orders from the streamer
// and updating the asset volumes in the database.
func (m *Manager) Run(ctx context.Context) { _ = "STUB: not implemented"; return }

// Streamer is an interface that defines methods for
// consuming orders from a stream.
type Streamer interface {
	// Consume should return a channel that streams orders.
	Consume(ctx context.Context) <-chan Order
}

// DB is an interface that defines methods for
// interacting with a database
type DB interface {
	io.Closer

	// FetchAssetVolume should retrieve the volume
	// for a given asset.
	FetchAssetVolume(ctx context.Context, asset string) (int64, error)

	// UpsertAssetVolume should update or insert the volume
	// for a given asset.
	UpsertAssetVolume(ctx context.Context, asset string, volume int64) error
}

// Order represents an order in the system.
type Order struct {
	// Asset is the identifier for the asset.
	Asset string `json:"asset"`

	// Quantity is the amount of the asset in the order.
	Quantity int64 `json:"quantity"`
}
