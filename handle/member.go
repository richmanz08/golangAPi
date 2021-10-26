package handle

type User struct {
	Id       int    `json:"id_member"`
	Firstame string `json:"firstname"`
	Lastname string `json:"Lastname"`
	Role     string `json:"role"`
	Email    string `json:"email"`
}

// var allUser []User

// func getallMember(w http.ResponseWriter, r *http.Request) {
// 	json.NewEncoder(w).Encode(allUser)
// }
