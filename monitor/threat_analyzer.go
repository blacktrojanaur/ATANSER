package monitor

import (
	"Atanser/logs"
	"strings"
)

// Analyze captured threat logs for severity & deletion behavior
func AnalyzeThreatLog(data string) {
	if strings.Contains(data, "spoof") || strings.Contains(data, "hijack") {
		logs.LogThreat("High severity threat confirmed from analysis")
		println("[ANALYZER] High severity network threat detected!")
	} else if strings.TrimSpace(data) == "" {
		logs.LogInfo("Possible layer 1 evidence deletion detected")
		println("[ANALYZER] Warning: Evidence might have been deleted!")
	} else {
		logs.LogInfo("Threat log analyzed: No critical flags")
		println("[ANALYZER] Threat scan completed — No critical threat!")
	}
}

// Detect layers of deletion from log patterns
func DetectDeletionLayers(logData string) int {
	layers := 0
	if strings.Contains(logData, "cipher wipe") {
		layers++
	}
	if strings.Contains(logData, "rm -rf") || strings.Contains(logData, "secure delete") {
		layers++
	}
	if strings.Contains(logData, "disk zero") || strings.Contains(logData, "block overwrite") {
		layers++
	}
	return layers
}

// Wrapper for reporting deletion layers
func ReportDeletionLayers(logData string) {
	l := DetectDeletionLayers(logData)
	if l > 0 {
		logs.LogThreat("Evidence deletion layers detected: " + string(rune(l)))
		println("[ANALYZER] Evidence deletion layers detected:", l)
	} else {
		logs.LogInfo("No evidence deletion layers found")
		println("[ANALYZER] No evidence deletion behavior found")
	}
}
