package login

import (
	"fmt"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)
type responseVerifyPIN struct {
	AccountId int32   `json:"account_id" `
	UserIndex string `json:"user_idx"`
	Username  string `json:"username" `
	ImageURL  string `json:"image_url"`
}
type userJsonStructofJWT struct {
	UserIndex string `json:"user_idx"`
	Username  string `json:"username" `
	PIN       string `json:"pin"`
	ImageURL  string `json:"image_url"`
}

type userScanTableStruct struct {
	UserID int32 `json:"idusers" binding:"required"`
	AccountId int32  `json:"account_id" binding:"required"`
	UserJWT string  `json:"users_jwt" binding:"required"`
}
type handleParamsRouteLoginStruct struct {
	Email    string `json:"email"  binding:"required"`
	Password string `json:"password"  binding:"required"`
}
type handleParamsRoutePINStruct struct {
	AccountId int32  `json:"account_id" binding:"required"`
	UserIndex string `json:"user_idx"  binding:"required"`
	PIN       string `json:"pin"  binding:"required"`
}
type responseStructLogin struct {
	AccountId int32        `json:"account_id" `
	Email     string       `json:"email" `
	Role      string       `json:"role" `
	UserJWT string  `json:"users_jwt"`
	Token     *TokenStruct `json:"token"`
}
 
type accountFullStruct struct {
	AccountId int32  `json:"account_id" `
	Password  string `json:"password" `
	Email     string `json:"mail" `
	Firstname string `json:"firstname" `
	Lastname  string `json:"lastname" `
	Role      string `json:"role" `
	Phone     string `json:"phone"`
	Status    string `json:"status"`
}



func LoginStreamingAccount(c *gin.Context) {
	var params handleParamsRouteLoginStruct
	var account responseStructLogin
	var newRowItem accountFullStruct


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


	for rows.Next() {

		err = rows.Scan(
			&newRowItem.AccountId,
			&newRowItem.Password,
			&newRowItem.Email,
			&newRowItem.Firstname,
			&newRowItem.Lastname,
			&newRowItem.Phone,
			&newRowItem.Role,
			&newRowItem.Status,
		)

		if err != nil {
			c.JSON(http.StatusBadRequest, "error")
			return
		}
	}
	if len(newRowItem.Email) == 0 {
		c.JSON(http.StatusBadRequest, "user not found")
		return
	}

	matchingPassword := CheckPasswordHash(params.Password, newRowItem.Password)

	if !matchingPassword {
		c.JSON(http.StatusBadRequest, "not match password")
		return
	}


	userOfAccount,err := QueryDataAccount(newRowItem.AccountId);
	if err != nil{
		c.JSON(http.StatusBadRequest, err)
		return	
	} 


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

	account.AccountId = newRowItem.AccountId
	account.Email = newRowItem.Email
	account.Role = newRowItem.Role
	account.UserJWT = userOfAccount.UserJWT

	c.JSON(http.StatusOK, account)

}

func VerifyPINStreamingAccount(c *gin.Context) {
	var params handleParamsRoutePINStruct
	var matchUser userJsonStructofJWT
	var stopedLoop bool = false
	var res responseVerifyPIN

	message_error1 := "Verify PIN Failed"

	if err := c.ShouldBindJSON(&params); err != nil {
		c.JSON(http.StatusBadRequest, message_error1)
		return
	}

	userOfAccount,err := QueryDataAccount(params.AccountId);
	if err != nil{
		c.JSON(http.StatusBadRequest, err)
		return	
	} 

	jwtUser := userOfAccount.UserJWT
	claims := jwt.MapClaims{}
	jsonJWT, err := jwt.ParseWithClaims(jwtUser, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(MY_APPLICATION_JWT_KEY), nil
	})
	if err != nil {
		c.JSON(http.StatusBadRequest, message_error1)
		return
	}

	claims, okr := jsonJWT.Claims.(jwt.MapClaims)

	if !okr {
		fmt.Println("error okr")
	}
	listUserArray := claims["user_list"].([]interface{})


	for _, item := range listUserArray {
		if !stopedLoop {
		
			var info userJsonStructofJWT
			if rec, ok := item.(map[string]interface{}); ok {
				for key, val := range rec {
					if key == "usr_idx" {
						info.UserIndex = fmt.Sprintf("%v", val)
					} else if key == "username" {
						info.Username = fmt.Sprintf("%v", val)
					} else if key == "pin" {
						info.PIN = fmt.Sprintf("%v", val)
					}else if  key == "image_url"{
						info.ImageURL = fmt.Sprintf("%v", val)
					}
				}
			}
			var a bool = info.UserIndex == params.UserIndex
			var b bool = info.PIN == params.PIN
			if a && b {
				matchUser.PIN = info.PIN
				matchUser.Username = info.Username
				matchUser.UserIndex = info.UserIndex
				matchUser.ImageURL = info.ImageURL
				stopedLoop = true
	
			}
		}

	}

	if len(matchUser.Username) == 0 {
		c.JSON(http.StatusBadRequest, message_error1)
		return
	}


	
	res.AccountId = params.AccountId
	res.UserIndex = matchUser.UserIndex
	res.Username = matchUser.Username
	res.ImageURL = matchUser.ImageURL
	c.JSON(http.StatusOK, res)
}


// ----------------- Function duplicate Helper common 

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func QueryDataAccount(account_id int32)(*userScanTableStruct,error){
	var Users userScanTableStruct
	rows, err := DB.Query("SELECT * FROM users WHERE account_id=?", account_id)
	if err != nil {
		return  nil,err
	}
	defer rows.Close()

	for rows.Next(){
		err = rows.Scan(
			&Users.UserID,
			&Users.AccountId,
			&Users.UserJWT,
		)		
		if err != nil {
			return nil,err
		}
	}

	if len(Users.UserJWT) == 0{
		return nil,err
	}

	return &userScanTableStruct{
		UserID :Users.UserID,
		AccountId :Users.AccountId,
		UserJWT:Users.UserJWT,
	},nil
}
