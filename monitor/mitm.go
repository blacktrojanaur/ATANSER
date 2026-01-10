package monitor

import "Atanser/logs"

func SimulateThreat(threat string) {
	logs.LogThreat("Simulated threat triggered: " + threat)
}
