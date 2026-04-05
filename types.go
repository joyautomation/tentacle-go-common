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

// ─── Orchestrator Types ─────────────────────────────────────────────────────

// DesiredServiceKV is the desired state for a module.
// Stored in the desired_services KV bucket, keyed by moduleId.
type DesiredServiceKV struct {
	ModuleID  string `json:"moduleId"`
	Version   string `json:"version"`
	Running   bool   `json:"running"`
	UpdatedAt int64  `json:"updatedAt"`
}

// ServiceStatusKV is the runtime status of a module as reported by the orchestrator.
// Stored in the service_status KV bucket (TTL 120s), keyed by moduleId.
type ServiceStatusKV struct {
	ModuleID          string   `json:"moduleId"`
	InstalledVersions []string `json:"installedVersions"`
	ActiveVersion     string   `json:"activeVersion,omitempty"`
	SystemdState      string   `json:"systemdState"`
	ReconcileState    string   `json:"reconcileState"`
	LastError         string   `json:"lastError,omitempty"`
	Runtime           string   `json:"runtime"`
	Category          string   `json:"category"`
	Repo              string   `json:"repo"`
	UpdatedAt         int64    `json:"updatedAt"`
}

// ModuleRegistryInfo describes a module from the orchestrator's registry.
type ModuleRegistryInfo struct {
	ModuleID       string              `json:"moduleId"`
	Repo           string              `json:"repo"`
	Description    string              `json:"description"`
	Category       string              `json:"category"`
	Runtime        string              `json:"runtime"`
	RequiredConfig []ModuleConfigField `json:"requiredConfig,omitempty"`
}

// ModuleConfigField describes a configuration field for a module.
type ModuleConfigField struct {
	EnvVar      string `json:"envVar"`
	Description string `json:"description"`
	Default     string `json:"default,omitempty"`
	Required    bool   `json:"required"`
}

// ModuleVersionInfo holds version info for a specific module.
type ModuleVersionInfo struct {
	ModuleID          string   `json:"moduleId"`
	InstalledVersions []string `json:"installedVersions"`
	LatestVersion     string   `json:"latestVersion,omitempty"`
	ActiveVersion     string   `json:"activeVersion,omitempty"`
}

// OrchestratorCommandRequest is sent to orchestrator.command (request/reply).
type OrchestratorCommandRequest struct {
	RequestID string `json:"requestId"`
	Action    string `json:"action"`
	ModuleID  string `json:"moduleId,omitempty"`
	Timestamp int64  `json:"timestamp"`
}

// OrchestratorCommandResponse is the reply to an orchestrator command.
type OrchestratorCommandResponse struct {
	RequestID string              `json:"requestId"`
	Success   bool                `json:"success"`
	Error     string              `json:"error,omitempty"`
	Modules   []ModuleRegistryInfo `json:"modules,omitempty"`
	Online    *bool               `json:"online,omitempty"`
	Versions  *ModuleVersionInfo  `json:"versions,omitempty"`
	Timestamp int64               `json:"timestamp"`
}
