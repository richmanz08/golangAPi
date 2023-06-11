package movies

import (
	"database/sql"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

var DB *sql.DB

func CreateMovie(c *gin.Context) {
	file, header, err := c.Request.FormFile("file")

	if err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("file err : %s", err.Error()))
		return
	}
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

	if len(form.Value["episodes"]) != 0 {
		movie.Episodes = convertInt32(form.Value["episodes"][0])
	}

	if len(form.Value["description"]) != 0 {
		movie.Description = form.Value["description"][0]
	}
	if len(form.Value["directors_id"]) != 0 {
		movie.Directors_id = form.Value["directors_id"][0]
	}
	if len(form.Value["casters_id"]) != 0 {
		movie.Casters_id = form.Value["casters_id"][0]
	}

	querySQL := `INSERT INTO movies(
		movie_name_local,
		movie_name_eng,
		type,
		rating,
		duration,
		quality_type,
		directory_name,
		year,
		episodes,
		description,
		directors_id,
		casters_id) VALUES(?,?,?,?,?,?,?,?,?,?,?,?)
	`
	data, err := DB.Prepare(querySQL)

	if err != nil {
		c.JSON(http.StatusBadRequest, "### Insert table failed ###")
	} else {

		data.Exec(
			movie.MovieNameLocal,
			movie.MovieNameEng,
			movie.Type,
			movie.Rating,
			movie.Duration,
			movie.QualityType,
			movie.DirectoryName,
			movie.Year,
			movie.Episodes,
			movie.Description,
			movie.Directors_id,
			movie.Casters_id)
		c.JSON(http.StatusCreated, movie)
	}
	defer data.Close()

}

func convertInt32(str string) int32 {
	num, err := strconv.ParseInt(str, 10, 32)
	if err != nil {
		return 0
	}
	return int32(num)
}
func convertFloat64(str string) float64 {
	num, err := strconv.ParseFloat(str, 32)
	if err != nil {
		return 0
	}
	return float64(num)
}
