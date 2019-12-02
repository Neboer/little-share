package lib

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"time"
	"strings"
)

type FileData struct {
	FileName               string
	FileSizeBytes          int64
	FileSurplusKeepSeconds int64
}

type FileTotalKeepTime map[string]time.Duration

func StoreToLocal(c *gin.Context, MaxSpaceUsage int64, maxKeepTimeDbJsonList *FileTotalKeepTime) error {
	defer func() {
		err := recover()
		if err != nil {
			if err == io.ErrUnexpectedEOF {
				c.AbortWithStatusJSON(500, gin.H{"server": "bad upload file(s)"})
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
	if header.Size >= MaxSpaceUsage-GetCurrentTotalFileSize() {
		c.AbortWithStatusJSON(413, gin.H{"error": "file too large!"})
	}
	(*maxKeepTimeDbJsonList)[filename] = TotalKeepTimeCalc(header.Size, GetCurrentTotalFileSize(), MaxSpaceUsage)
	WriteKeepTimeDB(maxKeepTimeDbJsonList)
	err = c.SaveUploadedFile(header, "files/"+filename)
	return err
}

func GetStoredFilesFolder() []os.FileInfo {
	f, _ := os.Open("files")
	i, _ := f.Readdir(-1)
	return i
}

const (
	OperationSuccessful = 0
	NoSuchFileError     = 1
	OtherError          = 2
)

type FileError int8

func DeleteOneFile(fileName string, maxKeepTimeDbJsonList *FileTotalKeepTime) FileError {
	if _, isExist := (*maxKeepTimeDbJsonList)[fileName]; isExist == true {
		delete(*maxKeepTimeDbJsonList, fileName)
		WriteKeepTimeDB(maxKeepTimeDbJsonList)
		err := os.Remove("./files/" + fileName)
		if err == nil {
			return OperationSuccessful
		} else if err == err.(*os.PathError) {
			return NoSuchFileError // 应该不会删不去文件的吧...
		} else {
			// 500了……
			log.Println(err.Error())
			return OtherError
		}
	} else {
		//　文件肯定不存在
		return NoSuchFileError
	}
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

func GetFileList(maxKeepTimeDbJsonList *FileTotalKeepTime) []FileData {
	var FileList []FileData
	i := GetStoredFilesFolder()
	for _, fi := range i {
		FileSurplusKeepTime := (*maxKeepTimeDbJsonList)[fi.Name()] - time.Now().Sub(fi.ModTime())
		fdt := FileData{fi.Name(), fi.Size(), int64(FileSurplusKeepTime / time.Second)}
		FileList = append(FileList, fdt)
	}
	return FileList
}

func TotalKeepTimeCalc(FileSizeBytes int64, CurrentTotalFileSizeBytes int64, MaxSpaceUsageBytes int64) time.Duration {
	var totalDays float64
	totalDays = (float64(MaxSpaceUsageBytes) - float64(CurrentTotalFileSizeBytes)) / float64(FileSizeBytes)
	return time.Duration(totalDays * 24 * float64(time.Hour))
}

func ReadMaxSpaceUsage() int64 {
	fileContent, _ := ioutil.ReadFile("maxSpaceUsage")
	maxSpaceUsageString := string(fileContent)
	maxSpaceUsage, _ := strconv.Atoi(strings.TrimSpace(maxSpaceUsageString))
	return int64(maxSpaceUsage)
}

func ReadKeepTimeDB() FileTotalKeepTime {
	maxKeepTimeDbJsonData, err := ioutil.ReadFile("filesMaxKeepTime.json")
	if err != nil {
		log.Fatalf("cannot load json data filesMaxKeepTime.json. %s", err.Error())
	}
	var maxKeepTimeDb FileTotalKeepTime
	err = json.Unmarshal(maxKeepTimeDbJsonData, &maxKeepTimeDb)
	if err != nil {
		log.Fatalf("cannot parse json data. %s", err.Error())
	}
	return maxKeepTimeDb
}

func WriteKeepTimeDB(FileTotalKeepTimeData *FileTotalKeepTime) {
	FileTotalKeepTimeBytes, _ := json.Marshal(FileTotalKeepTimeData)
	err := ioutil.WriteFile("filesMaxKeepTime.json", FileTotalKeepTimeBytes, 0755)
	if err != nil {
		log.Fatalf("cannot write file! %s", err.Error())
	}
}
