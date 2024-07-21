package login

import "database/sql"

var DB *sql.DB

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
	// UserID    int32  `json:"idusers" binding:"required"`
	Device    string `json:"is_device"  binding:"required"`
	LastLogin string `json:"last_login"  binding:"required"`
}

type ErrorMessageVerifyPIN struct {
	ErrerCode    int32  `json:"error_code"`
	Message      string `json:"error_message"`
	StayinDevice string `json:"are_logged_in_device"`
}

type ResponseExtractUserToken struct {
	AccountID int `json:"account_id"`
	UserID    int `json:"user_id"`
}

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