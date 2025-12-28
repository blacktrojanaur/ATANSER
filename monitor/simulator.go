package monitor

import "ATANSER/logs"

func RunSimulationMenu() {
	logs.LogInfo("Threat simulation menu launched")
	println("\n--- THREAT SIMULATION MODE ---")
	println("Simulating: ARP Spoofing")
	println("Simulating: BGP Hijack")
	println("Simulating: MITM Attack")
	logs.LogThreat("Simulated ARP Spoofing detected")
	logs.LogThreat("Simulated BGP Hijacking detected")
	logs.LogThreat("Simulated MITM detected")
}
