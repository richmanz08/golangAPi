package login

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type responseVerifyPIN struct {
	AccountId int32  `json:"account_id" `
	UserID    int32  `json:"idusers"`
	UserIndex int32  `json:"user_idx"`
	Username  string `json:"username" `
	ImageURL  string `json:"image_url"`
	Expire    bool   `json:"account_is_expire"`
}
type responseVerifyPINformatJWT struct {
	JwtoffVerifyPIN string `json:"access_jwt_pin"`
}

type myUserStruct struct {
	UserIndex int32  `json:"user_idx"`
	Username  string `json:"username" `
	ImageURL  string `json:"image_url"`
}

type userScanTableStruct struct {
	UserID    int32  `json:"idusers" binding:"required"`
	AccountId int32  `json:"account_id" binding:"required"`
	UserIndex int32  `json:"user_idx" binding:"required"`
	Username  string `json:"username" binding:"required"`
	ImageURL  string `json:"image_url" binding:"required"`
	PIN       string `json:"pin" binding:"required"`
}
type handleParamsRouteLoginStruct struct {
	Email    string `json:"email"  binding:"required"`
	Password string `json:"password"  binding:"required"`
}
type handleParamsRoutePINStruct struct {
	AccountId     int32  `json:"account_id" binding:"required"`
	UserIndex     int32  `json:"user_idx"  binding:"required"`
	PIN           string `json:"pin"  binding:"required"`
	ConnectionKey string `json:"force_connection_key"  binding:"required"`
}
type responseStructLogin struct {
	AccountId int32          `json:"account_id" `
	Email     string         `json:"email" `
	Role      string         `json:"role" `
	Reneval   string         `json:"reneval" `
	UserList  []myUserStruct `json:"user_list" `
	Token     *TokenStruct   `json:"token"`
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
	Reneval   string `json:"reneval"`
}

type checkReneval struct {
	Reneval string `json:"reneval" `
}
type SurviveParams struct {
	UserID int32  `json:"idusers" binding:"required"`
	Device string `json:"is_device"  binding:"required"`
}

