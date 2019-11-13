package lib

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"strconv"
	"time"
)

func StoreToLocal(c *gin.Context) error {
	defer func() {
		err := recover()
		if err != nil {
			log.Println(err)
		}
	}()
	_, header, err := c.Request.FormFile("file")
	if err != nil {
		c.AbortWithStatusJSON(499, gin.H{"error": "no 'file' in request body."})
		return err
	}
	filename := header.Filename
	fmt.Println(header.Filename)
	err = c.SaveUploadedFile(header, "files/"+filename)
	return err
}

func removeOutOfDateFiles() {

}

func KeepTimeChecker() {
	currentHour := time.Now().Hour()
	for {
		if time.Now().Hour() != currentHour {
			currentHour = time.Now().Hour()
		}
	}
}

func ReadMaxSpaceUsage(c gin.Context) {
	fileContent, _ := ioutil.ReadFile("maxSpaceUsage")
	maxSpaceUsageString := string(fileContent)
	maxSpaceUsage, _ := strconv.Atoi(maxSpaceUsageString)
	c.Keys["maxSpaceUsage"] = maxSpaceUsage
}
