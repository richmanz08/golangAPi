package video

import (
	// "bufio"
	// "bytes"

	"fmt"
	"net/http"

	// "net/http"

	// "log"

	"strings"

	// "time"

	// "io"
	// "log"
	// "net/http"
	// "os"

	"github.com/gin-gonic/gin"
)

type SubtitleURLStruct struct {
	MovieID  string `json:"mID"`
	Language string `json:"lang"`
}
type MediaURLStruct struct {
	MovieID string `json:"mID"`
}

var fileRoot = "assets/"

func ServerURLFileMediaM3U8(c *gin.Context) {
	var mediaOptions MediaURLStruct
	movieID := c.Request.URL.Query().Get("mID")
	mediaOptions.MovieID = movieID

	URLRoot := "movie/"
	fileName := "hotd_bandwidth" // waiting... db for know name file
	fileType := ".m3u8"

	result_file_name := fmt.Sprintf("http://localhost:8080/%s%s%s", URLRoot, fileName, fileType)
	c.JSON(http.StatusOK, result_file_name)
}
func ServerFileMedia(c *gin.Context) {

	file_name := c.Param("mID")
	fmt.Println("Filename was connected : ", file_name)
	// fileRoot := "assets/"
	typeFile := strings.Split(file_name, ".")
	typeFileName := typeFile[1]
	// fmt.Println("result file type TS :::",typeFileName == "ts")

	if typeFileName == "ts" {
		c.Writer.Header().Set("Content-Type", "application/octet-stream")
	} else {
		c.Writer.Header().Set("Content-Type", "application/x-mpegURL")
	}
	// rex := regexp.MustCompile("[0-9]+")
	// numberOfFileTS := rex.FindAllString(file_name, -1)
	// indexFileTS, err := strconv.Atoi(numberOfFileTS[0])
	// if err != nil {
	// 	fmt.Println("file index error")
	// }
	// if typeFileName == "ts" && indexFileTS > 5 {

	// 	newFilename := strings.Replace(file_name, "fhd", "low", 1)
	// 	c.File(fileRoot + newFilename)
	// 	fmt.Println("changed qulaity", newFilename)
	// 	return
	// }

	// fmt.Println("end log",file_name)
	c.File(fileRoot + file_name)
}

func ServerURLFileSubtitle(c *gin.Context) {
	var subtitleOptions SubtitleURLStruct

	movieID := c.Request.URL.Query().Get("mID")
	subtitle_lang := c.Request.URL.Query().Get("lang")

	subtitleOptions.MovieID = movieID
	subtitleOptions.Language = subtitle_lang

	// fmt.Println("movieID :::",movieID)
	// fmt.Println("subtitle_lang :::",subtitle_lang)
	// c.JSON(http.StatusOK,subtitleOptions)
	// fileRoot := "assets/"
	fileName := "example_subtitle" // waiting... db for know name file
	fileType := ".vtt"
	fileLang := strings.ToUpper(subtitleOptions.Language)
	result_file_name := fmt.Sprintf("http://localhost:8080/%s%s%s%s", fileRoot, fileName, fileLang, fileType)
	// fmt.Println("fileName :::",result_file_name)
	// fmt.Println("results path :::",fileRoot+result_file_name)

	// c.Writer.Header().Set("Content-Type","WEBVTT")
	c.JSON(http.StatusOK, result_file_name)
}

func ServerFileThumbnail(c *gin.Context) {
	c.Writer.Header().Set("Content-Type", "image/jpg")
	Filename := c.Param("file")
	// fileRoot := "assets/"
	c.File(fileRoot + Filename)
	c.Status(http.StatusOK)
}

//https://github.com/aofiee/Music-Streaming-HLS-Go-fiber/blob/main/main.go
