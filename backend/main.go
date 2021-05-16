package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)


func main()  {
	fmt.Println("hello world!")
	r := gin.Default()
	r.GET("/hello", func(c *gin.Context) {
		c.JSON(200,gin.H{
			"message":"你好",
		})

	})
	panic(r.Run())
}
