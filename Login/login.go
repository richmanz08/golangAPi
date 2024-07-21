package login

import (
	"net/http"

	"github.com/gin-gonic/gin"
)



func Logout(c *gin.Context) {
	au, err := ExtractTokenMetadata(c.Request)
	if err != nil {
		c.JSON(http.StatusUnauthorized, "unauthorized")
		return
	}
	deleted, delErr := DeleteAuth(au.AccessUuid)
	if delErr != nil || deleted == 0 { //if any goes wrong
		c.JSON(http.StatusUnauthorized, "unauthorized")
		return
	}
	c.JSON(http.StatusOK, "Successfully logged out")
}

// การแปลงข้อมูลที่ได้มาเป็น json
// resRp := &rp
// resRp2, _ := json.Marshal(resRp)
//  fmt.Println("Rp", string(resRp2))
