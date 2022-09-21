package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name string `gorm:"primary_key; column:user_name; type:varchar(100);"`
}

func (u User) TableName() string {
	/**
	可以用这个功能进行分表
	*/
	if u.Name == "aaa" {
		return "qm_users_aaa"
	}
	return "qm_users"
}

func main() {

	dsn := "root:mym123456@tcp(127.0.0.1:3306)/ginclass?charset=utf8mb4&parseTime=True&loc=Local"
	db, _ := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	error2 := db.AutoMigrate(&User{})
	if error2 != nil {
		fmt.Println(error2.Error())
	}
}
