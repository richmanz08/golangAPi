package movies

import (
	"api-webapp/common"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)



func AddMovie(c *gin.Context) {

	var MovieParams Movie
	var GroupMovie MovieGroup
	if err := c.ShouldBindJSON(&MovieParams); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
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

func GetAllMovie(c *gin.Context){
	var params ParamsMovies
	var movies []Movie


	movieGroupIdStr := c.DefaultQuery("MovieGroupID","0")
	pageSizeStr := c.DefaultQuery("pageSize", "10")
	currentStr := c.DefaultQuery("current", "1")

	MovieGroupID, err := strconv.Atoi(movieGroupIdStr)
	if err != nil {
		c.String(http.StatusBadRequest, "Invalid pagesize value")
		return
	}

	pageSize, err := strconv.Atoi(pageSizeStr)
	if err != nil {
		c.String(http.StatusBadRequest, "Invalid pagesize value")
		return
	}

	current, err := strconv.Atoi(currentStr)
	if err != nil {
		c.String(http.StatusBadRequest, "Invalid current value")
		return
	}

	params.PageSize = pageSize
	params.Current = current
	params.MovieGroupID = MovieGroupID
	if err := c.ShouldBindJSON(&params); err != nil {
		log.Println(err.Error())
	}
	dynamicQuery := DB.Model(&Movie{})

	if params.MovieGroupID != 0 {
		dynamicQuery = dynamicQuery.Where("movie_group_id = ?", params.MovieGroupID)
	}

	paginatedDB, err := common.Paginate(dynamicQuery, params.Current, params.PageSize)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	queryFindall := paginatedDB.Find(&movies)
	if queryFindall.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": queryFindall.Error})
		return
	}

	if len(movies) == 0 && params.Current > 1 {

		params.Current = 1
		dynamicQuery = DB.Model(&MovieGroup{})
		againPaginatedDB, err := common.Paginate(dynamicQuery, params.Current, params.PageSize)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		againQueryFindall := againPaginatedDB.Find(&movies)
		if againQueryFindall.Error != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": againQueryFindall.Error})
			return
		}

	}

	var totalCount int64
	countTotal := DB.Model(&Movie{}).Count(&totalCount)
	if countTotal.Error != nil {
		panic(countTotal.Error)
	}

	response := common.Response{
		Data: movies,
		Pagination: common.Pagination{
			PageSize: params.PageSize,
			Current:  params.Current,
			Total:    int(totalCount),
		},
		StatusCode: 200,
	}
	c.JSON(http.StatusOK,response)
}