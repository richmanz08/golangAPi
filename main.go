package main

import (
	L "api-webapp/Login"
	"database/sql"
	"log"
	"net/http"
	"time"

	// M "api-webapp/Member"
	VIDEO "api-webapp/Stream"
	P "api-webapp/another"
	COM "api-webapp/components"
	MOVIE "api-webapp/movies"
	"fmt"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)
type milotabletest struct {
	ID        uint   `gorm:"primarykey"`
	Name      string `gorm:"not null"`
	Email     string `gorm:"unique"`
	Same_email     string `gorm:"unique"`
	Age       int
	CreatedAt time.Time
	UpdatedAt time.Time
	
}
func main() {
	router := gin.Default()
	router.Use(CORSMiddleware())
	// emy.Lunchtext()
	// CONNECT DATABASE

	// db, err := sql.Open("mysql", "root:Xx0984437173@@tcp(127.0.0.1:3306)/app_database") //for client macOS
	dbOld,errConnection := sql.Open("mysql", "root:1234@tcp(127.0.0.1:3306)/app_database?parseTime=true") //for client PC Window
	if errConnection != nil {
		fmt.Println("Connect Database Failed")
		panic(errConnection.Error())

	}
	// by gorm
	db, err := gorm.Open(mysql.Open("root:1234@tcp(127.0.0.1:3306)/app_database?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})
	if err != nil {
		fmt.Println("Connect Database Failed")
		panic(err.Error())

	} else {
		fmt.Println("Connect Database Success")
	}
	// defer db.Close()

	db.AutoMigrate(&milotabletest{})
	db.AutoMigrate(&MOVIE.Movie_information{})

	// M.DB = db
	// COM.DB = db
	// COM.DBmember = db
	// C.DB = db
	L.DB = dbOld
	MOVIE.DB = db

	// Pass Variable to member.go

	//Login Authority Project Virify by jwt token

	// router.POST("/login", L.Login)
	
	router.POST("/logout", L.Logout)
	router.POST("/login-streaming",L.LoginStreamingAccount)
	router.POST("/verify-pin",L.VerifyPINStreamingAccount)
	router.POST("/survive-heal",L.SurviveHeal)
	router.DELETE("/kill-user",L.KillSurvive)
	// router.GET("/checklife-user",L.CheckUserIsSurvive)
	// router.GET("/kill-user",L.KillUserSurvive)
	// 1.1 session
	// router.GET("/save-session",L.CheckInsession)
	// router.GET("/checkout",L.CheckOutSession)
	// router.GET("/checkuser-islogged",L.CheckAreLoggedIN)
	// router.GET("/redis",L.CreateSession)
	// router.POST("/description", COM.Components)
	// My database api

	router.GET("/testenv", P.TestEnvironment)
	router.GET("/testusetoken", P.TestUseToken)
	router.POST("/upimage-local", P.TestUploadImageOnLocalHost)
	router.StaticFS("/public", http.Dir("public"))
	// router.StaticFS("/assets", http.Dir("assets"))

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
	router.GET("/media", VIDEO.ServerURLFileMediaM3U8)
	router.GET("/subtitle", VIDEO.ServerURLFileSubtitle)
	router.GET("/thumbnail/:file", VIDEO.ServerFileThumbnail)
	// router.GET("/media/{mId:[0-9]+}/stream/", VIDEO.StreamHandle)


	//####### movies ##########
	router.POST("/movies-create",MOVIE.CreateMovie)

	log.Fatal(router.Run(":8080"))
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Content-Type", "application/json")
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")
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
// npx kill-port 8080  ##### สำหรับ kill port
// netstat -a ####### สำหรับเช็ค port ที่มีในเครื่อง
