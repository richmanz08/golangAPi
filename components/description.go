package components

import (
	Ex "api-webapp/Login"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type DescriptionStruct struct {
	UserID       uint64 `json:"userid"`
	Title        string `json:"title"`
	Descriptions string `json:"description"`
}

func Components(c *gin.Context) {
	var Des *DescriptionStruct
	if err := c.ShouldBindJSON(&Des); err != nil {
		c.JSON(http.StatusUnprocessableEntity, "Invalid Json [COMPONENTS API : description.go]")
		return
	}

	tokenAuth, err := Ex.ExtractTokenMetadata(c.Request)
	fmt.Println("checktoken", tokenAuth)
	if err != nil {
		c.JSON(http.StatusUnauthorized, "Unauthorized")
	}
	userId, err := Ex.FetchAuth(tokenAuth)
	if err != nil {
		c.JSON(http.StatusUnauthorized, "Unauthorized")
	}
	Des.UserID = userId

	c.JSON(http.StatusCreated, Des)

}
