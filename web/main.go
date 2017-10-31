package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ufoscout/go_microservices/client"
)

// This function's name is a must. App Engine uses it to drive the requests properly.
func main() {
	// Starts a new Gin instance with no middle-ware
	r := gin.New()

	r.GET("/hello", func(c *gin.Context) {
		name := c.DefaultQuery("name", "Francesco")
		client.Hello(name)
		c.String(http.StatusOK, "")
	})

	//r.LoadHTMLGlob("templates/**/*.html")
	r.Static("/vendor", "./vendor")
	r.StaticFile("/", "./templates/index.html")

	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")
}
