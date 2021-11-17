package main

import (
	J "api-webapp/Books"
	L "api-webapp/Login"
	M "api-webapp/Member"
	COM "api-webapp/components"
	"database/sql"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

// var (
// 	router = gin.Default()
// )

func main() {
	router := gin.Default()
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
	//Login Authority Project Virify by jwt token

	router.POST("/login", L.Login)
	router.POST("/logout", L.Logout)
	router.POST("/description", COM.Components)
	// My database api
	router.GET("/memberall", M.GetallMember)
	// r.HandleFunc("/api/memberall", M.GetallMember).Methods("GET")
	// Book api
	router.GET("/books", J.GetBooks)

	// r.HandleFunc("/api/books", L.GetBooks).Methods("GET")
	// r.HandleFunc("/api/books/{id}", L.GetBook).Methods("GET")
	// r.HandleFunc("/api/books", L.CreateBook).Methods("POST")
	// r.HandleFunc("/api/books/{id}", L.UpdateBook).Methods("PUT")
	// r.HandleFunc("/api/books/{id}", L.DeleteBook).Methods("DELETE")

	log.Fatal(router.Run(":8080"))

}

// คำสั่ง [go run .] เพราะมันต้องเรียก package อื่นด้วย
