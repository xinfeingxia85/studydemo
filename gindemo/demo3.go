package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()

	r.GET("/morejson", func(c *gin.Context) {
		var msg struct {
			Name    string `json:"user"`
			Message string `json:"message"`
			Number  int    `json:"number"`
		}

		msg.Name = "fangkm"
		msg.Message = "ceshi"
		msg.Number = 10

		c.JSON(http.StatusOK, msg)
	})

	r.GET("/somexml", func(c *gin.Context) {
		c.XML(http.StatusOK, gin.H{
			"message": "fangkm",
			"age":     20,
		})
	})

	r.GET("/securejson", func(c *gin.Context) {
		name :=[]string{"lena", "austin", "foo"}
		c.SecureJSON(http.StatusOK, name)
	})
	
	r.GET("/jsonp", func(c *gin.Context) {
		data := struct {
			Name string
			Sex string
		}{"fangkm", "男"}

		c.JSONP(http.StatusOK, data)
	})

	r.Run()
}
