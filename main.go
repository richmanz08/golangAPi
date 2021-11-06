package main

import (
	L "api-webapp/Books"
	M "api-webapp/Member"
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	myPrint()

	// emy.Lunchtext()
	// CONNECT DATABASE
	db, err := sql.Open("mysql", "arnonpc:Xx0984437173@@tcp(127.0.0.1:3306)/myproject")
	if err != nil {
		fmt.Println("Connect Database Failed")
		panic(err.Error())

	} else {
		fmt.Println("Connect Database Success")
	}
	defer db.Close()
	// Pass Variable to member.go
	M.DB = db

	// My database api
	r.HandleFunc("/api/memberall", M.GetallMember).Methods("GET")
	// Book api
	r.HandleFunc("/api/books", L.GetBooks).Methods("GET")
	r.HandleFunc("/api/books/{id}", L.GetBook).Methods("GET")
	r.HandleFunc("/api/books", L.CreateBook).Methods("POST")
	r.HandleFunc("/api/books/{id}", L.UpdateBook).Methods("PUT")
	r.HandleFunc("/api/books/{id}", L.DeleteBook).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", r))

}

// คำสั่ง [go run .] เพราะมันต้องเรียก package อื่นด้วย
