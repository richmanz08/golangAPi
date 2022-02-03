package main

import (
	L "api-webapp/Login"
	M "api-webapp/Member"

	P "api-webapp/another"
	C "api-webapp/cloud"
	COM "api-webapp/components"
	"database/sql"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	router := gin.Default()
	router.Use(CORSMiddleware())
	// emy.Lunchtext()
	// CONNECT DATABASE

	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/shoppingplatform")
	if err != nil {
		fmt.Println("Connect Database Failed")
		panic(err.Error())

	} else {
		fmt.Println("Connect Database Success")
	}
	defer db.Close()
	M.DB = db
	COM.DB = db
	COM.DBmember = db
	C.DB =db
	L.DB = db

	// Pass Variable to member.go

	//Login Authority Project Virify by jwt token

	router.POST("/login", L.Login)
	router.POST("/logout", L.Logout)

	// router.POST("/description", COM.Components)
	// My database api
	// router.GET("/memberall", M.GetallMember)
	// r.HandleFunc("/api/memberall", M.GetallMember).Methods("GET")
	// Book api
	// router.GET("/books", J.GetBooks)
	// router.POST("/books", J.GetBooks)
	// router.GET("/api/books", J.GetBookById)
	// r.HandleFunc("/api/books", L.CreateBook).Methods("POST")
	// r.HandleFunc("/api/books/{id}", L.UpdateBook).Methods("PUT")
	// r.HandleFunc("/api/books/{id}", L.DeleteBook).Methods("DELETE")
router.GET("/testenv",P.TestEnvironment)
	router.GET("/testusetoken", P.TestUseToken)
	//#### Cloud Service ####
	router.POST("/cloud-storage-bucket",C.HandleFileUploadToBucket)
	router.PUT("/cloud-get-image",C.GetUrlFile)

	//##### Product ####
	router.GET("/allproduct", COM.ShowAllProduct)
	router.POST("/addproduct", COM.AddProDuct)
	router.PUT("/updateproduct", COM.UpdateProduct)
	router.DELETE("/deletedproduct/:id", COM.DeleteProduct)

	//#### Member #####
    router.POST("/addmember",COM.CreateUser)
	router.GET("/allusers",COM.ShowallUser)
	router.GET("/userbyid/:id",COM.GetUserById)
	router.DELETE("/deluserbyid/:id",COM.DeletedUser)
	router.PUT("/edituser",COM.EditUserById)

	log.Fatal(router.Run(":8080"))

}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Content-Type", "application/json")
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

// คำสั่ง [go run .] เพราะมันต้องเรียก package อื่นด้วย
