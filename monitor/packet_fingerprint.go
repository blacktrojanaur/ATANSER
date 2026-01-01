package monitor

import (
	"ATANSER/logs"
	"time"
)

func CapturePacketFingerprint(srcIP string, size int, proto string) {
	ts := time.Now().Format(time.RFC3339)
	entry := ts + " | IP=" + srcIP + " | SIZE=" + string(rune(size)) + " | PROTO=" + proto
	logs.LogInfo("Packet fingerprint: " + entry)
	println("[FINGERPRINT]", entry)
}
