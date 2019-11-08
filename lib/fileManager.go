package lib

import (
	"mime/multipart"
	"os"
	"time"
)

func storeToLocal(fileName string, fileSize int64, file multipart.File) error {
	newFile, e := os.Create("files/" + fileName)
	fileContent := make([]byte, fileSize)
	_, e = file.Read(fileContent)
	_, e = newFile.Write(fileContent)
	e = newFile.Close()
	return e
}

func removeOutOfDateFiles() {

}

func keepTimeChecker() {
	currentHour := time.Now().Hour()
	for {
		if time.Now().Hour() != currentHour {
			currentHour = time.Now().Hour()

		}
	}
}
