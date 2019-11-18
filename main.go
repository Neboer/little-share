package main

import (
	"github.com/Neboer/little-share/lib"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"strconv"
)

func main() {
	r := gin.Default()
	maxKeepTimeDbJsonList := lib.ReadKeepTimeDB()
	MaxSpaceUsage := lib.ReadMaxSpaceUsage()
	r.Use(static.Serve("/", static.LocalFile("front", true)))
	r.GET("/maxSpace", func(c *gin.Context) {
		c.String(200, strconv.Itoa(int(MaxSpaceUsage-lib.GetCurrentTotalFileSize())))
	})
	r.GET("/files", func(c *gin.Context) {
		FileList := lib.GetFileList(&maxKeepTimeDbJsonList)
		c.JSON(200, FileList)
	})
	r.DELETE("/files/:file", func(c *gin.Context) {
		FileNameNeedToDelete := c.Param("file")
		fileErr := lib.DeleteOneFile(FileNameNeedToDelete, &maxKeepTimeDbJsonList)
		if fileErr != lib.OperationSuccessful {
			if fileErr == lib.NoSuchFileError {
				c.String(404, "no such file: %s", FileNameNeedToDelete)
			} else {
				c.String(500, "internal server error.")
			}
		} else {
			c.String(200, "file %s has been delete.", FileNameNeedToDelete)
		}
	})
	r.Use(static.Serve("/files", static.LocalFile("files", false)))
	r.POST("/upload", func(c *gin.Context) {
		_ = lib.StoreToLocal(c, MaxSpaceUsage, &maxKeepTimeDbJsonList)
	})
	go lib.CheckAndDelete(&maxKeepTimeDbJsonList)
	_ = r.Run(":8081")

}
