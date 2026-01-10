package main

import (
	"Atanser/dashboard"
	"Atanser/integrity"
	"Atanser/logs"
	"Atanser/monitor"
	"Atanser/reroute"
	"os"
)

func main() {

	logs.InitLogger()
	logs.LogInfo("ATANSER system booting up...")
	// Test packet fingerprint capture
	monitor.CapturePacketFingerprint("192.168.1.55", 512, "UDP")

	// Snapshot current route table hash (placeholder for later real hash)
	integrity.TakeSnapshot("RoutingTable", "SIM_HASH_ABC_002")

	// Test packet fingerprint capture
	monitor.CapturePacketFingerprint("192.168.1.55", 512, "UDP")

	// Snapshot current route table hash (placeholder for later real hash)
	integrity.TakeSnapshot("RoutingTable", "SIM_HASH_ABC_002")

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
	// Read threat log file
	logBytes, err := os.ReadFile("logs/ATANSER_threats.log")
	if err == nil {
		monitor.AnalyzeThreatLog(string(logBytes))
		monitor.ReportDeletionLayers(string(logBytes))
	} else {
		println("No threat logs found for analysis")
	}

	println("\nSystem active on route:", route)
}
