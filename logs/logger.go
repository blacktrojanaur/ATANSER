package logs

import (
	"log"
	"os"
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
