package member

import (
	"database/sql"
	"encoding/json"
	"net/http"
)

type User struct {
	Id        int    `json:"id_member"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"Lastname"`
	Role      string `json:"role"`
	Email     string `json:"email"`
}

var DB *sql.DB

func GetallMember(w http.ResponseWriter, r *http.Request) {
	// Query  data on table
	var AllUser []User
	datamember, err_table_user := DB.Query("SELECT * FROM myproject.user")
	if err_table_user != nil {
		panic(err_table_user.Error())
	}
	defer datamember.Close()

	for datamember.Next() {
		var data User
		err_table_user = datamember.Scan(&data.Id, &data.Firstname, &data.Lastname, &data.Role, &data.Email)
		if err_table_user != nil {
			panic(err_table_user.Error())
		}
		AllUser = append(AllUser, User{Id: data.Id, Firstname: data.Firstname, Lastname: data.Lastname, Role: data.Role, Email: data.Email})

		// show renponse data from table
		// log.Printf(data.Firstame, data.Id, data.Lastname, data.Role)

	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(AllUser)

}
