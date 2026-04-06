package model

import "fmt"

type Flyable interface{
	Fly()
}


type Bird struct {
	Name string
}

func (b Bird) Fly(){
	fmt.Sprintln("s%", b.Name)
}



type Plane struct{
	Model string
}

func (p Plane) Fly(){
	fmt.Sprintf("s%", p.Model)
}

func LetFly(f Flyable){
	f.Fly()
}

//func main() {
//	b1 := Bird{Name: "xiao bird"}
//	p1 := Plane{Model : "波音131"}
//
//	LetFly(b1)
//	LetFly(p1)
//}
