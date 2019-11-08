package lib

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

// gin mid-ware for checking upload files and
func checkUpload(c *gin.Context) gin.HandlerFunc {
	return func(context *gin.Context) {
		if c.Request.URL.Path == "/upload" {
			fmt.Println("good")
		}
	}
}
