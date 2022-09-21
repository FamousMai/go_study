package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"os"
)

func main() {
	router := gin.Default()
	// 给表单限制上传大小 (默认 32 MiB)
	// router.MaxMultipartMemory = 8 << 20  // 8 MiB
	router.POST("/testUpload", func(c *gin.Context) {
		// 单文件
		file, _ := c.FormFile("file")
		name := c.PostForm("name")

		/**
		框架方法
		*/
		c.SaveUploadedFile(file, "./upload/"+file.Filename)

		/**
		自己实现
		*/
		in, _ := file.Open()
		defer in.Close()
		out, _ := os.Create("./upload/" + file.Filename)
		defer out.Close()
		io.Copy(out, in)

		c.JSON(200, gin.H{
			"msg":  file,
			"name": name,
		})

		/**
		TODO file 直接返回文件，但是用apipost，暂未实现
		*/
		c.Writer.Header().Add("Content-Disposition", fmt.Sprintf("attachment; filename=%s", file.Filename))
		c.File("./" + file.Filename)

		/**
		多文件
		*/
		form, _ := c.MultipartForm()
		files := form.File["file"]
		fmt.Println(files)
		for _, f := range files {
			log.Println(f.Filename)
		}
	})
	router.Run(":1010")
}
