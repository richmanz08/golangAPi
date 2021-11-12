package login

import (
	"encoding/json"
	"net/http"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// var books []Book

type User struct {
	ID        int    `json:"id"`
	Age       int    `json:"age"`
	Username  string `json:"username"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Role      string `json:"role"`
	Token     string `json:"token"`
}

// type Book struct {
// 	ID     string       `json:"id"`
// 	Isbn   string       `json:"isbn"`
// 	Title  string       `json:"title"`
// 	Author *loginStruct `json:"author"`
// }
type login struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"Lastname"`
}

func Authorityfunc(w http.ResponseWriter, r *http.Request) {
	var userData []User
	w.Header().Set("Content-Type", "application/json")
	var user login
	_ = json.NewDecoder(r.Body).Decode((&user))
	if user.Firstname != "" && user.Lastname != "" {
		token, err := GenerateToken(999)
		if err != nil {
			http.Error(w, "token generate fail", 400)
		}
		userData = append(userData, User{ID: 999, Age: 22, Username: "richmanz", Firstname: "arnon", Lastname: "ruengrueang", Role: "Administrator", Token: token})
		json.NewEncoder(w).Encode(userData)
	} else {

		http.Error(w, "Structure Failed", 400)
		return
	}

}
func GenerateToken(userId uint64) (string, error) {
	var err error
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["user_id"] = userId
	atClaims["exp"] = time.Now().Add(time.Minute * 15).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString([]byte(os.Getenv("ACCESS_SECRET")))
	if err != nil {
		return "", err
	}
	return token, nil
}
