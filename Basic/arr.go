package main

import (
	"fmt"
)

// func generate(numRows int) [][]int {
//     triangle := make([][]int, 0)

// 	for numRow := 0; numRow < numRows; numRow++ {
// 		layer := make([]int, 0)
// 		for numElement := 0; numElement <= numRow; numElement++ {
// 			value := 1
// 			if numRow > 1 && numElement > 0 && numElement < numRow {
// 				value = triangle[numRow-1][numElement-1] + triangle[numRow-1][numElement]
// 			}
// 			layer = append(layer, value)
// 		}
// 		triangle = append(triangle, [][]int{layer}...)
// 	}

// 	return triangle
// }

func main(){
	// 如何新增值到多維陣列
	t := make([][]int, 0) // make需要給大小
	fmt.Println("init :",t)

	// t[0][0] = {1,3} 錯的
	// t[0] = {1,3} 錯的
	layer := []int{1,3,4}
	layer2 := []int{1,2,4}

	// array 相加就會需要加上...，表示
	t = append(t, [][]int{layer}...)
	t = append(t, [][]int{layer2}...)
	fmt.Println("賦值 :",t)

	// 也可以這樣
	fmt.Println("arr :",[][]int{layer2,layer})
}