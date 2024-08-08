package video

import (
	"fmt"
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

var CONFIG_CONTENT_TYPE = "Content-Type"


func ServerFileMedia(c *gin.Context) {

	url := c.Param("name")

	if len(url) == 0 {
		c.JSON(http.StatusBadRequest, "ERROR LOAD FILE OR DESTINATION PATH")
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
	result := bucket + directoryName + "/video/" + url
	
	c.File(result)

}

func ServerURLFileSubtitle(c *gin.Context) {
	c.Writer.Header().Set("Content-Type", "text/vtt") // Set Content-Type to WebVTT
    
    directory := c.Param("directory")
    filename := c.Param("filename")
    
    bucket := os.Getenv("BUCKET_FILE_URL")
    
    // Construct the file path
    resultFileName := fmt.Sprintf("%s%s/subtitle/%s", bucket, directory, filename)
    fmt.Println("File path:", resultFileName)
    
    // Check if the file exists
    if _, err := os.Stat(resultFileName); os.IsNotExist(err) {
        c.JSON(http.StatusNotFound, gin.H{"error": "Subtitle file not found"})
        return
    }
    
    // Serve the file
    c.File(resultFileName)
}

func ServerFileThumbnail(c *gin.Context) {

	c.Writer.Header().Set("Content-Type", "image/jpeg")

	directory := c.Param("root")
	filename := c.Param("file")

	bucket := os.Getenv("BUCKET_FILE_URL")

	// Construct the file path
	result := fmt.Sprintf("%s%s/thumbnail/%s.jpeg", bucket, directory, filename)
	fmt.Println("Serving file:", result)

	// Check if the file exists
	if _, err := os.Stat(result); os.IsNotExist(err) {
		c.JSON(http.StatusNotFound, gin.H{"error": "File not found"})
		return
	}

	// Check if the file exists
	if _, err := os.Stat(result); os.IsNotExist(err) {
		c.JSON(http.StatusNotFound, gin.H{"error": "File not found"})
		return
	}

	// Serve the file
	c.File(result)
}

func ServerFilePoster(c *gin.Context){
	c.Writer.Header().Set("Content-Type", "image/jpeg")
	bucket := os.Getenv("BUCKET_FILE_URL")
	filename := c.Param("file")
	
	result := fmt.Sprintf("%spublic/%s", bucket, filename)

	// Check if the file exists
	if _, err := os.Stat(result); os.IsNotExist(err) {
		c.JSON(http.StatusNotFound, gin.H{"error": "File not found"})
		return
	}

	// Check if the file exists
	if _, err := os.Stat(result); os.IsNotExist(err) {
		c.JSON(http.StatusNotFound, gin.H{"error": "File not found"})
		return
	}
	

	// Serve the file
	c.File(result)
	
}

//https://github.com/aofiee/Music-Streaming-HLS-Go-fiber/blob/main/main.go
