package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func main() {
	router := SetupRouter()
	router.Run()
}

func SetupRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"hello": "world",
		})
	})

	router.POST("/plus/:x/:y", func(c *gin.Context) {
		x, _ := strconv.ParseInt(c.Param("x"), 10, 64)
		y, _ := strconv.ParseInt(c.Param("y"), 10, 64)

		c.JSON(http.StatusOK, gin.H{
			"result": x + y,
		})
	})

	router.POST("/multiple/:x/:y", func(c *gin.Context) {
		x, _ := strconv.ParseInt(c.Param("x"), 10, 64)
		y, _ := strconv.ParseInt(c.Param("y"), 10, 64)

		c.String(http.StatusOK, fmt.Sprint(x*y))
	})

	router.GET("/profile", func(c *gin.Context) {
		profile := map[string]string{
			"name": "Surya",
			"role": "Developer",
		}

		c.JSON(http.StatusOK, gin.H{
			"status": 200,
			"data":   profile,
		})
	})

	return router
}
