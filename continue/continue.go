package continueplay

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetMyContinuePlay(c *gin.Context) {
// 	au, err := AUTH.ExtractTokenMetadata(c.Request)
// 	fmt.Println(au.AccessUuid,au.UserId)
// 	if err != nil {
// 		c.JSON(http.StatusUnauthorized, "unauthorized")
// 		return
// 	}

// 	userId, err := AUTH.FetchAuth(au)
// 	if err != nil {
// 		c.JSON(http.StatusUnauthorized, "Unauthorized")
// 	}
// 	_userId := userId
// fmt.Println(_userId)
	var continuePlays []ContinuePlay
	queryTable := DB.Model(&ContinuePlay{}).Where("continue_user_id = ?", uint(1))
	if err := queryTable.Find(&continuePlays).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database query failed"})
		return
	}

	// Return the found records as JSON
	c.JSON(http.StatusOK, continuePlays)
}
