package main

import (
	"fmt"
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

	db.AutoMigrate(&Teacher{}, &Class{}, &Student{}, &IDCard{})

	i := IDCard{
		Num: 123456,
	}

	s := Student{
		StudentName: "mym",
		IDCard:      i,
	}

	t := Teacher{
		TeacherName: "老师傅",
		Students:    []Student{s}, // TODO 这里加了之后会添加失败
	}

	c := Class{
		ClassName: "老麦的班级",
		Students:  []Student{s},
	}

	_ = db.Create(&c).Error
	_ = db.Create(&t).Error
}
