package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

type PostParams struct {
	Name string `json:"name" uri:"name" form:"name" binding:"required"`
	Age  int    `json:"age" uri:"age" form:"age" binding:"required,mustBig"`
	Sex  bool   `json:"sex" uri:"sex" form:"sex" binding:"required"`
}

func mustBig(fl validator.FieldLevel) bool {
	age := fl.Field().Interface().(int)
	fmt.Println(age)

	if age <= 18 {
		return false
	}
	return true
}

func main() {
	r := gin.Default()

	/**
	示例1：基础请求
	*/
	r.GET("/path/:id", func(c *gin.Context) {
		id := c.Param("id")
		user := c.DefaultQuery("user", "mym default")
		pwd := c.Query("pwd")
		c.JSON(200, gin.H{
			"id":   id,
			"user": user,
			"pwd":  pwd,
		})
	})

	r.POST("/path", func(c *gin.Context) {
		id := c.PostForm("id")
		user := c.DefaultPostForm("user", "mym post default")
		pwd := c.PostForm("pwd")
		c.JSON(200, gin.H{
			"id":   id,
			"user": user,
			"pwd":  pwd,
		})
	})

	r.DELETE("/path/:id", func(c *gin.Context) {
		id := c.Param("id")
		c.JSON(200, gin.H{
			"id": id,
		})
	})

	r.PUT("/path", func(c *gin.Context) {
		id := c.PostForm("id")
		user := c.DefaultPostForm("user", "mym put default")
		pwd := c.PostForm("pwd")
		c.JSON(200, gin.H{
			"id":   id,
			"user": user,
			"pwd":  pwd,
		})
	})

	/**
	示例2：bind模式：ShouldBindJSON
	*/
	r.POST("/testBind", func(c *gin.Context) {
		var p PostParams
		err := c.ShouldBindJSON(&p)
		if err != nil {
			c.JSON(200, gin.H{
				"msg":  "报错了",
				"data": gin.H{},
			})
		} else {
			c.JSON(200, gin.H{
				"msg":  "请求成功",
				"data": p,
			})
		}
	})

	/**
	示例3：bind模式：ShouldBindUri
	*/
	r.POST("/testBindUri/:name/:age/:sex", func(c *gin.Context) {
		var p PostParams
		err := c.ShouldBindUri(&p)
		if err != nil {
			c.JSON(200, gin.H{
				"msg":  "报错了",
				"data": gin.H{},
			})
		} else {
			c.JSON(200, gin.H{
				"msg":  "请求成功",
				"data": p,
			})
		}
	})

	/**
	示例4：bind模式：ShouldBindQuery
	*/
	r.POST("/testBindQuery", func(c *gin.Context) {

		if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
			v.RegisterValidation("mustBig", mustBig)
		}

		var p PostParams
		err := c.ShouldBindQuery(&p)
		fmt.Println(err)
		if err != nil {
			c.JSON(200, gin.H{
				"msg":  "报错了",
				"data": gin.H{},
			})
			fmt.Println(err.Error())
		} else {
			c.JSON(200, gin.H{
				"msg":  "请求成功",
				"data": p,
			})
		}
	})

	r.Run(":1010") // listen and serve on 0.0.0.0:8080

}
