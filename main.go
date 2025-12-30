package main

import (
	"ATANSER/dashboard"
	"ATANSER/integrity"
	"ATANSER/logs"
	"ATANSER/monitor"
	"ATANSER/reroute"
	"os"
)

func main() {
	logs.InitLogger()
	logs.LogInfo("ATANSER system booting up...")

	// Run threat simulation menu
	monitor.RunSimulationMenu()

	// Choose a safe route
	route := reroute.SelectSafeRoute()

	// Read real routing table hash snapshot
	hashBytes, err := os.ReadFile("logs/routing_table.txt")
	if err == nil {
		integrity.StoreHashRecord("RoutingTableSnapshot", string(hashBytes))
	} else {
		logs.LogThreat("Routing table snapshot not found or unreadable")
	}

	// Launch CLI dashboard
	dashboard.ShowDashboard()

	println("\nSystem active on route:", route)
}
