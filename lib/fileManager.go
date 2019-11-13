package lib

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"strconv"
	"time"
)

func StoreToLocal(c *gin.Context) error {
	_, header, err := c.Request.FormFile("file")
	filename := header.Filename
	fmt.Println(header.Filename)
	err = c.SaveUploadedFile(header, "files/"+filename)
	//out, err := os.Create("./files/" + filename)
	//log.Println(err)
	//if err != nil {
	//	log.Println(err)
	//}
	//defer out.Close()
	//_, err = io.Copy(out, file)
	//if err != nil {
	//	log.Fatal(err)
	//}
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
