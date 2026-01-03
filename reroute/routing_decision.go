package reroute

import "ATANSER/logs"

func SelectSafeRoute() string {
	return "Server_B_Safe"
}

func TriggerReroute(reason string) string {
	logs.LogThreat("Reroute triggered due to: " + reason)
	return SelectSafeRoute()
}
