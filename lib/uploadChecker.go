package lib

import (
	"time"
)

func calculateKeepTime(fileSize int64) time.Duration {
	return time.Hour
}

// gin mid-ware for checking upload files and
