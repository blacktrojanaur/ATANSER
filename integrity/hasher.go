package integrity

import "Atanser/logs"

func StoreHashRecord(entity string, hash string) {
	logs.LogInfo("Hash stored for " + entity + " : " + hash)
}
