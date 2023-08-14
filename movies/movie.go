package movies

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)



func AddMovie(c *gin.Context) {

	var MovieParams Movie
	var GroupMovie MovieGroup
	if err := c.ShouldBindJSON(&MovieParams); err != nil {
		fmt.Println(err)
		return
	}

	// เช็คว่า MovieGroupID ตรงกับที่มีหรือไม่
	matchMovieGroupByID := DB.First(&GroupMovie,MovieParams.MovieGroupID)
	fmt.Println("matchMovieGroupByID response:", matchMovieGroupByID.Error)
	if matchMovieGroupByID.Error == gorm.ErrRecordNotFound {
		c.JSON(http.StatusBadRequest, "Not found Movie Group ? Please check")
		return
	}


	newMovieDetail := Movie{
		MovieGroupID: MovieParams.MovieGroupID,
		NameLocal:     MovieParams.NameLocal,
		NameEng:       MovieParams.NameEng,
		Type:          MovieParams.Type,
		Duration:      MovieParams.Duration,
		Status: MovieParams.Status,
		Season: MovieParams.Season,
		Episode: MovieParams.Episode,
		QualityType:   MovieParams.QualityType,
		DirectoryName: MovieParams.DirectoryName,
		Year:          MovieParams.Year,
		Description:   MovieParams.Description,
		DirectorsID:   MovieParams.DirectorsID,
		CastersID:     MovieParams.CastersID,
	}

	// 1. เช็คก่อนว่ามีซ้ำกันหรือไม่
	createFolderResponse := createFolder(newMovieDetail)
	if !createFolderResponse {
		c.JSON(http.StatusBadRequest, "Duplicate folder directory name")
		return
	}

	// 2. ทำการ insert ลง Database
	result := DB.Create(&newMovieDetail)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, "failed to insert table")
		return
	}

	c.JSON(http.StatusCreated, newMovieDetail)
}

