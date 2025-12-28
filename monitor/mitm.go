package monitor

import "ATANSER/logs"

func SimulateThreat(threat string) {
	logs.LogThreat("Simulated threat triggered: " + threat)
}
