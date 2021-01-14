package store

import (
	"context"

	"github.com/pkg/errors"

	"github.com/BenJetson/netwatch"
)

func (db *database) GetResources(ctx context.Context) ([]netwatch.Resource, error) {
	var res []netwatch.Resource

	err := db.SelectContext(ctx, &res, `
		SELECT
			r.resource_id,
			r.alias,
			t.alias AS type_alias
		FROM resource AS r
		JOIN resource_type AS t
			ON r.resource_type_id = t.resource_type_id`)

	return res, errors.Wrap(err, "failed to fetch resources")
}

func (db *database) GetResourcsOfType(ctx context.Context, t netwatch.ResourceType) ([]netwatch.Resource, error) {
	return nil, nil
}
