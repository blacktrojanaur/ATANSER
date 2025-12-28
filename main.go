package main

import (
	"ATANSER/dashboard"
	"ATANSER/integrity"
	"ATANSER/logs"
	"ATANSER/monitor"
	"ATANSER/reroute"
)

func main() {
	logs.InitLogger()
	logs.LogInfo("ATANSER system booting up...")

	// Trigger threat simulation
	monitor.RunSimulationMenu()

	// Choose a safe route
	route := reroute.SelectSafeRoute()

	// Store integrity hash record
	integrity.StoreHashRecord("RoutingTable", "SIMULATED_HASH_001")

	// Launch CLI dashboard
	dashboard.ShowDashboard()

	println("\nSystem active on route:", route)
}
