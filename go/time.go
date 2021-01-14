package netwatch

import (
	"database/sql/driver"
	"time"

	"github.com/pkg/errors"
)

// Time is a wrapper around time.Time that enables serialization for storage
// by and retrieval from a netwatch.Store.
type Time struct {
	time.Time
}

func (t *Time) Value() (driver.Value, error) {
	return t.Format(time.RFC3339), nil
}

func (t *Time) Scan(value interface{}) error {
	s, ok := value.(string)
	if !ok {
		return errors.New("cannot scan non-string value as Time")
	}

	var err error
	t.Time, err = time.Parse(time.RFC3339, s)
	return errors.Wrap(err, "could not parse as RFC3339 timestamp")
}

// Duration is a wrapper around time.Duration that enables serialization for
// storage by and retrieval from a netwatch.Store.
type Duration struct {
	time.Duration
}

func (d *Duration) Value() (driver.Value, error) {
	return int(d.Duration), nil
}

func (d *Duration) Scan(value interface{}) error {
	nsec, ok := value.(int)
	if !ok {
		return errors.New("cannot scan non-integer value as Duration")
	}

	d.Duration = time.Duration(nsec)
	return nil
}

var x = Duration{}
var y = x.Round()
