package store

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/BenJetson/netwatch"
)

func TestCreateResource(t *testing.T) {
	db, dbID := newTestDB(t)
	defer destroyTestDB(t, db, dbID)

	ctx := context.Background()

	r := netwatch.Resource{
		Name: "IP Printing Service",
		Type: netwatch.ResourceTypeHTTP,
	}

	err := db.CreateResource(ctx, r)
	require.NoError(t, err)

	res, err := db.GetResources(ctx)
	require.NoError(t, err, "fetching resources ought not fail")
	require.Len(t, res, 1, "one resource in ought to yield resource out")
	assert.NotZero(t, res[0].ID, "create resource ought to assign identifier")
	assert.Equal(t, r.Name, res[0].Name, "name in ought to equal name out")
	assert.Equal(t, r.Type, res[0].Type, "type in ought to equal type out")
}
