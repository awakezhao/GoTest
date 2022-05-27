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

func (u *User) String() string {
	return fmt.Sprintf("Name:%s\t Age:%d", u.Name, u.Age)
}

func TestGorm(t *testing.T) {
	dsn := "root:123456@(localhost)/account?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
	}

	user := User{Name: "Jinzhu2", Age: 18}

	var A User

	err = db.Where("name = ?", user.Name).First(&A).Error
	if err == nil { //已有
		fmt.Println(A.Name, A.Age)
	} else { // 没有
		result := db.Create(&user) // 通过数据的指针来创建
		if result.Error != nil {
			return
		} else {
			fmt.Println(result.RowsAffected)
		}
	}
}

// err = dbGlobal.Table(tableName).Where("nickname = ?", req.Nickname).First(&nickData).Error
// if err == nil {
// 	if nickData.PlayerID > 0 {
// 		msg.Head.Result = protocol.ErrorRoleNameExisted
// 		accSrv.SendMsgWithID(msg)
// 		return
// 	}
// }
