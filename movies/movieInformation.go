package movies

import (
	"errors"
	"fmt"
	"net/http"

	common "api-webapp/common"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
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
		NameLocal:   form.Value["NameLocal"][0],
		NameEng:     form.Value["NameEng"][0],
		Type:        form.Value["Type"][0],
		Status:      form.Value["Status"][0],
		Description: form.Value["Description"][0],
		PosterPath:  filepath,
	}

	result := DB.Create(&newMovie)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}
	c.JSON(http.StatusCreated, newMovie)
}

func GetAllMovieGroup(c *gin.Context) {
	params, err := readQueryStringMovieGroup(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid query string"})
		return
	}

	dynamicQuery, err := createDynamicQueryMovieGroup(params)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var moviesGroup []MovieGroup

	if err := handlePaginationAndQueryMovieGroup(dynamicQuery, &params, &moviesGroup); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var simplifiedMovies []ResponseMovieGroup
	for _, movie := range moviesGroup {
		var MovieRowItem Movie
		queryMovieByGroupID := DB.Model(&Movie{}).Where("movie_group_id = ?", movie.ID).Order("season DESC").First(&MovieRowItem)

		if queryMovieByGroupID.Error != nil {
			if errors.Is(queryMovieByGroupID.Error, gorm.ErrRecordNotFound) {
				fmt.Println("Not found this movie group id in table: movie", movie.ID)
			} else {
				fmt.Println("Error occurred:", queryMovieByGroupID.Error)
			}
		} else {
			fmt.Println("Found ::", MovieRowItem.NameEng, "Season:", MovieRowItem.Season)
		}

		simplifiedMovie := ResponseMovieGroup{
			ID:           movie.ID,
			NameLocal:    movie.NameLocal,
			NameEng:      movie.NameEng,
			Type:         movie.Type,
			Status:       movie.Status,
			PosterPath:   movie.PosterPath,
			Description:  movie.Description,
			CreatedAt:    movie.CreatedAt,
			UpdatedAt:    movie.UpdatedAt,
			MovieSeason:  int(MovieRowItem.Season),
			MovieTime:    MovieRowItem.Duration,
			MovieQuality: MovieRowItem.QualityType,
			MovieYear:    int(MovieRowItem.Year),
		}

		simplifiedMovies = append(simplifiedMovies, simplifiedMovie)
	}

	var totalCount int64
	countQuery, err := createDynamicQueryMovieGroup(params)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := countQuery.Count(&totalCount).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to count total records"})
		return
	}

	response := common.Response{
		Data: simplifiedMovies,
		Pagination: common.Pagination{
			PageSize: params.PageSize,
			Current:  *params.Current,
			Total:    int(totalCount),
		},
		StatusCode: http.StatusOK,
	}

	c.JSON(http.StatusOK, response)

}

func GetOneInformationMovie(c *gin.Context) {
	movieID := c.Param("id")

	var movie Movie
	query := DB.First(&movie, movieID)
	if query.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Movie not found"})
		return
	}

	response := Movie{
		ID:            movie.ID,
		MovieGroupID:  movie.MovieGroupID,
		NameLocal:     movie.NameLocal,
		NameEng:       movie.NameEng,
		Type:          movie.Type,
		Status:        movie.Status,
		Duration:      movie.Duration,
		QualityType:   movie.QualityType,
		Year:          movie.Year,
		Casters:       movie.Casters,
		Directors:     movie.Directors,
		Description:   movie.Description,
		DirectoryName: movie.DirectoryName,
		Season: movie.Season,
		CreatedAt:     movie.CreatedAt,
		UpdatedAt:     movie.UpdatedAt,
	}

	c.JSON(http.StatusOK, response)
}
