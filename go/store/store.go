package store

import (
	"context"
	"io"

	"github.com/google/uuid"

	"github.com/BenJetson/netwatch"
)

type DataStore interface {
	io.Closer

	CreateResource(ctx context.Context, r netwatch.Resource) error
	UpdateResource(ctx context.Context, r netwatch.Resource) error
	DeleteResource(ctx context.Context, id uuid.UUID) error

	GetResources(ctx context.Context) ([]netwatch.Resource, error)
	GetResourcsOfType(ctx context.Context, t netwatch.ResourceType) ([]netwatch.Resource, error)
}
