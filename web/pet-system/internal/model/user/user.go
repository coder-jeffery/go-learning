package user

import  "time"

//数值类型、布尔类型、字符串类型、派生类型

type User struct{
	ID int
	Name string
	Address string
	Age int
	Weight float32
	Salary string
	State bool
	BirthTime time.Location  // // 输出：2026-01-01 00:00:00 +0800 CST
	BirthDate time.Time
}