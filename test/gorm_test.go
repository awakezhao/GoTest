package test

import (
	"fmt"
	"testing"

	_ "github.com/go-sql-driver/mysql"
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

func TestGorm(t *testing.T) {
	dsn := "root:123456@(localhost)/account?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
	}

	user := User{Name: "Jinzhu", Age: 18}

	result := db.Create(&user) // 通过数据的指针来创建
	if result.Error != nil {
		return
	} else {
		fmt.Println(result.RowsAffected)
	}
}
