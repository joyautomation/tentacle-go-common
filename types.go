// Package common provides shared NATS message types for tentacle Go modules.
//
// These types match the canonical schema defined in tentacle-proto and must
// produce identical JSON when marshaled with encoding/json. Any changes here
// should be reflected in the corresponding .proto files.
package common

// ServiceHeartbeat is published every 10s to the service_heartbeats KV bucket.
type ServiceHeartbeat struct {
	ServiceType string                 `json:"serviceType"`
	ModuleID    string                 `json:"moduleId"`
	LastSeen    int64                  `json:"lastSeen"`
	StartedAt   int64                  `json:"startedAt"`
	Version     string                 `json:"version,omitempty"`
	Metadata    map[string]interface{} `json:"metadata,omitempty"`
}

// ServiceEnabledKV is the value stored in the service_enabled KV bucket.
type ServiceEnabledKV struct {
	ModuleID  string `json:"moduleId"`
	Enabled   bool   `json:"enabled"`
	UpdatedAt int64  `json:"updatedAt"`
}

// ServiceLogEntry is published to service.logs.{serviceType}.{moduleId}.
type ServiceLogEntry struct {
	Timestamp   int64  `json:"timestamp"`
	Level       string `json:"level"`
	Message     string `json:"message"`
	ServiceType string `json:"serviceType"`
	ModuleID    string `json:"moduleId"`
	Logger      string `json:"logger,omitempty"`
}

// BrowseProgressMessage is published during async browse operations
// to {protocol}.browse.progress.{browseId}.
type BrowseProgressMessage struct {
	BrowseID      string `json:"browseId"`
	ModuleID      string `json:"moduleId"`
	DeviceID      string `json:"deviceId"`
	Phase         string `json:"phase"` // "discovering", "expanding", "reading", "caching", "completed", "failed"
	TotalTags     int    `json:"totalTags"`
	CompletedTags int    `json:"completedTags"`
	ErrorCount    int    `json:"errorCount"`
	Message       string `json:"message,omitempty"`
	Timestamp     int64  `json:"timestamp"`
}
