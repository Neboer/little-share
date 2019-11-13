package main

import (
	"github.com/Neboer/little-share/lib"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"log"
	"os"
)

func main() {
	r := gin.Default()
	r.Use(static.Serve("/", static.LocalFile("front", true)))
	r.Use(lib.CheckUpload())
	r.GET("/files", func(c *gin.Context) {
		f, _ := os.Open("files")
		i, _ := f.Readdir(-1)
		for _, fi := range i {
			log.Println(fi.Name())
		}
	})
	r.POST("/upload", func(c *gin.Context) {
		_ = lib.StoreToLocal(c)
	})
	// Listen and Server in 0.0.0.0:8080
	_ = r.Run(":8081")
}
