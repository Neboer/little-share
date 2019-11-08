package main

import (
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/neboer/little-share/lib"
	"log"
	"os"
)

func main() {
	r := gin.Default()

	// if Allow DirectoryIndex
	//r.Use(static.Serve("/", static.LocalFile("/tmp", true)))
	// set prefix
	//r.Use(static.Serve("/static", static.LocalFile("/tmp", true)))

	r.Use(static.Serve("/", static.LocalFile("front", true)))
	r.GET("/files", func(c *gin.Context) {
		f, _ := os.Open("files")
		i, _ := f.Readdir(-1)
		for _, fi := range i{
			log.Println(fi.Name())
		}
	})
	r.POST("/upload", func(c *gin.Context) {
		fileForm, _ := c.FormFile("file")
		fileUpload, _ := fileForm.Open()
		log.Println(fileForm.Filename)
	})
	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")
}
