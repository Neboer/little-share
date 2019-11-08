package main

import (
	"./lib"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"log"
	"os"
)

func main() {
	r := gin.Default()
	r.Use(static.Serve("/", static.LocalFile("front", true)))
	r.GET("/files", func(c *gin.Context) {
		f, _ := os.Open("files")
		i, _ := f.Readdir(-1)
		for _, fi := range i {
			log.Println(fi.Name())
		}
	})
	r.POST("/upload", func(c *gin.Context) {
		fileForm, _ := c.FormFile("file")
		fileUpload, _ := fileForm.Open()
		_ = lib.StoreToLocal(fileForm.Filename, fileForm.Size, fileUpload)
		log.Println(fileForm.Filename)
	})
	// Listen and Server in 0.0.0.0:8080
	_ = r.Run(":8080")
}
