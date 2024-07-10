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
		Casters: MovieParams.Casters,
		Directors: MovieParams.Directors,
	}

	log.Println(newMovieDetail)

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
	


	params, err := parseQueryParams(c)
	if err != nil {
		c.String(http.StatusBadRequest, "Invalid query parameter value")
		return
	}




	if err := c.ShouldBindJSON(&params); err != nil {
		log.Println(err.Error())
	}



	dynamicQuery := createDynamicQuery(params)
	totalCount, err := countTotalRecords(dynamicQuery)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var movies []Movie
	if err := handlePaginationAndQuery(dynamicQuery, params, &movies); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
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


// this help funtion for refactor code cleanup
func parseQueryParams(c *gin.Context) (ParamsMovies, error) {
	var params ParamsMovies

	movieGroupIdStr := c.DefaultQuery("MovieGroupID", "0")
	pageSizeStr := c.DefaultQuery("pageSize", "10")
	currentStr := c.DefaultQuery("current", "1")
	seasonStr := c.DefaultQuery("season", "0")

	movieGroupID, err := strconv.Atoi(movieGroupIdStr)
	if err != nil {
		return params, err
	}

	pageSize, err := strconv.Atoi(pageSizeStr)
	if err != nil {
		return params, err
	}

	current, err := strconv.Atoi(currentStr)
	if err != nil {
		return params, err
	}

	season, err := strconv.Atoi(seasonStr)
	if err != nil {
		return params, err
	}

	params.MovieGroupID = movieGroupID
	params.PageSize = pageSize
	params.Current = current
	params.Season = season

	return params, nil
}
func createDynamicQuery(params ParamsMovies) *gorm.DB {
	dynamicQuery := DB.Model(&Movie{})

	if params.MovieGroupID != 0 {
		dynamicQuery = dynamicQuery.Where("movie_group_id = ?", params.MovieGroupID)
	}

	if params.Season != 0 {
		dynamicQuery = dynamicQuery.Where("season = ?", params.Season)
	}

	return dynamicQuery
}
func countTotalRecords(dynamicQuery *gorm.DB) (int64, error) {
	var totalCount int64
	countTotal := dynamicQuery.Count(&totalCount)
	if countTotal.Error != nil {
		return 0, countTotal.Error
	}
	return totalCount, nil
}
func handlePaginationAndQuery(dynamicQuery *gorm.DB, params ParamsMovies, movies *[]Movie) error {
	paginatedDB, err := common.Paginate(dynamicQuery, params.Current, params.PageSize)
	if err != nil {
		return err
	}

	queryFindAll := paginatedDB.Find(movies)
	if queryFindAll.Error != nil {
		return queryFindAll.Error
	}

	if len(*movies) == 0 && params.Current > 1 {
		params.Current = 1
		dynamicQuery = createDynamicQuery(params)
		paginatedDB, err = common.Paginate(dynamicQuery, params.Current, params.PageSize)
		if err != nil {
			return err
		}

		queryFindAll = paginatedDB.Find(movies)
		if queryFindAll.Error != nil {
			return queryFindAll.Error
		}
	}

	return nil
}