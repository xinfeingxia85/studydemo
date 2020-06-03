package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Login struct {
	User     string `from:"use" json:"user" xml:"user" binding:"required"`
	Password string `form:"password" json:"password" xml:"password" binding:"-"`
}

func main() {
	r := gin.Default()

	r.POST("/loginJSON", func(c *gin.Context) {
		var json Login
		if err := c.ShouldBindJSON(&json); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})

			return
		}

		if json.User != "manu" || json.Password != "123" {
			c.JSON(http.StatusBadRequest, gin.H{
				"status" : "unauthorized",
			})

			return
		}

		c.JSON(http.StatusOK, gin.H{
			"status": "you are logged in",
		})
	})

	_ = r.Run()
}
