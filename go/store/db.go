package store

import (
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/sqlite3"
	_ "github.com/golang-migrate/migrate/v4/source/pkger"
	"github.com/jmoiron/sqlx"
	"github.com/markbates/pkger"
	_ "github.com/mattn/go-sqlite3"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

type database struct {
	*sqlx.DB
	log *logrus.Entry
}

func newDatabase(log *logrus.Entry, filename string) (*database, error) {
	// Create sqlx database instance.
	db, err := sqlx.Open("sqlite3", filename)
	if err != nil {
		return nil, errors.Wrap(err, "failed to load database file")
	}

	// Migrate database to latest version.
	migrateDriver, err := sqlite3.WithInstance(db.DB, &sqlite3.Config{})
	if err != nil {
		return nil, errors.Wrap(err, "failed to create migration driver")
	}

	m, err := migrate.NewWithDatabaseInstance(
		"pkger://"+pkger.Include("/store/migrations"),
		"sqlite3",
		migrateDriver,
	)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create migration harness")
	}

	err = m.Up()
	if err != nil && err != migrate.ErrNoChange {
		return nil, errors.Wrap(err, "failed to migrate database up to latest")
	}

	// Success! Return driver with fully migrated database and log attached.
	return &database{
		DB:  db,
		log: log,
	}, nil
}

func NewDataStore(log *logrus.Entry, filename string) (DataStore, error) {
	return newDatabase(log, filename)
}

// Transact will begin a database transaction and execute the given transaction
// handler.
//
// Assuming no error is returned by the transaction handler, the transaction
// will be committed automatically upon return of the transaction handler.
//
// If the transaction handler panics, the transaction will rollback, then resume
// panicking up the stack.
//
// If the transaction handler returns a non-nil error value, the transaction
// will rollback and the causal error is returned.
//
// Adapted from: https://stackoverflow.com/a/23502629
func (db *database) Transact(txHandler func(tx *sqlx.Tx) error) (err error) {
	tx, err := db.Beginx()
	if err != nil {
		return errors.Wrap(err, "failed to begin transaction")
	}

	defer func() {
		if r := recover(); r != nil {
			_ = tx.Rollback()
			panic(r)
		}

		if err != nil {
			_ = tx.Rollback()
			return
		}

		tx.Commit()
	}()

	err = txHandler(tx)
	return
}
