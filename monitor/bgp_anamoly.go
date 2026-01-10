package monitor

import (
	"Atanser/integrity"
	"Atanser/logs"
)

func DetectBGPChange(currentHash string) {
	if currentHash != integrity.RoutingTableHash {
		logs.LogThreat("Routing table mismatch! Possible BGP hijacking")
	}
}
