package lib

import (
	"log"
	"os"
	"time"
)

func CheckAllRequireFilesAndFolders(){
	info, err := os.Stat("files")
	if err != nil {
		log.Fatalf("cannot find files folder. %s", err.Error())
	} else if !info.IsDir() {
		log.Fatalf("files is a folder not file!")
	}
	info, err = os.Stat("filesMaxKeepTime.json")
	if err != nil {
		log.Fatalf("cannot find filesMaxKeepTime.json. %s", err.Error())
	} else if info.IsDir() {
		log.Fatalf("filesMaxKeepTime.json is not a folder but file!")
	}
	info, err = os.Stat("maxSpaceUsage")
	if err != nil {
		log.Fatalf("cannot find filesMaxKeepTime.json. %s", err.Error())
	} else if info.IsDir() {
		log.Fatalf("filesMaxKeepTime.json is not a folder but file!")
	}
}

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
