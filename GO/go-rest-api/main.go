package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {

	fmt.Println("Hello World")

	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://Core2931:1150@cluster0.6zoeef8.mongodb.net/test"))
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	err = client.Connect(ctx)
	db := client.Database("test")
	coll := db.Collection("customersusers")

	// var result struct {
	// 	Name     string `bson:"name"`
	// 	Email    string `bson:"email"`E
	// 	Userid   string `bson:"userID"`
	// 	Platform string `bson:"platform"`
	// }
	var result struct {
		Username  string `bson:"username"`
		Email     string `bson:"email"`
		Firstname string `bson:"firstname"`
		Lastname  string `bson:"lastname"`
		Platform  string `bson:"platform"`
	}
	filter := bson.M{}
	ctx, _ = context.WithTimeout(context.Background(), 5*time.Second)
	err = coll.FindOne(ctx, filter).Decode(&result)

	// cursor, err := coll.Find(context.TODO(), filter)
	// err = cursor.Decode(&result)
	if err != nil {
		log.Fatal(err)
	}
	// fmt.Printf("Platform : %s \nUserID : %s \nName : %s \nEmail : %s \n", result.Platform, result.Userid, result.Name, result.Email)

	r := gin.New()
	// // r.GET("/", getListproducts())
	// // r.POST("/products", createProduct())
	// // r.DELETE("/products/:id", deleteProduct())
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, result)
	})

	// r.POST("/products", func(c *gin.Context) {
	// 	var product Product

	// 	if err := c.ShouldBindJSON(&product); err != nil {
	// 		c.JSON(http.StatusBadRequest, gin.H{
	// 			"error": err.Error(),
	// 		})
	// 		return
	// 	}
	// 	products = append(products, product)
	// 	c.JSON(http.StatusCreated, product)
	// })
	// r.DELETE("/products/:id", func(c *gin.Context) {
	// 	id := c.Param("id")

	// 	for i, a := range products {
	// 		if a.ID == id {
	// 			products = append(products[:i], products[i+1:]...)
	// 			break
	// 		}
	// 	}
	// 	c.Status(http.StatusNoContent)

	// })
	r.Run()

}

type Product struct {
	ID     string `json:"id"`
	Title  string `json:"Title"`
	Author string `json:"author"`
}

var products = []Product{
	{ID: "1", Title: "Bitcoind", Author: "Satoshi"},
	{ID: "2", Title: "Etherium", Author: "Vitalik"},
}

// func getListproducts(c *gin.Context) {
// 	c.JSON(http.StatusOK, products)
// }

// func createProduct(c *gin.Context) {
// 	var product Product

// 	if err := c.ShouldBindJSON(&product); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"error": err.Error(),
// 		})
// 		return
// 	}
// 	products = append(products, product)

// }

// func deleteProduct(c *gin.Context) {
// 	id := c.Param("id")

// 	for i, a := range products {
// 		if a.ID == id {
// 			products = append(products[:i], products[i+1:]...)
// 			break
// 		}
// 	}

// 	c.Status(http.StatusNoContent)

// }
