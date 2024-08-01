package continueplay

import (
	AUTH "api-webapp/Login"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetMyContinuePlay(c *gin.Context) {
	extractResult, err := AUTH.ExtractUserTokenMetadata(c.Request)

	if err != nil {
		c.JSON(http.StatusUnauthorized, "Token user error")
		return
	}

	var continuePlays []ResponseContinuePlay
	queryTable := DB.Model(&ContinuePlay{}).Where("continue_user_id = ?", uint(extractResult.UserID))
	if err := queryTable.Find(&continuePlays).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database query failed"})
		return
	}

	// Return the found records as JSON
	c.JSON(http.StatusOK, continuePlays)
}

func CreateMyContinuePlay(c *gin.Context) {
	extractResult, err := AUTH.ExtractUserTokenMetadata(c.Request)

	if err != nil {
		c.JSON(http.StatusUnauthorized, "Token user error")

	}

	var params ContinuePlayParams

	if err := c.ShouldBindJSON(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

	}


	newContinueRow := ContinuePlay{
		ContinueUserID: uint(extractResult.UserID),
		MovieGroupID:   params.MovieGroupID,
		MovieID:        params.MovieID,
		StampTime:      params.StampTime,
		EndTime:        params.EndTime,
	}
	if params.Season != 0 {
		newContinueRow.Season = params.Season
	}

	var UpdateContinuePlay ContinuePlay
	dynamicQuery := DB.Model(&ContinuePlay{})
	dynamicQuery.Where("movie_id = ?",params.MovieID)
	dynamicQuery.Where("continue_user_id = ?",extractResult.UserID)

	isDuplicatedContinueRow := dynamicQuery.First(&UpdateContinuePlay)

	if isDuplicatedContinueRow.Error != nil {

		result := DB.Create(&newContinueRow)
		if result.Error != nil {
			c.JSON(http.StatusBadRequest, "failed to insert table")

		}
		c.JSON(http.StatusCreated, "create a new record success")
	} else {

		UpdateContinuePlay.StampTime = params.StampTime
		result := DB.Save(&UpdateContinuePlay)
		if result.Error != nil {
			c.JSON(http.StatusBadRequest, "failed to insert table")

		}
		c.JSON(http.StatusCreated, "Updated record success")
	}

	

}
