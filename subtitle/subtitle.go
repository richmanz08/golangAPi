package subtitle

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

func GetAllFilenameSubtitleOfMovie(c *gin.Context) {

	// Read and parse JSON body
	var bodyParam ParamsSubtitle
	if err := c.BindJSON(&bodyParam); err != nil {
		log.Printf("Failed to parse JSON: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	// Open the directory
	folderLocation := fmt.Sprintf("D:/streamingfile/%s/subtitle/", bodyParam.Directory)
	dir, err := os.Open(folderLocation)
	if err != nil {
		log.Printf("Failed to open directory: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to open directory"})
		return
	}
	defer dir.Close()

	// Read the files in the directory
	fileInfos, err := dir.Readdir(-1)
	if err != nil {
		log.Printf("Failed to read directory: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to read directory"})
		return
	}

	// Prepare the array for JSON response
	var subtitleFiles []SubtitleFile

	// Iterate through fileInfos and create SubtitleFile objects
	for _, fileInfo := range fileInfos {
		if !fileInfo.IsDir() {
			// Determine language from file name
			language := strings.TrimSuffix(fileInfo.Name(), ".vtt")
			language = strings.ToUpper(language)

			// Check if language is known (ENGLISH or THAI)
			if language == "ENGLISH" || language == "THAI" {
				subtitleFile := SubtitleFile{
					Language: language,
					FileName: fileInfo.Name(),
				}
				subtitleFiles = append(subtitleFiles, subtitleFile)
			}
		}
	}

	// Convert subtitleFiles slice to JSON
	c.JSON(http.StatusOK, subtitleFiles)
}
