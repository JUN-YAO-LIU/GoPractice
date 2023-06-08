package main

import (
	"fmt"
)

func main() {
	// 第一種方式
	// 建立一個型別為 map 的變數，其中的 key 都會是 string，value 也都會是 string
	colors1 := map[string]string{}
	fmt.Println(colors1)
	// 也可以直接帶值
	  colors2 := map[string]string{
		  "red":   "#ff0000",
		  "green": "#4bf745",
	  }

	  fmt.Println(colors2)
  
	// 第二種方式
	var colors3 map[string]string
	fmt.Println(colors3)
  
	// 第三種方式，使用 make 建立 Map。
	// Key 的型別是 string，Value 是 int
	colors4 := make(map[string]int)
	colors4["red"] = 10
	  fmt.Println(colors4) // map[red:10]
  
	// 如果對於鍵值的數量有概念的話，也可以給定初始大小，將有助於效能提升（The little Go book）
	m := make(map[string]int, 100)
	fmt.Println(m)

	m2 := make([][]string,7)
	fmt.Println(m2)
  }