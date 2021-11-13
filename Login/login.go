package login

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Userdetail struct {
	ID        int          `json:"id"`
	Age       int          `json:"age"`
	Username  string       `json:"username"`
	Firstname string       `json:"firstname"`
	Lastname  string       `json:"lastname"`
	Role      string       `json:"role"`
	Token     *TokenStruct `json:"token"`
}

//A sample use
type User struct {
	ID       uint64 `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Phone    string `json:"phone"`
}

var user = User{
	ID:       1,
	Username: "username",
	Password: "password",
	Phone:    "49123454322", //this is a random number
}

func Login(c *gin.Context) {
	var userData []Userdetail
	var u User
	if err := c.ShouldBindJSON(&u); err != nil {
		c.JSON(http.StatusUnprocessableEntity, "Invalid json provided")
		return
	}
	//compare the user from the request, with the one we defined:
	if user.Username != u.Username || user.Password != u.Password {
		c.JSON(http.StatusUnauthorized, "Please provide valid login details")
		return
	}
	ts, err := CreateToken(user.ID)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}
	saveErr := CreateAuth(user.ID, ts)
	if saveErr != nil {
		c.JSON(http.StatusUnprocessableEntity, saveErr.Error())
	}
	// tokens := map[string]string{
	// 	"access_token":  ts.AccessToken,
	// 	"refresh_token": ts.RefreshToken,
	// }

	userData = append(userData, Userdetail{ID: 999, Age: 22, Username: "rizhmanz", Firstname: "arnon", Lastname: "reas", Role: "admin",
		Token: &TokenStruct{
			AccessToken: ts.AccessToken, RefrshToken: ts.RefreshToken,
		}})

	c.JSON(http.StatusOK, userData)
}

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
