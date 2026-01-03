package tunnel

import "Atanser/logs"

func RebuildTunnel(route string) {
	logs.LogInfo("Tunnel REBUILT on fallback route: " + route)
	println("Tunnel rebuilt on route:", route)
	TunnelActive = true
	CurrentRoute = route
}
