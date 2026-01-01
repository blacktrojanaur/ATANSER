package tunnel

import "ATANSER/logs"

func StartTunnel(route string) {
	TunnelActive = true
	CurrentRoute = route
	logs.LogInfo("Tunnel STARTED on route: " + route)
	println("Tunnel started on route:", route)
}

func StopTunnel(reason string) {
	TunnelActive = false
	logs.LogThreat("Tunnel STOPPED due to: " + reason)
	println("Tunnel stopped! Reason:", reason)
}

func RebuildTunnel(route string) {
	TunnelActive = true
	CurrentRoute = route
	logs.LogInfo("Tunnel REBUILT on route: " + route)
	println("Tunnel rebuilt on route:", route)
}

func KillTunnelOnThreat(threat string) {
	StopTunnel("Threat detected: " + threat)
}
