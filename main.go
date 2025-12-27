package main

import (
	"ATANSER/logs"
	"ATANSER/monitor"
)

func main() {
	logs.InitLogger()
	logs.LogInfo("ATANSER system booting up...")

	// Simulated fingerprints for now (you replace with real ones later)
	fakeMAC := "AA:BB:CC:DD:EE:11"
	trustedMAC := "AA:BB:CC:DD:EE:22"

	fakeRouteHash := "HASH_1234"
	trustedRouteHash := "HASH_9999"

	// Run detectors
	monitor.DetectARPSpoof(fakeMAC, trustedMAC)
	monitor.DetectBGPChange(fakeRouteHash, trustedRouteHash)

	println("ATANSER Phase 3 modules running...")
}
