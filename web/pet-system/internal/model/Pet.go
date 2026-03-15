package internal

import (
	"fmt"
	"time"
)

type Pet struct{
	ID int
	Name string
	Weight float32
	Color string
	Birth time.Location
	Dead bool
}


func (p Pet) Getinfo() string{
	//return fmt.Sprintln("id: %d, name: %s, weight: %s, color: %s, brith: %s, dead: %s,", p.ID, p.Name, p.Weight, p.Color, p.Birth, p.Dead)
	return fmt.Sprintln("id: %D, name: %s ", p.ID, p.Name)
}

func (p *Pet) Update(newname string) error{
	if newname == ""{
		return fmt.Errorf("name cannot be empty")
	}
	p.Name = newname
	return  nil
}
//func main() {
//	p := Pet{ID: 1, Name: "jim", Weight: 11.1, Color: "red", Birth: time.Location{}, Dead: false}
//	fmt.Println(p)
//}