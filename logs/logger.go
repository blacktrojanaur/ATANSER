package logs

import (
	"log"
	"os"
	"time"
)

func InitLogger() {
	file, err := os.OpenFile("logs/ATANSER_threats.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err == nil {
		log.SetOutput(file)
	} else {
		log.Println("Failed to log to file, using default stderr")
	}
}

func LogThreat(threat string) {
	log.Println("[THREAT]", threat)
}

func LogInfo(info string) {
	log.Println("[INFO]", info)
}

// Expose log reading to GUI
func ReadThreatLogs() string {
	data, err := os.ReadFile("ATANSER_threats.log")
	if err != nil {
		return ""
	}
	return string(data)
}
func LogEvent(msg string) {
	f, _ := os.OpenFile("logs/runtime.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	defer f.Close()
	f.WriteString(time.Now().Format(time.RFC3339) + " :: " + msg + "\n")
}
