package reroute

import "ATANSER/logs"

func SelectSafeRoute() string {
	safe := "Fallback_Server_B"
	logs.LogInfo("Safe route selected: " + safe)
	return safe
}
func TriggerReroute(reason string) string {
	logs.LogThreat("Reroute triggered due to: " + reason)
	return SelectSafeRoute()
}
