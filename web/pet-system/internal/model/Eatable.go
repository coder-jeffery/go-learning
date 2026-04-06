package model

import "fmt"

type Eatable interface{
	Eat()
}

type Dog struct {
	Name string
}

func (d Dog) Eat(){
	fmt.Printf("%s", d.Name)
}

type Cat struct {
	Description string
}

func (c Cat) Eat(){
	fmt.Printf("s%", c.Description)
}


func CommonPet(eat Eatable){
	eat.Eat()
}


