package lib

import (
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"mime/multipart"
	"os"
	"strconv"
	"time"
)

func StoreToLocal(fileName string, fileSize int64, file multipart.File) error {
	newFile, e := os.Create("files/" + fileName)
	fileContent := make([]byte, fileSize)
	_, e = file.Read(fileContent)
	_, e = newFile.Write(fileContent)
	e = newFile.Close()
	return e
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
