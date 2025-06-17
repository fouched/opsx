package models

import (
	"encoding/json"
	"time"
)

type RemoteApp struct {
	ID             string          `db:"id,omitempty" schema:"id"`
	Name           string          `db:"name" schema:"name"`
	Type           string          `db:"type" schema:"type"`
	Host           string          `db:"host" schema:"host"`
	Port           int             `db:"port" schema:"port"`
	Status         string          `db:"status" schema:"status"`
	HealthURL      string          `db:"health_url" schema:"healthUrl"`
	CommandStart   string          `db:"command_start" schema:"commandStart"`
	CommandStop    string          `db:"command_stop" schema:"commandStop"`
	CommandRestart string          `db:"command_restart" schema:"commandRestart"`
	LastChecked    time.Time       `db:"last_checked" schema:"lastChecked"`
	Metadata       json.RawMessage `db:"metadata" schema:"metadata"`
}
