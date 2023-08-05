package movies

import (
	"database/sql"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strconv"

	"github.com/gin-gonic/gin"
)

var DB *sql.DB

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

	var movie IMovie
	movie.MovieNameLocal = form.Value["movie_name_local"][0]
	movie.MovieNameEng = form.Value["movie_name_eng"][0]
	movie.Type = form.Value["type"][0]
	movie.Rating = convertFloat64(form.Value["rating"][0])
	movie.Duration = convertFloat64(form.Value["duration"][0])
	movie.QualityType = form.Value["quality_type"][0]
	movie.PosterURL = filepath
	movie.DirectoryName = form.Value["directory_name"][0]
	movie.Year = convertInt32(form.Value["year"][0])
	movie.Episodes = convertInt32(form.Value["episodes"][0])
	movie.Description = form.Value["description"][0]
	movie.Directors_id = form.Value["directors_id"][0]
	movie.Casters_id = form.Value["casters_id"][0]


	querySQL := `INSERT INTO movies(
		movie_name_local,
		movie_name_eng,
		type,
		rating,
		duration,
		quality_type,
		poster_url,
		directory_name,
		year,
		episodes,
		description,
		directors_id,
		casters_id) VALUES(?,?,?,?,?,?,?,?,?,?,?,?,?)
	`
	data, err := DB.Prepare(querySQL)

	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, "failed to insert table")
		return
	}
	defer data.Close()

	result, err := data.Exec(
		movie.MovieNameLocal,
		movie.MovieNameEng,
		movie.Type,
		movie.Rating,
		movie.Duration,
		movie.QualityType,
		movie.PosterURL,
		movie.DirectoryName,
		movie.Year,
		movie.Episodes,
		movie.Description,
		movie.Directors_id,
		movie.Casters_id)

	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, "failed to create movie")
		return
	}

	// Retrieve the last inserted record ID
	lastInsertID, err := result.LastInsertId()
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, "failed to retrieve last insert ID")
		return
	}
	response, err := DB.Query("SELECT * FROM movies WHERE id=?", lastInsertID)
	if err != nil {
		fmt.Println(err)
	}
	defer response.Close()
	for response.Next() {
		var new IMovie
		err = response.Scan(
			&new.ID,
			&new.MovieNameLocal,
			&new.MovieNameEng,
			&new.Episodes,
			&new.CreateAt,
			&new.UpdateAt,
			&new.DeleteAt,
			&new.Type,
			&new.Rating,
			&new.Duration,
			&new.Description,
			&new.QualityType,
			&new.Directors_id,
			&new.Casters_id,
			&new.PosterURL,
			&new.DirectoryName,
			&new.Year,
		)
		if err != nil {
			panic(err.Error())
		}
		movie.ID = new.ID
		movie.CreateAt = new.CreateAt
		movie.UpdateAt = new.UpdateAt

	}

	createFolderResponse := createFolder(movie)
	if !createFolderResponse {
		c.JSON(http.StatusBadRequest, "duplicate ?")
		return
	}

	c.JSON(http.StatusCreated, movie)

}

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
func createFolder(movie IMovie) bool{
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
