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
	Mail     string `json:"mail" binding:"required"`
	Name     string `json:"name" binding:"required"`
	Surname  string `json:"surname" binding:"required"`
	Phone    string `json:"phone" binding:"required"`
	Role     string `json:"role" binding:"required"`
}
type memberFullStruct struct {
	AccountId int32  `json:"account_id" `
	Username  string `json:"username" `
	Password  string `json:"password" `
	Mail      string `json:"mail" `
	Name      string `json:"name" `
	Surname   string `json:"surname" `
	Phone     string `json:"phone"`
	Role      string `json:"role" `
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
				&new.Mail,
				&new.Name,
				&new.Surname,
				&new.Phone,
				&new.Role)
			if err != nil {
				panic(err.Error())
			}
			member = append(member,
				memberFullStruct{AccountId: new.AccountId,
					Username: new.Username,
					Password: new.Password,
					Mail:     new.Mail,
					Name:     new.Name,
					Surname:  new.Surname,
					Phone:    new.Phone,
					Role:     new.Role})
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
			&new.Mail,
			&new.Name,
			&new.Surname,
			&new.Phone,
			&new.Role)
		if err != nil {
			panic(err.Error())
		}
		member.AccountId = new.AccountId
		member.Username = new.Username
		member.Password = new.Password
		member.Mail = new.Mail
		member.Name = new.Name
		member.Surname = new.Surname
		member.Phone = new.Phone
		member.Role = new.Role
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
	data, err := DB.Prepare("UPDATE members SET account_id=?,username=?,mail=?,name=?,surname=?,phone=?,role=? WHERE account_id=?")
	if err != nil {
		fmt.Println(err)
	} else {
		if _, err := data.Exec(member.AccountId, member.Username,member.Mail, member.Name, member.Surname, member.Phone, member.Role, member.AccountId); err != nil {
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
		data, err := DB.Prepare("INSERT INTO members( username,password,mail,name,surname,phone,role) VALUES(?,?,?,?,?,?,?)")
		if err != nil {
			c.JSON(http.StatusBadRequest, "### Insert table failed ###")
		} else {
          
			hash, _ := HashPassword(member.Password)
			data.Exec(member.Username, hash, member.Mail, member.Name, member.Surname, member.Phone, member.Role)
			memberDetail.Username = member.Username
			memberDetail.Password = hash
			memberDetail.Mail = member.Mail
			memberDetail.Name = member.Name
			memberDetail.Surname = member.Surname
			memberDetail.Phone = member.Phone
			memberDetail.Role = member.Phone						
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
