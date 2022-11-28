package video

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func ServerFileM3U8(c *gin.Context) {
	file_name := c.Param("filename")
	fmt.Println("Filename was connected : ",file_name)
	fileRoot := "public/"
	file, err := os.Open(fileRoot+file_name)
	fmt.Println("File location is : ",fileRoot+file_name)
	if err != nil {
		c.Status(http.StatusNotFound)
		return
	}
	defer file.Close()
	// stFile, err := file.Stat()
	data := make([]byte, 100)
	count, err := file.Read(data)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(count)
	// wrFile := os.WriteFile(stFile)
	

	// c.File(file)
	// c.Status(http.StatusOK,file)


}