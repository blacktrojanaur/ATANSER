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

	// Test simulation
	monitor.SimulateThreat("ARP_SPOOFING")
	route := reroute.SelectSafeRoute()
	integrity.StoreHashRecord("RoutingTable", "FAKE_HASH_001")
	dashboard.ShowDashboard()

	println("System active on route:", route)
}