type ErrorMessageVerifyPIN struct {
	ErrerCode int32 `json:"error_code"`
	Message string `json:"error_message"`
	StayinDevice string `json:"are_logged_in_device"`
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
			&newRowItem.Reneval,
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
	log.Println(newRowItem)
	matchingPassword := ChekHash(params.Password, newRowItem.Password)

	if !matchingPassword {
		c.JSON(http.StatusBadRequest, "not match password")
		return
	}

	userOfAccount, err := QueryDataAllAccount(newRowItem.AccountId)
	if err != nil {
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
	account.Reneval = newRowItem.Reneval
	account.Role = newRowItem.Role
	account.UserList = userOfAccount

	c.JSON(http.StatusOK, account)

}

func VerifyPINStreamingAccount(c *gin.Context) {
	var params handleParamsRoutePINStruct
	var res responseVerifyPINformatJWT
	var Errresponse ErrorMessageVerifyPIN
	
	// 1. handle body params
	if err := c.ShouldBindJSON(&params); err != nil {
		Errresponse.ErrerCode = 100
		Errresponse.Message = "params struct error."
		c.JSON(http.StatusBadRequest,Errresponse)
		return
	}

	// 2. check account is Reneval ?
	checkReneval := QueryGetRenevalofAccount(params.AccountId)

	// 3. query rows in database
	userOfAccount, err := QueryDataAccount(params.AccountId, params.UserIndex)
	if err != nil {
		Errresponse.ErrerCode = 101
		Errresponse.Message = "user is not found is system."
		c.JSON(http.StatusBadRequest, Errresponse)
		return
	}

	// 4. check password is matching ?
	verifyPin := ChekHash(params.PIN, userOfAccount.PIN)
	if !verifyPin {
		Errresponse.ErrerCode = 102
		Errresponse.Message = "pin is not match."
		c.JSON(http.StatusBadRequest, Errresponse)
		return
	}
	convertToint64, _ := strconv.ParseInt(params.ConnectionKey, 10, 0)
	// 5. check if this user exists in the system.
	if params.ConnectionKey != "notkey" && int32(convertToint64) == userOfAccount.UserID {
		// 5.1 check a key connection & verify
		convertToint64, _ := strconv.ParseInt(params.ConnectionKey, 10, 0)
		isSurvive := ChekUserIsSurviveInSystem(int32(convertToint64))
		if !isSurvive {
			Errresponse.ErrerCode = 103
			Errresponse.Message = "Connection key error.Please reset cookie in a browser!"
			c.JSON(http.StatusBadRequest, Errresponse)
			return
		}
	} else {
		// 5.2 not found a key connection from browser
		isSurvive := ChekUserIsSurviveInSystem(userOfAccount.UserID)
		if isSurvive {
			Errresponse.ErrerCode = 104
			Errresponse.Message = "This user is currently on another device."
			Errresponse.StayinDevice = CheckDeviceUserInSystem(userOfAccount.UserID)
			c.JSON(http.StatusBadRequest, Errresponse)
			return
		}
	}

	// 6. create jwt
	jwtOfaccessPINverify, _ := CreateJWTofPIN(responseVerifyPIN{
		AccountId: userOfAccount.AccountId,
		UserID:    userOfAccount.UserID,
		UserIndex: userOfAccount.UserIndex,
		Username:  userOfAccount.Username,
		ImageURL:  userOfAccount.ImageURL,
		Expire:    checkReneval,
	})

	res.JwtoffVerifyPIN = jwtOfaccessPINverify

	// end...
	c.JSON(http.StatusOK, res)
}

func SurviveHeal(c *gin.Context) {
	// set session login user on Redis server ...
	var params SurviveParams
	// 1. handle body params
	if err := c.ShouldBindJSON(&params); err != nil {
		c.JSON(http.StatusBadRequest, "body request failed")
		return
	}
	// 2. set user is survive on redis
	_, era := client.Set(strconv.Itoa(int(params.UserID)), params.Device, 120*time.Minute).Result()
	if era != nil {
		c.JSON(http.StatusBadGateway, "Error created session on redis server")
		return
	}
	// end ...
	c.JSON(http.StatusOK, "create success")
}

func KillSurvive(c *gin.Context) {
	// 1. handle header params
	userID := c.Request.URL.Query().Get("userID")

	// 2. remove user survive in system redis server
	_, err := client.Del(userID).Result()
	if err != nil {
		c.JSON(http.StatusBadRequest, "user not found")
		return
	}
	// end ...
	c.JSON(http.StatusOK, "remove success")
}

// ----------------- Function duplicate Helper common

func ChekUserIsSurviveInSystem(userID int32) bool {
	_, err := client.Get(strconv.Itoa(int(userID))).Result()
	return err == nil
}
func CheckDeviceUserInSystem(userID int32) string{
	value, err := client.Get(strconv.Itoa(int(userID))).Result()
	if err != nil {
		return ""
	}
	return value
}

func Hash(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}
func ChekHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func QueryGetRenevalofAccount(account_id int32) bool {
	var RenevalDate checkReneval
	current_date := time.Now()
	rows, err := DB.Query("SELECT reneval FROM accounts WHERE account_id=?", account_id)
	if err != nil {
		return false
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&RenevalDate.Reneval)
		if err != nil {
			return false
		}
	}
	// loc, _ := time.LoadLocation("UTC")
	// createdAt := time.Now().In(loc).Add(2 * time.Hour)
	// now := time.Now()
	// fmt.Println(now.Format(time.UnixDate))
	// fmt.Println(createdAt)

	date, error := time.Parse("2006-01-02 15:04:05", RenevalDate.Reneval)

	if error != nil {
		fmt.Println(error)
		// return
	}
	// log.Println("Value of date: ", date)
	// myTime, err := time.Parse("2 Jan 06 03:04PM", "10 Nov 10 11:00PM")
	// 	myTime, err := time.Parse("2023-01-01 21:01:22", RenevalDate.Reneval)
	// if err != nil {
	// 	panic(err)
	// }
	// now := time.Now()
	// fmt.Println(current_date.Before(date))
	// 	log.Println("TIME NOW :::",current_date.Format("2006-01-02 15:04:05"))
	// 	log.Println("RENEVAL CURRENT :::",date.Add(time.Hour*720))
	// fmt.Println("Expire is",current_date.After(date.Add(time.Hour*720)))
	// log.Println("TIME NOW :::",current_date.Format("2006-01-02 15:04:05"))
	// log.Println("RENEVAL CURRENT :::",date)
	return current_date.After(date.Add(time.Hour * 720))
}

func QueryDataAccount(account_id, usr_idx int32) (*userScanTableStruct, error) {
	var Users userScanTableStruct
	rows, err := DB.Query("SELECT * FROM users WHERE account_id=? and usr_idx=?", account_id, usr_idx)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(
			&Users.UserID,
			&Users.AccountId,
			&Users.UserIndex,
			&Users.Username,
			&Users.ImageURL,
			&Users.PIN,
		)
		if err != nil {
			return nil, err
		}
	}

	return &userScanTableStruct{
		UserID:    Users.UserID,
		AccountId: Users.AccountId,
		UserIndex: Users.UserIndex,
		Username:  Users.Username,
		ImageURL:  Users.ImageURL,
		PIN:       Users.PIN,
	}, nil
}

func QueryDataAllAccount(account_id int32) ([]myUserStruct, error) {
	var Users userScanTableStruct
	var newArray []myUserStruct
	rows, err := DB.Query("SELECT * FROM users WHERE account_id=?", account_id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(
			&Users.UserID,
			&Users.AccountId,
			&Users.UserIndex,
			&Users.Username,
			&Users.ImageURL,
			&Users.PIN,
		)
		if err != nil {
			return nil, err
		}
		newArray = append(newArray, myUserStruct{UserIndex: Users.UserIndex, Username: Users.Username, ImageURL: Users.ImageURL})
	}
	return newArray, nil
}

