package main

import (
	"fmt"
)

type InitCar struct{
	name string
	price int
}

func Contructor(n string,p int)InitCar{
	return InitCar{name:n,price:p}
}

func (this *InitCar) Get() (int,string){
	return this.price,this.name
}

func main(){
	var newCar = InitCar{name:"Jim",price:123}
	fmt.Println(newCar.Get())
}