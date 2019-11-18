package lib

import "time"

func CheckAndDelete(maxKeepTimeDbJsonList *FileTotalKeepTime) {
	for {
		fileList := GetFileList(maxKeepTimeDbJsonList)
		for _, fileItem := range fileList {
			//　总保存时长小于一个小时
			if fileItem.FileSurplusKeepTime <= time.Hour {
				DeleteOneFile(fileItem.FileName, maxKeepTimeDbJsonList)
			}
		}
	}
}

// gin mid-ware for checking upload files and
