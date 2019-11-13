package main

import (
	"github.com/Neboer/little-share/lib"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Use(static.Serve("/", static.LocalFile("front", true)))
	//r.Use(lib.CheckUpload())
	r.GET("/files", func(c *gin.Context) {
		FileList := lib.GetFileList()
		c.JSON(200, FileList)
	})
	r.POST("/upload", func(c *gin.Context) {
		_ = lib.StoreToLocal(c)
	})
	_ = r.Run(":8081")
}
