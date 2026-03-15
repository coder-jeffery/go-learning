package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
)
import _ "github.com/jinzhu/gorm/dialects/mysql"

type Student struct {
	Name string
	Age  int
}

func main() {
	//db, err := gorm.Open("mysql", "root:root@tcp(127.0.0.1:3306)/trefsys?charset=utf8&parseTime=True")
	//db, err := gorm.Open("mysql", "root:root@tcp(127.0.0.1:3306)/trefsys?charset=utf8&parseTime=True&loc=Asia%2FShanghai")
	db, err := gorm.Open("mysql", "root:root@tcp(127.0.0.1:3306)/trefsys?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	//https://github.com/go-gorm
	//go get github.com/jinzhu/gorm
	//创建表结构
	db.CreateTable(&Student{})
	//db.Table("student").Create(&Student{})

	//删除表结构
	db.DropTable(&Student{})
	//db.DropTableIfExists(&Student{})
	//db.DropTable("student")

	//判断表是否存在
	flag := db.HasTable(&Student{})
	fmt.Println(flag)

	fmt.Printf("\n execute success......... %t\n", flag)
}
