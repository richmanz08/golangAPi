package components

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

var DBmember *sql.DB

type memberStruct struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Email     string `json:"email" binding:"required"`
	FirstName     string `json:"firstname" binding:"required"`
	LastName  string `json:"lastname" binding:"required"`
	Phone    string `json:"phone" binding:"required"`
	Role     string `json:"role" binding:"required"`
	Status   string `json:"status" binding:"required"`
}
type memberFullStruct struct {
	AccountId int32  `json:"account_id" `
	Username  string `json:"username" `
	Password  string `json:"password" `
	Email      string `json:"email" `
	FirstName      string `json:"firstname" `
	LastName   string `json:"lastname" `
	Phone     string `json:"phone"`
	Role      string `json:"role" `
	Status   string `json:"status"`
}
func HashPassword(password string)(string, error){
	bytes, err := bcrypt.GenerateFromPassword([]byte(password),14)
	return string(bytes), err
}
func CheckPasswordHash(password,hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
func ShowallUser(c *gin.Context) {

	var member []memberFullStruct
	data, err := DB.Query("SELECT * FROM members")
	if err != nil {
		fmt.Println(err)
	} else {
		for data.Next() {
			var new memberFullStruct
			err = data.Scan(&new.AccountId,
				&new.Username,
				&new.Password,
				&new.Email,
				&new.FirstName,
				&new.LastName,
				&new.Phone,
				&new.Role,
				&new.Status,
			)
			if err != nil {
				panic(err.Error())
			}
			member = append(member,
				memberFullStruct{AccountId: new.AccountId,
					Username: new.Username,
					Password: new.Password,
					Email:     new.Email,
					FirstName:     new.FirstName,
					LastName:  new.LastName,
					Phone:    new.Phone,
					Role:     new.Role,
					Status:new.Status,
				})
		}
		c.JSON(http.StatusOK, member)
	}
	defer data.Close()
}
func GetUserById(c *gin.Context) {
	itemid := c.Param("id")
	var member memberFullStruct
	data, err := DB.Query("SELECT * FROM members WHERE account_id=?", itemid)
	if err != nil {
		fmt.Println(err)
	}
	defer data.Close()

	for data.Next() {
		var new memberFullStruct
		err = data.Scan(&new.AccountId,
			&new.Username,
			&new.Password,
			&new.Email,
			&new.FirstName,
			&new.LastName,
			&new.Phone,
			&new.Role,
			&new.Status,
		)
		if err != nil {
			panic(err.Error())
		}
		member.AccountId = new.AccountId
		member.Username = new.Username
		member.Password = new.Password
		member.Email = new.Email
		member.FirstName = new.FirstName
		member.LastName = new.LastName
		member.Phone = new.Phone
		member.Role = new.Role
		member.Status = new.Status
	}
	c.JSON(http.StatusOK, member)

}

func EditUserById(c *gin.Context) {
	var member memberFullStruct
	if err := c.ShouldBindJSON(&member); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(member)
	data, err := DB.Prepare("UPDATE members SET account_id=?,username=?,email=?,firstname=?,lastname=?,phone=?,role=?,status=? WHERE account_id=?")
	if err != nil {
		fmt.Println(err)
	} else {
		if _, err := data.Exec(member.AccountId, member.Username,member.Email, member.FirstName, member.LastName, member.Phone, member.Role,member.Status, member.AccountId); err != nil {
			fmt.Println("update failed")
			c.JSON(http.StatusBadRequest, "### update failed ### ")
		} else {
			c.JSON(http.StatusOK, member)
		}
	}
	defer data.Close()
}
func CreateUser(c *gin.Context) {

	var member memberStruct
	var memberDetail memberStruct
	if err := c.ShouldBindJSON(&member); err != nil {
		c.JSON(http.StatusBadRequest, "### Structure params failed ###")
	} else {
		data, err := DB.Prepare("INSERT INTO members(username,password,email,firstname,lastname,phone,role,status) VALUES(?,?,?,?,?,?,?,?)")
		
		if err != nil {
			c.JSON(http.StatusBadRequest, "### Insert table failed ###")
		} else {
          
			hash, _ := HashPassword(member.Password)
			data.Exec(member.Username, hash, member.Email, member.FirstName, member.LastName, member.Phone, member.Role,member.Status)
			memberDetail.Username = member.Username
			memberDetail.Password = hash
			memberDetail.Email = member.Email
			memberDetail.FirstName = member.FirstName
			memberDetail.LastName = member.LastName
			memberDetail.Phone = member.Phone
			memberDetail.Role = member.Phone
			memberDetail.Status = member.Status						
			c.JSON(http.StatusCreated, memberDetail)
		}
		defer data.Close()
	}

}
func DeletedUser(c *gin.Context) {
	itemid := c.Param("id")

	data, err := DB.Prepare("DELETE FROM members WHERE account_id=?")
	if err != nil {
		fmt.Println(err)
		return
	} else {
		if _, err := data.Exec(itemid); err != nil {
			fmt.Println("smt.Exec failed: ", err)
		}
		c.JSON(http.StatusOK, "deleted user success")
	}
	defer data.Close()

}
// สอนอัพโหลดไฟล์นะ
// https://tutorialedge.net/golang/go-file-upload-tutorial/ 