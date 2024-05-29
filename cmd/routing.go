package main

import (
	ports "HEXAGONAL-GO/cmd/services/dashboard-api/ports/drivers"

	"github.com/gin-gonic/gin"
)

func CreateRouter(userAdapter ports.ForUser) *gin.Engine {
	router := gin.Default()

	// Create GetUser API
	router.GET("/getUser", func(c *gin.Context) {
		username := c.Query("username")
		user, err := userAdapter.GetUser(username)
		if err != nil {
			c.JSON(500, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(200, user)
	})

	return router
}
