package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Persons struct {
	ID   string `uri:"id" binding:"required,uuid"`
	Name string `uri:"name" binding:"required"`
}

type MyForm struct {
	Colors []string `form:"colors[]"`
}

var fromStr = `
<form action="/post" method="POST">
    <p>Check some colors</p>
    <label for="red">Red</label>
    <input type="checkbox" name="colors[]" value="red" id="red">
    <label for="green">Green</label>
    <input type="checkbox" name="colors[]" value="green" id="green">
    <label for="blue">Blue</label>
    <input type="checkbox" name="colors[]" value="blue" id="blue">
    <input type="submit">
</form>
`

func main() {
	r := gin.Default()
	r.GET("/form", func(c *gin.Context) {
		c.Header("content-type", "text/html;charset=uft-8")
		c.String(http.StatusOK, fromStr)
	})

	//r.GET("/:name/:id", func(c *gin.Context) {
	//	var person Persons
	//	if err := c.ShouldBindUri(&person); err != nil {
	//		c.JSON(400, gin.H{
	//			"msg": err.Error(),
	//		})
	//
	//		return
	//	}
	//
	//	c.JSON(http.StatusOK, gin.H{
	//		"name": person.Name,
	//		"id":   person.ID,
	//	})
	//})

	r.POST("/post", func(c *gin.Context) {
		var myForm MyForm
		if err := c.ShouldBind(&myForm);err !=nil{
			c.JSON(400, gin.H{
				"err": err.Error(),
			})
			return
		}

		c.JSON(200, gin.H{
			"color": myForm.Colors,
		})
	})

	r.Run()
}
