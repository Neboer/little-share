package main

import (
	"github.com/Neboer/little-share/lib"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"strconv"
)

func main() {
	r := gin.Default()
	MaxSpaceUsage := lib.ReadMaxSpaceUsage()
	r.Use(static.Serve("/", static.LocalFile("front", true)))
	r.Use(static.Serve("/files", static.LocalFile("files", false)))
	r.GET("/maxSpace", func(c *gin.Context) {
		c.String(200,strconv.Itoa(int(MaxSpaceUsage-lib.GetCurrentTotalFileSize())))
	})
	r.GET("/files", func(c *gin.Context) {
		FileList := lib.GetFileList(MaxSpaceUsage)
		c.JSON(200, FileList)
	})
	r.POST("/upload", func(c *gin.Context) {
		_ = lib.StoreToLocal(c)
	})
	_ = r.Run(":8081")
}
