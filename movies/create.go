package movies

import (
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"os/exec"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// var DB *sql.DB
var DB *gorm.DB

func CreateMovie(c *gin.Context) {
	file, header, err := c.Request.FormFile("file")

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to read the file"})
		return
	}
	defer file.Close()

	form, err := c.MultipartForm()

	if err != nil {
		c.JSON(http.StatusBadRequest, "create movie failed")
	}


	filepath, errUpload := uploadImage(file,header)
	if errUpload != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to upload poster"})
		return
	}
	
	newMovieInformation := Movie_information{
		MovieNameLocal:  form.Value["movie_name_local"][0],
		MovieNameEng:form.Value["movie_name_eng"][0],
		Type: form.Value["type"][0],
		Rating:convertFloat64(form.Value["rating"][0]),
		Duration: convertFloat64(form.Value["duration"][0]),
		QualityType: form.Value["quality_type"][0],
		PosterURL:filepath,
		DirectoryName:form.Value["directory_name"][0],
		Year:  convertInt32(form.Value["year"][0]),
		Episodes: convertInt32(form.Value["episodes"][0]),
		Description: form.Value["description"][0],
		DirectorsID: form.Value["directors_id"][0],
		CastersID: form.Value["casters_id"][0],
	}

	// 1. เช็คก่อนว่ามีซ้ำกันหรือไม่
	createFolderResponse := createFolder(newMovieInformation)
	if !createFolderResponse {
		c.JSON(http.StatusBadRequest, "duplicate folder directory ?")
		return
	}

	// 2. ทำการ insert ลง Database
	result := DB.Create(&newMovieInformation)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, "failed to insert table")
		return
	}

	


	c.JSON(http.StatusCreated, newMovieInformation)

}

// FUNCTION :) SUPPORT CREAT MOVIE API

func convertInt32(str string) int32 {
	num, err := strconv.ParseInt(str, 10, 32)
	if err != nil {
		return 0
	}
	return int32(num)
}
func convertFloat64(str string) float64 {
	floatValue, err := strconv.ParseFloat(str, 64)
	if err != nil {
		return 0.0
	}
	return floatValue
}
func createFolder(movie Movie_information) bool{
	var bucketFolder = "D:\\streamingfile\\"

	// Replace "YourDesiredFolderName" with the name of the folder you want to create
	folderName := movie.DirectoryName
	// Replace "D:\\" with the desired path on the D drive where you want to create the folder
	folderPath := bucketFolder + folderName

	// Create the command to execute
	cmd := exec.Command("cmd", "/c", "mkdir", folderPath)

	errCommand := cmd.Run()
	if errCommand != nil {
		fmt.Println("Error creating folder:", errCommand)
		 return false
		// os.Exit(1)
	}

	// thumbnail sub folder
	thumbnailPathFolder := bucketFolder + folderName + "\\thumbnail"
	// Create the command to execute
	cmdthumbnail := exec.Command("cmd", "/c", "mkdir", thumbnailPathFolder)
	cmdthumbnail.Run()

	// subtitle sub folder
	subtitlePathFolder := bucketFolder + folderName + "\\subtitle"
	// Create the command to execute
	cmdsubtitle := exec.Command("cmd", "/c", "mkdir", subtitlePathFolder)
	cmdsubtitle.Run()

	// audio sub folder
	audioPathFolder := bucketFolder + folderName + "\\audio"
	// Create the command to execute
	cmdaudio := exec.Command("cmd", "/c", "mkdir", audioPathFolder)
	cmdaudio.Run()

	// video sub folder
	videoPathFolder := bucketFolder + folderName + "\\video"
	// Create the command to execute
	cmdvideo := exec.Command("cmd", "/c", "mkdir", videoPathFolder)
	cmdvideo.Run()
	return true
}
func uploadImage(file multipart.File,header *multipart.FileHeader) (string, error){
	filename := header.Filename
	out, err := os.Create("public/" + filename)
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()
	
	_, err = io.Copy(out, file)
	if err != nil {
		log.Fatal(err)
	}
	filepath := "http://localhost:8080/public/" + filename
	return filepath, nil
}