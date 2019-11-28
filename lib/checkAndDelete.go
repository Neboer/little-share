package lib

import "time"

func CheckAndDelete(maxKeepTimeDbJsonList *FileTotalKeepTime) {
	for {
		fileList := GetFileList(maxKeepTimeDbJsonList)
		for _, fileItem := range fileList {
			//　总保存时长小于0
			if fileItem.FileSurplusKeepTime <= time.Duration(0) {
				DeleteOneFile(fileItem.FileName, maxKeepTimeDbJsonList)
			}
		}
	}
}

// gin mid-ware for checking upload files and
