package integrity

import (
	"Atanser/logs"
	"time"
)

func TakeSnapshot(component string, value string) {
	ts := time.Now().Format(time.RFC3339)
	record := ts + " | COMPONENT=" + component + " | VALUE_HASH=" + value
	logs.LogInfo("Snapshot taken: " + record)
	println("[SNAPSHOT]", record)
}
