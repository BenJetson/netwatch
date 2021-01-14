package netwatch

import (
	"time"

	"github.com/google/uuid"
)

type ResourceGroup struct {
	Name      string
	Resources []Resource
}

type ResourceType string

const (
	ResourceTypePing ResourceType = "ping"
	ResourceTypeHTTP ResourceType = "http"
)

type Resource struct {
	ID            uuid.UUID     `json:"id" db:"resource_id"`
	Name          string        `json:"name" db:"alias"`
	Active        bool          `json:"active" db:"active"`
	Type          ResourceType  `json:"type" db:"type_alias"`
	CheckInterval time.Duration `json:"check_interval" db:"check_interval"`
	NextCheck     time.Time     `json:"next_check" db:"next_check"`
}

type Status string

const (
	StatusOK    Status = "OK"
	StatusWarn  Status = "WARN"
	StatusAlarm Status = "ALARM"
)

type ResourceStatus struct {
	ResourceID uuid.UUID
	Timestamp  time.Time
	Status     Status
	Details    string
}
