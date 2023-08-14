package movies

import (
	"net/http"

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