func CreateJWTofPIN(data responseVerifyPIN) (string, error) {
	optionJWT := jwt.MapClaims{}
	optionJWT["account_id"] = data.AccountId
	optionJWT["user_idx"] = data.UserIndex
	optionJWT["idusers"] = data.UserID
	optionJWT["username"] = data.Username
	optionJWT["image_url"] = data.ImageURL
	optionJWT["account_is_expire"] = data.Expire
	groupOption := jwt.NewWithClaims(jwt.SigningMethodHS256, optionJWT)

	jwt, err := groupOption.SignedString([]byte(os.Getenv("ACCESS_SECRET")))
	if err != nil {
		return err.Error(), err
	}
	return jwt, nil
}

//save concept code
// func VerifyPINStreamingAccount(c *gin.Context) {
// 	var params handleParamsRoutePINStruct
// 	// var matchUser userJsonStructofJWT
// 	// var stopedLoop bool = false
// 	var res responseVerifyPINformatJWT

// 	message_error1 := "Verify PIN Failed"

// 	if err := c.ShouldBindJSON(&params); err != nil {
// 		c.JSON(http.StatusBadRequest, message_error1)
// 		return
// 	}

// 	// go func() {
// 		checkReneval :=  QueryGetRenevalofAccount(params.AccountId)
// 		// fmt.Println(checkReneval)
// 	// }()

// 	userOfAccount, err := QueryDataAccount(params.AccountId, params.UserIndex)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, err)
// 		return
// 	}
// 	// fmt.Println(Hash(params.PIN))
// 	verifyPin := ChekHash(params.PIN, userOfAccount.PIN)
// 	if !verifyPin {
// 		c.JSON(http.StatusBadRequest, " PIN not match")
// 		return
// 	}
// 	// jwtUser := userOfAccount.UserJWT
// 	// claims := jwt.MapClaims{}
// 	// jsonJWT, err := jwt.ParseWithClaims(jwtUser, claims, func(token *jwt.Token) (interface{}, error) {
// 	// 	return []byte(MY_APPLICATION_JWT_KEY), nil
// 	// })
// 	// if err != nil {
// 	// 	c.JSON(http.StatusBadRequest, message_error1)
// 	// 	return
// 	// }

// 	// claims, okr := jsonJWT.Claims.(jwt.MapClaims)

// 	// if !okr {
// 	// 	fmt.Println("error okr")
// 	// }
// 	// listUserArray := claims["user_list"].([]interface{})

// 	// for _, item := range listUserArray {
// 	// 	if !stopedLoop {

// 	// 		var info userJsonStructofJWT
// 	// 		if rec, ok := item.(map[string]interface{}); ok {
// 	// 			for key, val := range rec {
// 	// 				if key == "usr_idx" {
// 	// 					info.UserIndex = fmt.Sprintf("%v", val)
// 	// 				} else if key == "username" {
// 	// 					info.Username = fmt.Sprintf("%v", val)
// 	// 				} else if key == "pin" {
// 	// 					info.PIN = fmt.Sprintf("%v", val)
// 	// 				}else if  key == "image_url"{
// 	// 					info.ImageURL = fmt.Sprintf("%v", val)
// 	// 				}
// 	// 			}
// 	// 		}
// 	// 		var a bool = info.UserIndex == params.UserIndex
// 	// 		var b bool = info.PIN == params.PIN
// 	// 		if a && b {
// 	// 			matchUser.PIN = info.PIN
// 	// 			matchUser.Username = info.Username
// 	// 			matchUser.UserIndex = info.UserIndex
// 	// 			matchUser.ImageURL = info.ImageURL
// 	// 			stopedLoop = true

// 	// 		}
// 	// 	}

// 	// }

// 	// if len(matchUser.Username) == 0 {
// 	// 	c.JSON(http.StatusBadRequest, message_error1)
// 	// 	return
// 	// }
//  jwtOfaccessPINverify,err :=	CreateJWTofPIN(responseVerifyPIN{
// 		AccountId: userOfAccount.AccountId,
// 		UserIndex: userOfAccount.UserIndex,
// 		Username: userOfAccount.Username,
// 		ImageURL: userOfAccount.ImageURL,
// 		Expire:checkReneval,
// 	})
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, "Generate JWT Failed")
// 		return
// 	}
// 	res.JwtoffVerifyPIN = jwtOfaccessPINverify
// 	// res.AccountId = userOfAccount.AccountId
// 	// res.UserIndex = userOfAccount.UserIndex
// 	// res.Username = userOfAccount.Username
// 	// res.ImageURL = userOfAccount.ImageURL
// 	c.JSON(http.StatusOK, res)
// }

// for make sessions verifyPin
// https://gowebexamples.com/sessions
