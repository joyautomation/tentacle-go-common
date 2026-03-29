package common

import "strings"

// SanitizeDeviceIDForSubject replaces spaces and NATS-invalid characters in a device ID.
func SanitizeDeviceIDForSubject(id string) string {
	r := strings.NewReplacer(" ", "_", ".", "_", "*", "_", ">", "_")
	return r.Replace(id)
}
