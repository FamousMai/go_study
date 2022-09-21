package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type HelloWord struct {
	Name string
	Sex  bool
	Age  int
	gorm.Model
}

func main() {
	// 参考 https://github.com/go-sql-driver/mysql#dsn-data-source-name 获取详情
	dsn := "root:mym123456@tcp(127.0.0.1:3306)/ginclass?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println(db)
		fmt.Println(err.Error())
	}

	// 迁移/创建表
	error2 := db.AutoMigrate(&HelloWord{})
	if error2 != nil {
		fmt.Println(error2.Error())
	}

	// 新增数据
	//db.Create(&HelloWord{
	//	Name: "grant",
	//	Age: 81,
	//	Sex: false,
	//})

	/**
	查询 单条数据
	*/
	var hello HelloWord

	db.First(&hello)
	fmt.Println(hello)

	db.First(&hello, "name = ?", "mym")
	fmt.Println(hello)

	/**
	查询 查多条数据
	*/
	var hello2 []HelloWord
	db.Find(&hello2)
	fmt.Println(hello2)

	db.Find(&hello2, "age < ?", 21)
	fmt.Println(hello2)

	db.Where("age > ?", 21).Find(&hello2)
	fmt.Println(hello2)

	db.Where("age > ?", 21).Or("").Find(&hello2)
	fmt.Println(hello2)

	/**
	修改 单个字段
	*/
	db.Where("id = ?", 2).First(&HelloWord{}).Update("name", "mym3333")

	/**
	修改 单个用户 多个字段
	*/
	db.Where("id = ?", 1).First(&HelloWord{}).Updates(HelloWord{
		Name: "mym shuai",
		Age:  28,
	})

	/**
	修改 单个用户 多个字段 map
	*/
	db.Where("id = ?", 3).First(&HelloWord{}).Updates(map[string]interface{}{
		"Name": "mym map",
		"Sex":  false,
		"Age":  28,
	})

	/**
	修改 多个用户 多个字段 map
	*/
	db.Where("id in (?)", []int{1, 2}).Find(&HelloWord{}).Updates(map[string]interface{}{
		"Name": "mym piliang",
		"Sex":  true,
		"Age":  28,
	})

	/**
	删除 软删除
	*/
	db.Where("id in (?)", []int{1, 2}).Delete(&HelloWord{})

	/**
	删除 硬删除
	*/
	db.Where("id in (?)", []int{1, 2}).Unscoped().Delete(&HelloWord{})
}
