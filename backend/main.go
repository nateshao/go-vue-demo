package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main()  {
	//fmt.Println("hello world!")
	//r := gin.Default()
	//r.GET("/hello", func(c *gin.Context) {
	//	c.JSON(200,gin.H{
	//		"message":"你好",
	//	})
	//
	//})
	//panic(r.Run())

	r := gin.Default()
	r.GET("/api/auth/register", func(c *gin.Context) {
		// 获取参数
		name := c.PostForm("name")
		telephone := c.PostForm("telephone")
		password := c.PostForm("password")
		if len(telephone) != 11 {
			c.JSON(http.StatusUnprocessableEntity,gin.H{
				"code":422,
				"msg":"手机号必须为11位",
			})
		}
		if len(password) != 1 {
			
		}
		if len(name) != 1 {
			
		}
		// 判断手机号不存在

		// 
	})
}



















































