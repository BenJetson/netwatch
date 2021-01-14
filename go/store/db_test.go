package store

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const testDBFileFormat = "/tmp/netwatch-test-%s.db"

func destroyTestDB(t *testing.T, db *database, id uuid.UUID) {
	err := db.Close()
	assert.NoError(t, err, "closing DB failed")

	testDBFile := fmt.Sprintf(testDBFileFormat, id.String())
	err = os.Remove(testDBFile)
	assert.NoError(t, err, "destroying DB file failed")
}

func newTestDB(t *testing.T) (db *database, dbID uuid.UUID) {
	var err error

	dbID, err = uuid.NewUUID()
	require.NoError(t, err, "cannot proceed with test without test DB ID")

	testDBFile := fmt.Sprintf(testDBFileFormat, dbID.String())

	log := logrus.New()
	log.SetOutput(ioutil.Discard)
	testLog := log.WithField("source", "test")

	db, err = newDatabase(testLog, testDBFile)
	require.NoError(t, err, "cannot proceed with test without test database")
	return
}
