package movies

import (
	"net/http"

	common "api-webapp/common"

	"github.com/gin-gonic/gin"
)

func CreateInformationMovie(c *gin.Context) {
	file, header, err := c.Request.FormFile("file")

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	defer file.Close()

	form, err := c.MultipartForm()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	filepath, err := uploadImage(file, header)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	newMovie := MovieGroup{
		NameLocal:  form.Value["NameLocal"][0],
		NameEng:    form.Value["NameEng"][0],
		Type:       form.Value["Type"][0],
		Status:     form.Value["Status"][0],
		PosterPath: filepath,
	}

	
	result := DB.Create(&newMovie)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}
	c.JSON(http.StatusCreated, newMovie)
}


func GetAllInformationMovie(c *gin.Context){

	var params ParamsMovieGroup

	if err := c.ShouldBindJSON(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}


	var movies  []MovieGroup

	dynamicQuery := DB.Model(&MovieGroup{})

	if params.Name != "" {
 		dynamicQuery = dynamicQuery.Where("name_eng LIKE ? OR name_local LIKE ?","%"+params.Name+"%","%"+params.Name+"%")
	}


   if params.Status != ""  {
		dynamicQuery = dynamicQuery.Where("status IN (?)",params.Status)
	}

	paginatedDB, err := common.Paginate(dynamicQuery, params.Current, params.PageSize)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error":err.Error()})
		return 
    }


	queryFindall := paginatedDB.Find(&movies)


	if queryFindall.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": queryFindall.Error})
		return
	}
	var simplifiedMovies []ResponseMovieGroup
	for _, movie := range movies {
		simplifiedMovie := ResponseMovieGroup{
			ID:         movie.ID,
			NameLocal:  movie.NameLocal,
			NameEng:    movie.NameEng,
			Type:       movie.Type,
			Status:     movie.Status,
			PosterPath: movie.PosterPath,
			CreatedAt:  movie.CreatedAt,
			UpdatedAt:  movie.UpdatedAt,
		}
		simplifiedMovies = append(simplifiedMovies, simplifiedMovie)
	}

	var totalCount int64
	countTotal := DB.Model(&MovieGroup{}).Count(&totalCount)
	if countTotal.Error != nil {
		panic(countTotal.Error)
	}

	response := common.Response{
		Data: simplifiedMovies,
		Pagination: common.Pagination{
			PageSize: params.PageSize,
			Current:  1,
			Total: int(totalCount) ,
		},
		StatusCode: 200,
	}

	c.JSON(http.StatusOK, response)


}