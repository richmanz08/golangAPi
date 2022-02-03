package login

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

var DB *sql.DB

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
type memberFullStruct struct {
	AccountId int32        `json:"account_id" `
	Username  string       `json:"username" `
	Password  string       `json:"password" `
	Mail      string       `json:"mail" `
	Name      string       `json:"name" `
	Surname   string       `json:"surname" `
	Phone     string       `json:"phone"`
	Role      string       `json:"role" `
	Token     *TokenStruct `json:"token"`
}
type loginParamsStruct struct {
	Username string `json:"username"  binding:"required"`
	Password string `json:"password"  binding:"required"`
}

// var user = User{
// 	ID:       1,
// 	Username: "username",
// 	Password: "password",
// 	Phone:    "49123454322", //this is a random number
// }

// func Login(c *gin.Context) {
// 	var userData []Userdetail
// 	var u User
// 	if err := c.ShouldBindJSON(&u); err != nil {
// 		c.JSON(http.StatusUnprocessableEntity, "Invalid json provided")
// 		return
// 	}
// 	//compare the user from the request, with the one we defined:
// 	if user.Username != u.Username || user.Password != u.Password {
// 		c.JSON(http.StatusUnauthorized, "Please provide valid login details")
// 		return
// 	}
// 	ts, err := CreateToken(user.ID)
// 	if err != nil {
// 		c.JSON(http.StatusUnprocessableEntity, err.Error())
// 		return
// 	}
// 	saveErr := CreateAuth(user.ID, ts)
// 	if saveErr != nil {
// 		c.JSON(http.StatusUnprocessableEntity, saveErr.Error())
// 	}
// 	// tokens := map[string]string{
// 	// 	"access_token":  ts.AccessToken,
// 	// 	"refresh_token": ts.RefreshToken,
// 	// }

// 	userData = append(userData, Userdetail{ID: 999, Age: 22, Username: "rizhmanz08", Firstname: "arnon", Lastname: "Rungrueng", Role: "admin",
// 		Token: &TokenStruct{
// 			AccessToken: ts.AccessToken, RefrshToken: ts.RefreshToken,
// 		}})

// 	c.JSON(http.StatusOK, userData)
// }

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func Login(c *gin.Context) {
	var m memberFullStruct
	var u loginParamsStruct

	if err := c.ShouldBindJSON(&u); err != nil {
		c.JSON(http.StatusUnprocessableEntity, "Invalid json provided")
		return
	}
	data, err := DB.Query("SELECT * FROM members WHERE username=? ", u.Username)
	if err != nil {
		fmt.Println("connect table fail")
	}
	defer data.Close()

	for data.Next() {
		var new memberFullStruct
		err = data.Scan(&new.AccountId,
			&new.Username,
			&new.Password,
			&new.Mail,
			&new.Name,
			&new.Surname,
			&new.Phone,
			&new.Role)

		if err != nil {
			panic(err.Error())
		}

		if new.Username == u.Username {
			MatchingPassword := CheckPasswordHash(u.Password, new.Password)
			// fmt.Print(match)
			if MatchingPassword == true {
				m.AccountId = new.AccountId
				m.Username = new.Username
				m.Password = new.Password
				m.Mail = new.Mail
				m.Name = new.Name
				m.Surname = new.Surname
				m.Phone = new.Phone
				m.Role = new.Role
				ts, err := CreateToken(uint64(m.AccountId))
				if err != nil {
					c.JSON(http.StatusUnprocessableEntity, err.Error())
					return
				}
				saveErr := CreateAuth(uint64(m.AccountId), ts)
				if saveErr != nil {
					c.JSON(http.StatusUnprocessableEntity, saveErr.Error())
				}
				m.Token = &TokenStruct{
					AccessToken: ts.AccessToken, RefreshToken: ts.RefreshToken,
				}

				c.JSON(http.StatusOK, m)
				
			}else{
				c.JSON(http.StatusBadRequest, "login failed")
			}

		}

	}

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


// การแปลงข้อมูลที่ได้มาเป็น json
// resRp := &rp 
// resRp2, _ := json.Marshal(resRp)
//  fmt.Println("Rp", string(resRp2))