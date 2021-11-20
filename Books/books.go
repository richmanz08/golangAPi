package books

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Book struct {
	Id     int     `json:"id"`
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

var CategoryBook []Book

func StringToInt(s string) (i int) {
	i, _ = strconv.Atoi(s)
	return
}

// Get All Books
func GetBooks(c *gin.Context) {
	var books []Book
	var req userImplement
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, "error get method post Structure failed")
		return
	} else {
		// fmt.Println("arnontest", req.Id, req.Name)
		// show data from params

		//Mock Data - @todo - implement DB
		books = append(books, Book{Id: 1, Isbn: "448743", Title: "Book one", Author: &Author{Firstname: "JARD", Lastname: "scumdown"}})
		books = append(books, Book{Id: 2, Isbn: "448744", Title: "Book two", Author: &Author{Firstname: "Smith", Lastname: "scumdown"}})
		books = append(books, Book{Id: 3, Isbn: "448745", Title: "Book tree", Author: &Author{Firstname: "header", Lastname: "scumdown"}})
		// c.JSON(http.StatusOK, gin.H{"data": books})
		c.JSON(http.StatusOK, books)
	}
	CategoryBook = books

}

// Get book by id From Path http://localhost:8080/api/books?id=1
func GetBookById(c *gin.Context) {
	idFrompath := c.Request.URL.Query().Get("id")

	// c.String(http.StatusOK, "First ID: "+idFrompath+"\n")
	// Loop through books and find with id
	newId := StringToInt(idFrompath)
	// idInteger := strconv.AppendInt([]byte(idFrompath))
	fmt.Println("ID test", newId)
	for _, item := range CategoryBook {
		if item.Id == newId {
			c.JSON(http.StatusOK, item)
			return
		}

	}
	c.JSON(http.StatusNotFound, "not found item Books id = "+idFrompath)

	// strVar := "100"
	// intVar, err := strconv.Atoi(strVar)
	// fmt.Println(intVar, err, reflect.TypeOf(intVar))

}

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
