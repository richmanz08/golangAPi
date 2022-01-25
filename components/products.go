package components

import (
	"database/sql"
	"fmt"
	"net/http"

	MES "api-webapp/Message"

	"github.com/gin-gonic/gin"
)

type ProductStruct struct {
	// ProductID   int32   `json:"product_id"`
	ProductName string  `json:"product_name"`
	Description string  `json:"description"`
	ProductType string  `json:"product_type"`
	BgColor     string  `json:"bgColor"`
	Price       float64 `json:"price"`
}
type AllProductStruct struct {
	ProductID   int32   `json:"product_id"`
	ProductName string  `json:"product_name"`
	Description string  `json:"description"`
	ProductType string  `json:"product_type"`
	BgColor     string  `json:"bgColor"`
	Price       float64 `json:"price"`
}

var DB *sql.DB

func ShowAllProduct(c *gin.Context) {
	var allProduct []AllProductStruct
	data, err := DB.Query("SELECT * FROM products")
	if err != nil {
		fmt.Println(err)
	}
	defer data.Close()

	// fmt.Println(allProduct)
	for data.Next() {
		var newdata AllProductStruct
		err = data.Scan(&newdata.ProductID, &newdata.ProductName, &newdata.Description, &newdata.ProductType, &newdata.BgColor, &newdata.Price)
		if err != nil {
			panic(err.Error())
		}
		allProduct = append(allProduct, AllProductStruct{ProductID: newdata.ProductID, ProductName: newdata.ProductName, Description: newdata.Description, ProductType: newdata.ProductType, BgColor: newdata.BgColor, Price: newdata.Price})

		// show renponse data from table
		// log.Printf(data.Firstame, data.Id, data.Lastname, data.Role)

	}
	c.JSON(http.StatusOK, allProduct)
}

func AddProDuct(c *gin.Context) {
	// DATABASE AUTO_INCREMENT product_id
	var Product ProductStruct
	if err := c.ShouldBindJSON(&Product); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(Product)
	data, err := DB.Prepare("INSERT INTO products( product_name,description,product_type,bgColor,price) VALUES(?,?,?,?,?)")
	if err != nil {
		fmt.Println(err)
	}
	defer data.Close()

	data.Exec(Product.ProductName, Product.Description, Product.ProductType, Product.BgColor, Product.Price)

	c.JSON(http.StatusCreated, Product)
}

func UpdateProduct(c *gin.Context) {
	var Update AllProductStruct
	if err := c.ShouldBindJSON(&Update); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(Update)
	data, err := DB.Prepare("UPDATE products SET product_id=?,product_name=?, description=?,product_type=?,bgColor=?,price=? WHERE product_id=?")
	if err != nil {
		fmt.Println(err)
	}
	defer data.Close()	
	if _, err := data.Exec(Update.ProductID, Update.ProductName, Update.Description, Update.ProductType, Update.BgColor, Update.Price, Update.ProductID); err != nil {
		fmt.Println("smt.Exec failed: ", err)
	}
	c.JSON(http.StatusOK, MES.Update_Message)
}

func DeleteProduct(c *gin.Context) {
	// itemid := c.Param("product_id")
	// var req getAccountRequest
	itemid := c.Param("id")

	data, err := DB.Prepare("DELETE FROM products WHERE product_id=?")
	if err != nil {
		fmt.Println(err)
		return
	}

	defer data.Close()
	if _, err := data.Exec(itemid); err != nil {
		fmt.Println("smt.Exec failed: ", err)
	}
	c.JSON(http.StatusOK, MES.Deleted_Success)
}
