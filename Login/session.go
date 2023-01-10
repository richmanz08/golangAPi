package login

import (
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/sessions"
)

var (
	// key must be 16, 24 or 32 bytes long (AES-128, AES-192 or AES-256)
	key   = []byte(os.Getenv("SESSION_COOKIE"))
	store = sessions.NewCookieStore(key)
)

type UserSession struct {
	UserID    int32  `json:"idusers"`
	Device string `json:"is_device" `
	Online     bool `json:"is_online"`
}

var UserArray []UserSession

func CheckInsession(c *gin.Context){
	userID := c.Request.URL.Query().Get("uID")
	log.Println("UserID :::",userID)
	session, _ := store.Get(c.Request,userID)
	var User UserSession
	
	
	newID ,err:=strconv.ParseInt(userID,0,32) 
	if err != nil{
		c.Status(http.StatusBadRequest)
		return
	}
	result := int32(newID)
	User.UserID =result
	User.Device = "Window x64 Chrome"
	User.Online = true
	session.Values["UserID"] = userID
	session.Values["Device"] = "Window x64 Chrome"
	session.Values["Online"] = true
    session.Save(c.Request, c.Writer)


	UserArray = append(UserArray, UserSession{UserID:result,Device:"Window x64 Chrome",Online: true  })
	c.JSON(http.StatusOK,User)
}
func CheckAreLoggedIN(c *gin.Context) {
	log.Println("LIST ALL:::",UserArray)
	var User UserSession
	userID := c.Request.URL.Query().Get("uID")
    session, _ := store.Get(c.Request, userID)
	// log.Println(session)
    // Check if user is authenticated
     _, ok := session.Values["UserID"]; 
	 if !ok{
		c.Status(http.StatusBadRequest)
	  return
   }
	 device, ok := session.Values["Device"]; 
	 if !ok{
		c.Status(http.StatusBadRequest)
	  return
   }
	 online,ok := session.Values["Online"]; 
	 if !ok{
		  c.Status(http.StatusBadRequest)
        return
	 }
	
	 newID ,err:=strconv.ParseInt(userID,0,32) 
	 if err != nil{
		 c.Status(http.StatusBadRequest)
		 return
	 }
	 result := int32(newID)
	 User.UserID =result
	User.Device = device.(string)
	User.Online = online.(bool)
	c.JSON(http.StatusOK,UserArray)
}
func CheckOutSession(c *gin.Context) {
	userID := c.Request.URL.Query().Get("uID")
    session, _ := store.Get(c.Request, userID)
	// log.Println(session)
    // Revoke users authentication
    session.Values["Online"] = false
    session.Save(c.Request, c.Writer)
	c.JSON(http.StatusOK,"is logout session :::"+userID)
}