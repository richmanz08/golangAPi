package video

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

type SubtitleURLStruct struct {
	MovieID  string `json:"mID"`
	Language string `json:"lang"`
}
type MediaURLStruct struct {
	MovieID string `json:"mID"`
}


var filerootSubtitle = "D:/streamingfile/house_of_dragon/subtitle/"
var CONFIG_CONTENT_TYPE = "Content-Type"

func ServerURLFileMediaM3U8(c *gin.Context) {
	var mediaOptions MediaURLStruct
	movieID := c.Request.URL.Query().Get("mID")
	mediaOptions.MovieID = movieID

	URLRoot := "movie/"
	fileName := "hotd_bandwidth" // waiting... db for know name file
	fileType := ".m3u8"

	resultFileName := fmt.Sprintf("http://localhost:8080/%s%s%s", URLRoot, fileName, fileType)
	c.JSON(http.StatusOK, resultFileName)
}
func ServerFileMedia(c *gin.Context) {

	url := c.Param("name")

	if len(url) == 0 {
		c.JSON(http.StatusBadRequest,"ERROR LOAD FILE OR DESTINATION PATH")
		return
	}

	typeFile := strings.Split(url, ".")
	typeFileName := typeFile[1]

	var directoryName string

	if len(typeFile[0]) != 0 {
		dir := strings.Split(typeFile[0], "q")
		directoryName = dir[0]
	}

	
	

	


	if typeFileName == "ts" {
		c.Writer.Header().Set(CONFIG_CONTENT_TYPE, "application/octet-stream")
	} else {
		c.Writer.Header().Set(CONFIG_CONTENT_TYPE, "application/x-mpegURL")
	}

	bucket := os.Getenv("BUCKET_FILE_URL")
	result := bucket + directoryName + "/video/"+ url
log.Println(result)
	c.File(result)

}

func ServerURLFileSubtitle(c *gin.Context) {
	var subtitleOptions SubtitleURLStruct

	movieID := c.Request.URL.Query().Get("mID")
	subtitleLang := c.Request.URL.Query().Get("lang")

	subtitleOptions.MovieID = movieID
	subtitleOptions.Language = subtitleLang

	// fmt.Println("movieID :::",movieID)
	// fmt.Println("subtitle_lang :::",subtitle_lang)
	// c.JSON(http.StatusOK,subtitleOptions)
	// fileRoot := "assets/"
	fileName := "example_subtitle" // waiting... db for know name file
	fileType := ".vtt"
	fileLang := strings.ToUpper(subtitleOptions.Language)
	resultFileName := fmt.Sprintf("http://localhost:8080/%s%s%s%s", filerootSubtitle, fileName, fileLang, fileType)
	// fmt.Println("fileName :::",result_file_name)
	// fmt.Println("results path :::",fileRoot+result_file_name)

	// c.Writer.Header().Set("Content-Type","WEBVTT")
	c.JSON(http.StatusOK, resultFileName)
}

func ServerFileThumbnail(c *gin.Context) {
	c.Writer.Header().Set(CONFIG_CONTENT_TYPE, "image/jpeg")
	directory := c.Param("root")
	Filename := c.Param("file")

	bucket := os.Getenv("BUCKET_FILE_URL")
	result := bucket + directory + "/thumbnail/"+ Filename
	fmt.Println(result)
	c.File(result)
	c.Status(http.StatusOK)
}

//https://github.com/aofiee/Music-Streaming-HLS-Go-fiber/blob/main/main.go
