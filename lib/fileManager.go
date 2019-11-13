package lib

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"time"
)

func StoreToLocal(c *gin.Context) error {
	defer func() {
		err := recover()
		if err != nil {
			if err == io.ErrUnexpectedEOF {
				log.Println("uploader terminate the connection before it complete.")
			}
			log.Println(err)
		}
	}()
	_, header, err := c.Request.FormFile("file")
	if err != nil {
		c.AbortWithStatusJSON(400, gin.H{"error": "no 'file' in request body."})
		log.Println(err)
		return err
	}
	filename := header.Filename
	fmt.Println(header.Filename)
	err = c.SaveUploadedFile(header, "files/"+filename)
	return err
}

type FileData struct {
	FileName            string
	FileSizeBytes       int64
	FileSurplusKeepTime time.Duration
}

func GetFileList() []FileData {
	var FileList []FileData
	f, _ := os.Open("files")
	i, _ := f.Readdir(-1)
	for _, fi := range i {
		fd := FileData{fi.Name(), fi.Size(), time.Now().Sub(fi.ModTime())}
		FileList = append(FileList,fd)
	}
	return FileList
}

func removeOutOfDateFiles() {

}

func TotalKeepTime(fileSizeBytes int64) {
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
