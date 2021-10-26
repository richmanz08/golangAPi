package main

import (
	L "api-webapp/func"
	"api-webapp/handle"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

type Book struct {
	ID     string  `json:"id"`
	Isbn   string  `json:"isbn"`
	Title  string  `json:"title"`
	Author *Author `json:"author"`
}

type Author struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"Lastname"`
}

var allUser []handle.User

func getallMember(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(allUser)
}

var books []Book

// Get All Books
func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

// Get Single Books
func getBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) //Get params
	// Loop through books and find with id
	for _, item := range books {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Book{})
}

// Create a New Book
func createBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var book Book
	_ = json.NewDecoder(r.Body).Decode(&book)
	book.ID = strconv.Itoa(rand.Intn(100000)) //Mock Id - not safe
	books = append(books, book)
	json.NewEncoder(w).Encode(book)

} //Update a New Book
func updateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range books {
		if item.ID == params["id"] {
			books = append(books[:index], books[index+1:]...)
			var book Book
			_ = json.NewDecoder(r.Body).Decode(&book)
			book.ID = params["id"]
			books = append(books, book)
			json.NewEncoder(w).Encode(book)
			return

		}
	}
	json.NewEncoder(w).Encode(books)
}

// Delete a New Book
func deleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range books {
		if item.ID == params["id"] {
			books = append(books[:index], books[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(books)
}

func main() {
	r := mux.NewRouter()
	myPrint()
	L.Another()

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

	// Query  data on table

	datamember, err_table_user := db.Query("SELECT * FROM myproject.user")
	if err_table_user != nil {
		panic(err.Error())
	}
	for datamember.Next() {
		var data handle.User
		err = datamember.Scan(&data.Id, &data.Firstame, &data.Lastname, &data.Role, &data.Email)
		if err != nil {
			panic(err.Error())
		}
		allUser = append(allUser, handle.User{Id: data.Id, Firstame: data.Firstame, Lastname: data.Lastname, Role: data.Role, Email: data.Email})
		// show renponse data from table
		// log.Printf(data.Firstame, data.Id, data.Lastname, data.Role)

	}

	//Mock Data - @todo - implement DB
	// books = append(books, Book{ID: "1", Isbn: "448743", Title: "Book one", Author: &Author{Firstname: "JARD", Lastname: "scumdown"}})
	// books = append(books, Book{ID: "2", Isbn: "448744", Title: "Book two", Author: &Author{Firstname: "Smith", Lastname: "scumdown"}})
	// My database api
	r.HandleFunc("/api/memberall", getallMember).Methods("GET")
	// Book api
	r.HandleFunc("/api/books", getBooks).Methods("GET")
	r.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	r.HandleFunc("/api/books", createBook).Methods("POST")
	r.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")
	r.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", r))

}

// คำสั่ง [go run .] เพราะมันต้องเรียก package อื่นด้วย
