package login

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type handleParamsRouteLoginStruct struct {
	Email    string `json:"email"  binding:"required"`
	Password string `json:"password"  binding:"required"`
}
type handleParamsRoutePINStruct struct {
	AccountId int32  `json:"account_id" binding:"required"`
	UserIndex string `json:"user_index"  binding:"required"`
	PIN string `json:"pin"  binding:"required"`
}
type responseStructLogin struct {
	AccountId int32        `json:"account_id" `
	Email     string       `json:"email" `
	Role      string       `json:"role" `
	Token     *TokenStruct `json:"token"`
}

type accountFullStruct struct {
	AccountId int32  `json:"account_id" `
	Password  string `json:"password" `
	Username  string `json:"username" `
	Email     string `json:"mail" `
	Firstname string `json:"firstname" `
	Lastname  string `json:"lastname" `
	Role      string `json:"role" `
	Phone     string `json:"phone"`
	Status    string `json:"status"`
	ImageURL  string `Json:"image_path"`
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func LoginStreamingAccount(c *gin.Context) {
	var params handleParamsRouteLoginStruct

	if err := c.ShouldBindJSON(&params); err != nil {
		c.JSON(http.StatusBadRequest, "login failed")
		return
	}
	rows, err := DB.Query("SELECT * FROM accounts WHERE email=?  ", params.Email)
	if err != nil {
		c.JSON(http.StatusBadRequest, "connect table fail")
		return
	}
	defer rows.Close()

	var account responseStructLogin
	var newRowItem accountFullStruct

	for rows.Next() {

		err = rows.Scan(
			&newRowItem.AccountId,
			&newRowItem.Username,
			&newRowItem.Password,
			&newRowItem.Email,
			&newRowItem.Firstname,
			&newRowItem.Lastname,
			&newRowItem.Phone,
			&newRowItem.Role,
			&newRowItem.Status,
			&newRowItem.ImageURL,
		)

		if err != nil {
			c.JSON(http.StatusBadRequest, "error")
			return
		}
	}
	if len(newRowItem.Email) == 0  {
		c.JSON(http.StatusBadRequest, "user not found")
		return
	}

	matchingPassword := CheckPasswordHash(params.Password, newRowItem.Password)

	if !matchingPassword {
		c.JSON(http.StatusBadRequest, "not match password")
		return
	}

	account.AccountId = newRowItem.AccountId
	account.Email = newRowItem.Email
	account.Role = newRowItem.Role
	token, err := CreateToken(uint64(account.AccountId))
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}
	saveErr := CreateAuth(uint64(account.AccountId), token)
	if saveErr != nil {
		c.JSON(http.StatusUnprocessableEntity, saveErr.Error())
	}
	account.Token = &TokenStruct{
		AccessToken: token.AccessToken, RefreshToken: token.RefreshToken,
	}

	c.JSON(http.StatusOK, account)
	
}

func VerifyPINStreamingAccount(c *gin.Context){
	var params handleParamsRoutePINStruct

	if err := c.ShouldBindJSON(&params); err != nil {
		c.JSON(http.StatusBadRequest, "Verify PIN Failed")
		return
	}
}