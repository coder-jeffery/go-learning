package main

import "fmt"

//nested structure

type Address struct {
	Province string
	City     string
}

type Driver struct {
	Name    string
	Age     int
	Gender  string
	Address Address
}

func main() {

	d1 := Driver{
		Name:    "John",
		Age:     18,
		Gender:  "male",
		Address: Address{Province: "Shanghai", City: "Shanghai"},
	}
	fmt.Printf("d1: %#v\n", d1)
	fmt.Println(d1.Name, d1.Age, d1.Gender, d1.Address)
}
