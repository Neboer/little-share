package lib

import (
	"github.com/gin-gonic/gin"
	"os"
	"syscall"
	"time"
)

func getFreeSpaceBytes() int64 {
	var stat syscall.Statfs_t
	wd, err := os.Getwd()
	if err != nil {
		_, _ = os.Stderr.WriteString(err.Error())
	}
	err = syscall.Statfs(wd, &stat)
	// Available blocks * size per block = available space in bytes
	return int64(stat.Bavail) * stat.Bsize
}

func calculateKeepTime(fileSize int64) time.Duration {
	freeSpace := getFreeSpaceBytes()
	if fileSize > freeSpace {

	}
	return time.Hour
}

// gin mid-ware for checking upload files and

func CheckUpload() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.URL.Path == "/upload" && c.Request.Method == "post" {
			//fmt.Println("good")
		} else {
			//fmt.Println(c.Request.URL.Path)
		}
	}
}
