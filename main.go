package main

import (
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
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
	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "test")
	})
	r.POST("/upload", func(c *gin.Context) {
		fileForm, _ := c.FormFile("file")
		fileUpload, _ := fileForm.Open()
		fileContent := make([]byte, 999)
		fileUpload.Read(fileContent)
		newFile, _ := os.Create("files/" + fileForm.Filename)
		newFile.Write(fileContent)
		newFile.Close()
		log.Println(fileForm.Filename)
	})
	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")
}
