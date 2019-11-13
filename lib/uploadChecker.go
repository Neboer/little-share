package lib

import (
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

func calculateKeepTime(fileSize int64) time.Duration {
	return time.Hour
}

// gin mid-ware for checking upload files and

func CheckUpload() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.URL.Path == "/upload" && c.Request.Method == "POST" {
			log.Println("good")
		} else {
			//fmt.Println(c.Request.URL.Path)
		}
	}
}
