package main

import (
	L "api-webapp/Login"
	"net/http"

	// M "api-webapp/Member"

	P "api-webapp/another"
	C "api-webapp/cloud"
	COM "api-webapp/components"
	VIDEO "api-webapp/video"
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

	// db, err := sql.Open("mysql", "root:Xx0984437173@@tcp(127.0.0.1:3306)/app_database") //for client macOS
	db, err := sql.Open("mysql", "root:1234@tcp(127.0.0.1:3306)/app_database") //for client PC Window
	if err != nil {
		fmt.Println("Connect Database Failed")
		panic(err.Error())

	} else {
		fmt.Println("Connect Database Success")
	}
	defer db.Close()
	// M.DB = db
	COM.DB = db
	COM.DBmember = db
	C.DB = db
	L.DB = db

	// Pass Variable to member.go

	//Login Authority Project Virify by jwt token

	router.POST("/login", L.Login)
	router.POST("/logout", L.Logout)

	// router.POST("/description", COM.Components)
	// My database api

	router.GET("/testenv", P.TestEnvironment)
	router.GET("/testusetoken", P.TestUseToken)
	router.POST("/upimage-local", P.TestUploadImageOnLocalHost)
	router.StaticFS("/public", http.Dir("public"))
	router.StaticFS("/assets", http.Dir("assets"))

	//#### Cloud Service ####
	// router.POST("/cloud-storage-bucket", C.HandleFileUploadToBucket)
	// router.PUT("/cloud-get-image", C.GetUrlFile)

	//##### Product ####
	// router.GET("/allproduct", COM.ShowAllProduct)
	// router.POST("/addproduct", COM.AddProDuct)
	// router.PUT("/updateproduct", COM.UpdateProduct)
	// router.DELETE("/deletedproduct/:id", COM.DeleteProduct)

	//#### Member #####
	router.POST("/addmember", COM.CreateUser)
	router.GET("/allusers", COM.ShowallUser)
	router.GET("/userbyid/:id", COM.GetUserById)
	router.DELETE("/deluserbyid/:id", COM.DeletedUser)
	router.PUT("/edituser", COM.EditUserById)
	router.GET("/sumofmember", COM.CounterMember)

	//###### video-streaming ######
	// router.GET("/movie",COM.VideoStreamingRender)
	router.GET("/movie/:mID", VIDEO.ServerFileMedia)
	router.GET("/media", VIDEO.ServerURLFileMedia)
	router.GET("/subtitle", VIDEO.ServerURLFileSubtitle)
	router.GET("/thumbnail/:file", VIDEO.ServerFileThumbnail)
	// router.GET("/media/{mId:[0-9]+}/stream/", VIDEO.StreamHandle)

	log.Fatal(router.Run(":8080"))

}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// c.Writer.Header().Set("Content-Type", "application/octet-stream")
		c.Writer.Header().Set("Content-Type", "application/json")
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")
		// c.Writer.Header().Set("Content-Type", "application/octet-stream")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

// คำสั่ง [go run .] เพราะมันต้องเรียก package อื่นด้วย
//หากต้องการ auto run เมื่อ save nodemon --exec go run main.go
//หากต้องการเคลียแคช go clean -cache -modcache -i -r
