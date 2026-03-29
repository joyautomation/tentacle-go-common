package common

// DeadBandConfig defines RBE (Report By Exception) thresholds for a variable.
type DeadBandConfig struct {
	Value   float64 `json:"value"`            // only publish if change exceeds this amount
	MinTime int64   `json:"minTime,omitempty"` // ms — suppress publishes more frequent than this
	MaxTime int64   `json:"maxTime,omitempty"` // ms — force publish if exceeded, 0 = disabled
}

// PlcDataMessage is published when a monitored variable changes value.
// Subject: {moduleId}.data.{deviceId}.{sanitizedVariableId}
type PlcDataMessage struct {
	ModuleID    string          `json:"moduleId"`
	DeviceID    string          `json:"deviceId"`
	VariableID  string          `json:"variableId"`
	Value       interface{}     `json:"value"`
	Timestamp   int64           `json:"timestamp"`
	Datatype    string          `json:"datatype"`
	Description string          `json:"description,omitempty"`
	Deadband    *DeadBandConfig `json:"deadband,omitempty"`
	DisableRBE  bool            `json:"disableRBE,omitempty"`
}
