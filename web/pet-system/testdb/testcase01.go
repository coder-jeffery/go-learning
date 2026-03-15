package main

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// 定义模型
type User struct {
	gorm.Model
	Name  string
	Email string `gorm:"uniqueIndex"`
	Age   int
}

func main() {
	// 连接字符串：user=用户名 password=密码 dbname=数据库名 host=地址 port=端口 sslmode=disable
	dsn := "host=localhost user=myuser password=mypassword dbname=mydb port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		//panic("连接数据库失败: " + err.Error())
	}

	// 自动迁移表（创建/更新表结构）
	db.AutoMigrate(&User{})

	// 插入数据
	user := User{Name: "张三", Email: "zhangsan@example.com", Age: 20}
	db.Create(&user)
	fmt.Printf("插入用户ID: %d\n", user.ID)

	// 查询数据
	var queryUser User
	db.First(&queryUser, 1) // 根据ID=1查询
	fmt.Printf("查询到用户: %s, %s\n", queryUser.Name, queryUser.Email)
}