package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Class struct {
	gorm.Model
	ClassName string
	Students  []Student
}

type Student struct {
	gorm.Model
	StudentName string
	ClassID     uint
	IDCard      IDCard
	// 多对多
	Teachers []Teacher `gorm:"many2many:student_teachers;"`
}

type IDCard struct {
	gorm.Model
	StudentID uint
	Num       int
}

type Teacher struct {
	gorm.Model
	TeacherName string
	Students    []Student `gorm:"many2many:student_teachers;"`
}

func main() {
	dsn := "root:mym123456@tcp(127.0.0.1:3306)/ginclass?charset=utf8mb4&parseTime=True&loc=Local"
	db, _ := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	fmt.Println(db)

	db.AutoMigrate(&Teacher{}, &Student{}, &Class{}, &IDCard{})

	r := gin.Default()

	/**
	新建学生
	*/
	r.POST("/student", func(context *gin.Context) {
		var student Student
		_ = context.BindJSON(&student)
		db.Create(&student)
	})

	/**
	查询
	*/
	r.GET("/student/:ID", func(context *gin.Context) {
		id := context.Param("ID")
		var student Student
		_ = context.BindJSON(&student)
		db.Preload("Teachers").Preload("IDCard").Where("id = ?", id).First(&student)

		context.JSON(200, gin.H{
			"s": student,
		})

	})

	r.GET("/class/:ID", func(context *gin.Context) {
		id := context.Param("ID")
		var class Class
		db.Preload("Students").Preload("Students.Teachers").Preload("Students.IDCard").Where("id = ?", id).First(&class)

		context.JSON(200, gin.H{
			"c": class,
		})
	})

	r.Run(":1010")

}
