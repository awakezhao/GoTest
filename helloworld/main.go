// package必须是文件的代码的第一行，不包括注释

package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	Name string
	Age  int32
}

func (u *User) TableName() string {
	return "user"
}

func main() {
	fmt.Println("===============hello, world==================")
	Run()
}

func Run() {
	// 参考 https://github.com/go-sql-driver/mysql#dsn-data-source-name 获取详情
	dsn := "root:123456@tcp(127.0.0.1:3306)/account?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(db.Config)

	user := User{
		Name: "zhz01",
		Age:  18,
	}
	db.Create(&user)
}
