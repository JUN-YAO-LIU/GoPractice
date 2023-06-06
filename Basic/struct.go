package main

import(
	"fmt"
)


func main(){

	// type 和 struct都是小寫
	type Person struct{
		name string
		age int
	}

	// variable
	var str string
	var p *string
	str = "123"

	var person1 Person

	person1.name = "Jim"
	person1.age = 18

	fmt.Println(str)
	fmt.Println(person1)

	p = &str
	fmt.Println(p)
	fmt.Println(*p)
}