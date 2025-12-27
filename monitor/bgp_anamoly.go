package monitor

import "ATANSER/logs"

func DetectBGPChange(currentHash string, trustedHash string) {
	if trustedHash == "" {
		trustedHash = currentHash
		logs.LogInfo("Stored Routing Table fingerprint hash: " + trustedHash)
	}

	if currentHash != trustedHash {
		logs.LogThreat("Possible BGP Hijacking — Routing table changed unexpectedly!")
	}
}
