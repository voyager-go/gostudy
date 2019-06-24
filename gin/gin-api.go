package main

import (
	"github.com/gin-gonic/gin"
)

func GetHomePage(c *gin.Context)  {
	c.JSON(200, "Hello gin")
}
func PostHomePage(c *gin.Context)  {
	c.JSON(200,  gin.H{
		"message": "Post HomePage",
	})
}
func QueryString(c *gin.Context)  {
	id   := c.Query("id")
	name := c.Query("name")
	c.JSON(200, gin.H{
		"id":   id,
		"name": name,
	})
}
func PathParameters(c *gin.Context)  {
	name := c.Param("name")
	age  := c.Param("age")
	c.JSON(200, gin.H{
		"name": name,
		"age":  age,
	})
}
func main() {
	r := gin.Default()
	r.GET("/", GetHomePage)
	r.GET("/post", PostHomePage)
	r.GET("/query", QueryString) // /query?id=2&name=zwj
	r.GET("/path/:name/:age", PathParameters) // /path/zwj/23
	r.Run(":8989")
}
