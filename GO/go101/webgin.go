package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var Product struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       int    `json:"price"`
}

func main() {
	router := gin.Default()

	router.GET("/user/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.String(http.StatusOK, "Hello %s", name)
	})
	router.GET("/user/:name/*action", func(c *gin.Context) {
		name := c.Param("name")
		action := c.Param("action")
		message := name + " is " + action
		c.String(http.StatusOK, message)
	})

	router.Run(":8080")
}
