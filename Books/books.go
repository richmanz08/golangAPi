package books

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
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
type userImplement struct {
	Id   int    `json:"id" binding:"required"`
	Name string `json:"name" binding:"required"`
}

// Get All Books
func GetBooks(c *gin.Context) {
	var req userImplement
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, "error get method post Structure failed")
		return
	} else {
		fmt.Println("arnontest", req.Id, req.Name)
		// show data from params
		var books []Book

		//Mock Data - @todo - implement DB
		books = append(books, Book{ID: "1", Isbn: "448743", Title: "Book one", Author: &Author{Firstname: "JARD", Lastname: "scumdown"}})
		books = append(books, Book{ID: "2", Isbn: "448744", Title: "Book two", Author: &Author{Firstname: "Smith", Lastname: "scumdown"}})
		books = append(books, Book{ID: "3", Isbn: "448745", Title: "Book tree", Author: &Author{Firstname: "header", Lastname: "scumdown"}})
		// c.JSON(http.StatusOK, gin.H{"data": books})
		c.JSON(http.StatusOK, books)
	}

}

// Get Single Books
// func GetBook(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	params := mux.Vars(r) //Get params
// 	// Loop through books and find with id
// 	for _, item := range books {
// 		if item.ID == params["id"] {
// 			json.NewEncoder(w).Encode(item)
// 			return
// 		}
// 	}
// 	json.NewEncoder(w).Encode(&Book{})
// }

// // Create a New Book
// func CreateBook(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	var book Book
// 	_ = json.NewDecoder(r.Body).Decode(&book)
// 	book.ID = strconv.Itoa(rand.Intn(100000)) //Mock Id - not safe
// 	books = append(books, book)
// 	json.NewEncoder(w).Encode(book)

// } //Update a New Book
// func UpdateBook(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	params := mux.Vars(r)
// 	for index, item := range books {
// 		if item.ID == params["id"] {
// 			books = append(books[:index], books[index+1:]...)
// 			var book Book
// 			_ = json.NewDecoder(r.Body).Decode(&book)
// 			book.ID = params["id"]
// 			books = append(books, book)
// 			json.NewEncoder(w).Encode(book)
// 			return

// 		}
// 	}
// 	json.NewEncoder(w).Encode(books)
// }

// // Delete a New Book
// func DeleteBook(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	params := mux.Vars(r)
// 	for index, item := range books {
// 		if item.ID == params["id"] {
// 			books = append(books[:index], books[index+1:]...)
// 			break
// 		}
// 	}
// 	json.NewEncoder(w).Encode(books)
// }
