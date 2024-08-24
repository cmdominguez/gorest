package main

import (
	"gorest/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func setupRoutes(r *gin.Engine) {
	utils.LogAlgo("Algo")
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": 5,
		})
	})
}

func main() {
	r := gin.Default()
	setupRoutes(r)
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
