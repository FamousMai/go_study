package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func middleOne() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("我在方法1前")
		c.Next()
		fmt.Println("我在方法1后")
	}
}

func middleTwo() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("我在方法2前")
		c.Next()
		fmt.Println("我在方法2后")
	}
}

func main() {
	r := gin.Default()
	/**
	.Use 就可以使用中间件了
	*/
	v1 := r.Group("v1").Use(middleOne(), middleTwo())
	v1.GET("testRouter", func(c *gin.Context) {
		fmt.Println("我在分组方法v1内部")
		c.JSON(200, gin.H{
			"success": true,
		})
	})

	r.Run(":1010")
}
