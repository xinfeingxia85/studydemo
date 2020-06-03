package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"os"
	"time"
)

type Person struct {
	Name    string `form:"name"`
	Address string `form:"address"`
}

func main() {
	//禁止控制台颜色
	gin.DisableConsoleColor()
	gin.SetMode(gin.ReleaseMode)

	//创建日志文件
	f, _ := os.Create("log/gin.log")
	gin.DefaultWriter = io.MultiWriter(f)

	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
			param.ClientIP,
			param.TimeStamp.Format(time.RFC1123),
			param.Method,
			param.Path,
			param.Request.Proto,
			param.StatusCode,
			param.Latency,
			param.Request.UserAgent(),
			param.ErrorMessage,
		)
	}))

	//router
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
			"code":    200,
		})
	})

	r.GET("/user/:id", func(c *gin.Context) {
		log.Println(c.GetQuery("keyword"))
		log.Println(c.Query("keyword"))
		fmt.Println(c.Params.Get("id"))
	})

	r.GET("/look", startPage)
	r.POST("/look", startPage)
	r.GET("fangkm/:id/*action", func(c *gin.Context) {
		log.Println(c.Param("id"))
		log.Println(c.Param("action"))
	})

	v1 := r.Group("/v1")
	{
		v1.GET("/user/*action", func(c *gin.Context) {
			c.String(200, "URL: %s", c.Request.URL)
		})

		v1.GET("/group/:id", func(c *gin.Context) {
			c.String(200, "URL: %s", c.Request.URL)
		})
	}

	// Simple group: v1
	v2 := r.Group("/v2")
	{
		v2.GET("/login", func(c *gin.Context) {
			c.String(200, "%s", c.Request.URL)
		})
		v2.GET("/submit", func(c *gin.Context) {
			c.String(200, "%s", c.Request.URL)

		})
		v2.GET("/read", func(c *gin.Context) {
			c.String(200, "%s", c.Request.URL)
		})
	}

	_ = r.Run() // listen and serve on 0.0.0.0:8080
}

func startPage(c *gin.Context) {
	var person Person

	if c.ShouldBind(&person) == nil {
		log.Println(person.Name)
		log.Println(person.Address)
	}
	c.String(200, "sucess")
}
