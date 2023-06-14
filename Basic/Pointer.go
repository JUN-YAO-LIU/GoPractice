package main

import(
	"fmt"
)


func main(){

	// 定義這個變數可以儲存，記憶體位置
	var p *int
	i := 10

	// 一般變數取出記憶體位置
	p = &i

	fmt.Println(p)

	// 指向記憶體位置的值
	fmt.Println(*p)

	// 指向位置的值，做++
	*p++
	fmt.Println(i)
}