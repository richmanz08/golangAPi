package continueplay

import (
	AUTH "api-webapp/Login"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetMyContinuePlay(c *gin.Context) {
	extractResult,err := AUTH.ExtractUserTokenMetadata(c.Request)

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
