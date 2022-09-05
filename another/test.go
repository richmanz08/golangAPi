package another

import (
	Ex "api-webapp/Login"
	MES "api-webapp/Message"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func TestEnvironment(c *gin.Context){
	var envs map[string]string
    envs, err := godotenv.Read(".env")
	if err != nil {
        log.Fatal("Error loading .env file")
    }
	A := envs["PATH_CLOUD_STORAGE"]
	B:= envs["BUCKET_NAME_CLOUD_STORAGE"]
	fmt.Print(A,B)
	c.JSON(http.StatusOK,A+B)
}
func TestUseToken(c *gin.Context) {
	//เช็คว่า มี token ในระบบหรือไม่
	tokenAuth, err := Ex.ExtractTokenMetadata(c.Request)
	fmt.Println("checktoken", tokenAuth)
	if err != nil {
		c.JSON(http.StatusUnauthorized, MES.Token_Error)
	}
	//เช็คหาว่าเป็นใคร
	userId, err := Ex.FetchAuth(tokenAuth)
	fmt.Println("checktoken", userId)
	if err != nil {
		c.JSON(http.StatusUnauthorized, MES.Token_timeout)
	}else{
		// TO DO WORKING API
		c.JSON(http.StatusCreated, MES.Token_Validator)
	}
	
	// Des.UserID = userId

	
}

func TestUploadImageOnLocalHost(c *gin.Context){
	fmt.Println("File Upload Endpoint Hit")
	//รับได้ทั้งข้อมูลและไฟล์
	file,header,err := c.Request.FormFile("file")

	if err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("file err : %s", err.Error()))
		return
	}

	  filename := header.Filename
	  out, err := os.Create("public/" + filename)
	  if err != nil {
		log.Fatal(err)
	  }
	  defer out.Close()
	  _, err = io.Copy(out, file)
	  if err != nil {
		log.Fatal(err)
	  }
	  filepath := "http://localhost:8080/public/" + filename
  c.JSON(http.StatusOK, gin.H{"filepath": filepath})
}