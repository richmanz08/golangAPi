package movies

import (
	"api-webapp/common"
	"fmt"
	"log"
	"net/http"

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

func GetRecommendMovieNow(c *gin.Context){
	var movieGroupData MovieGroup
	if err := DB.Order("updated_at asc").First(&movieGroupData).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to fetch recomend movie"})
		return
	}

	var mv Movie

	if err := DB.Where("movie_group_id = ?", movieGroupData.ID).First(&mv) .Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to fetch recomend movie"})
		return
	}
	posterURL := fmt.Sprintf("image/%s",  movieGroupData.PosterPath)
	response := ResponseMovieRecommendNow{
		MovieGroupID: int(mv.MovieGroupID),
		NameLocal: mv.NameLocal,
		NameEng: mv.NameEng,
		Description:mv.Description,
		Type: mv.Type,
		MovieQuality: mv.QualityType,
		MovieTime:mv.Duration,
		PosterURL:posterURL,
		UpdatedAt: mv.UpdatedAt,
	}

	c.JSON(http.StatusOK, response)
}


