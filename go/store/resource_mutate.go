package store

import (
	"context"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"

	"github.com/BenJetson/netwatch"
)

func (db *database) CreateResource(ctx context.Context, r netwatch.Resource) error {
	var err error
	if r.ID, err = uuid.NewUUID(); err != nil {
		return errors.Wrap(err, "failed to create new resource ID")
	}

	return db.Transact(func(tx *sqlx.Tx) error {
		result, err := tx.ExecContext(ctx, `
			INSERT INTO resource (
				resource_id,
				alias,
				resource_type_id
			) VALUES (?, ?, (
				SELECT resource_type_id
					FROM resource_type
					WHERE alias = ?
			))`, r.ID, r.Name, r.Type)

		if err != nil {
			return errors.Wrap(err, "failed to insert resource")
		}

		n, err := result.RowsAffected()
		if err != nil {
			return errors.Wrap(err, "failed to determine result of insert")
		} else if n != 1 {
			return errors.Wrap(err, "insert ought to touch one row only")
		}

		return nil
	})
}
func (db *database) UpdateResource(ctx context.Context, r netwatch.Resource) error {
	return nil
}
func (db *database) DeleteResource(ctx context.Context, id uuid.UUID) error {
	return nil
}
