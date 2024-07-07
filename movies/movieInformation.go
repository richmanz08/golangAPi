package movies

import (
	"net/http"
	"strconv"

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
		Description: form.Value["Description"][0],
		PosterPath: filepath,
	}

	result := DB.Create(&newMovie)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}
	c.JSON(http.StatusCreated, newMovie)
}

func GetAllInformationMovie(c *gin.Context) {

	var params ParamsMovieGroup
	var movies []MovieGroup

	params.Name = c.Param("Name")
	params.Status = c.Param("Status")
	pageSizeStr := c.DefaultQuery("pageSize", "10")
	currentStr := c.DefaultQuery("current", "1")

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
	c.ShouldBindJSON(&params)

	dynamicQuery := DB.Model(&MovieGroup{})

	if params.Name != "" {
		dynamicQuery = dynamicQuery.Where("name_eng LIKE ? OR name_local LIKE ?", "%"+params.Name+"%", "%"+params.Name+"%")
	}

	if params.Status != "" {
		dynamicQuery = dynamicQuery.Where("status IN (?)", params.Status)
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

	var simplifiedMovies []ResponseMovieGroup
	for _, movie := range movies {
		simplifiedMovie := ResponseMovieGroup{
			ID:         movie.ID,
			NameLocal:  movie.NameLocal,
			NameEng:    movie.NameEng,
			Type:       movie.Type,
			Status:     movie.Status,
			PosterPath: movie.PosterPath,
			Description: movie.Description,
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
			Current:  params.Current,
			Total:    int(totalCount),
		},
		StatusCode: 200,
	}

	c.JSON(http.StatusOK, response)

}

func GetOneInformationMovie( c * gin.Context){
	movieID := c.Param("id")
	
	var movie Movie
	query := DB.First(&movie, movieID)
	if query.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Movie not found"})
		return
	}

	response := Movie{
		ID:          movie.ID,
		MovieGroupID: movie.MovieGroupID,
		NameLocal:   movie.NameLocal,
		NameEng:     movie.NameEng,
		Type:        movie.Type,
		Status:      movie.Status,
		Duration: movie.Duration,
		QualityType: movie.QualityType,
		Year: movie.Year,
		Casters: movie.Casters,
		Directors: movie.Directors,
		Description: movie.Description,
		DirectoryName: movie.DirectoryName,
		CreatedAt:   movie.CreatedAt,
		UpdatedAt:   movie.UpdatedAt,
	}

	c.JSON(http.StatusOK, response)
}
