package monitor

import (
	"ATANSER/integrity"
	"ATANSER/logs"
)

func DetectBGPChange(currentHash string) {
	if currentHash != integrity.RoutingTableHash {
		logs.LogThreat("Routing table mismatch! Possible BGP hijacking")
	}
}
