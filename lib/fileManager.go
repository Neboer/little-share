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

func GetStoredFilesFolder() []os.FileInfo {
	f, _ := os.Open("files")
	i, _ := f.Readdir(-1)
	return i
}

func calculateTotalFileSize(i []os.FileInfo) int64 {
	var CurrentTotalFileSize int64 = 0
	for _, fi := range i {
		CurrentTotalFileSize += fi.Size()
	}
	return CurrentTotalFileSize
}

func GetCurrentTotalFileSize() int64 {
	i := GetStoredFilesFolder()
	return calculateTotalFileSize(i)
}

func GetFileList(MaxSpaceUsageBytes int64) []FileData {
	var FileList []FileData
	i := GetStoredFilesFolder()
	CurrentTotalFileSize := calculateTotalFileSize(i)
	for _, fi := range i {
		FileSurplusKeepTime := TotalKeepTimeCalc(fi.Size(), CurrentTotalFileSize, MaxSpaceUsageBytes) - time.Now().Sub(fi.ModTime())
		fdt := FileData{fi.Name(), fi.Size(), FileSurplusKeepTime}
		FileList = append(FileList, fdt)
	}
	return FileList
}

func TotalKeepTimeCalc(FileSizeBytes int64, CurrentTotalFileSizeBytes int64, MaxSpaceUsageBytes int64) time.Duration {
	totalTime := (MaxSpaceUsageBytes - CurrentTotalFileSizeBytes) / FileSizeBytes
	return time.Duration(totalTime) * time.Hour
}

func ReadMaxSpaceUsage() int64 {
	fileContent, _ := ioutil.ReadFile("maxSpaceUsage")
	maxSpaceUsageString := string(fileContent)
	maxSpaceUsage, _ := strconv.Atoi(maxSpaceUsageString)
	return int64(maxSpaceUsage)
}
