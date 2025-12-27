package monitor

import (
	"ATANSER/logs"
	"log"
)

func DetectARPSpoof(currentMAC string, trustedMAC string) {
	if trustedMAC == "" {
		trustedMAC = currentMAC
		logs.LogInfo("Stored trusted Gateway MAC fingerprint: " + trustedMAC)
	}

	if currentMAC != trustedMAC {
		log.Println("ARP spoof detected!")
		logs.LogThreat("ARP spoofing attempt detected! TrustedMAC=" + trustedMAC + " CurrentMAC=" + currentMAC)
	}
}
