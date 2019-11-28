package lib

import (
	"log"
	"time"
)

func CheckAndDelete(maxKeepTimeDbJsonList *FileTotalKeepTime) {
	log.Println("Will delete outdated files.")
	for {
		fileList := GetFileList(maxKeepTimeDbJsonList)
		for _, fileItem := range fileList {
			//　每分钟检查一次，删除剩余保存时长很少的文件(1秒)
			if fileItem.FileSurplusKeepSeconds <= 1 {
				DeleteOneFile(fileItem.FileName, maxKeepTimeDbJsonList)
				log.Printf("Delete outdated file: %s", fileItem.FileName)
			}
		}
		time.Sleep(time.Minute)
	}
}

// gin mid-ware for checking upload files and
