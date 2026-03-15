package main

import (
	"encoding/json"
	"fmt"
)

//结构体字段可见性和JSON序列化
//如果一个GO语言定义标识符是首字符大写 那么就对外可见
//JSON: javascript object notation

type Cat struct {
	ID   int
	Name string
}

func newCat(id int, name string) Cat {
	return Cat{
		ID:   id,
		Name: name,
	}
}

type class struct {
	Title string `json:"title"`
	Cat   []Cat  `json:"cat_list" db:"cat_list" xml:"cat"`
}

func main() {

	c1 := class{
		Title: "虎斑英短猫",
		Cat:   make([]Cat, 0, 20),
	}

	for i := 0; i < 5; i++ {
		temStu := newCat(i, fmt.Sprintf("%d cat: ", i))
		c1.Cat = append(c1.Cat, temStu)
	}
	fmt.Printf("%#v\n ", c1)

	//JSON 序列化 Go语言数据转换为JSON格式字符串
	data, err := json.Marshal(c1)
	if err != nil {
		fmt.Println("json marshal failed err", err)
		return
	}
	fmt.Println("************************************************************")
	fmt.Printf("%s\n", data)
	fmt.Printf("%T\n", data)

	//JSON反序列化
	jsonStr := `{"Title":"虎斑英短猫","Cat":[{"ID":0,"Name":"0 cat: "},{"ID":1,"Name":"1 cat: "},{"ID":2,"Name":"2 cat: "},{"ID":3,"Name":"3 cat: "},{"ID":4,"Name":"4 cat: "}]}`
	var c2 class
	json.Unmarshal([]byte(jsonStr), &c2)
	if err != nil {
		fmt.Println("json unmarshal failed , err ", err)
		return
	}
	fmt.Printf("%#v\n", c2)

}
