package main

import (
	"fmt"
)

// 此樓往前走，有沒有比自己高? 沒有才記下自己位置。
func findBuildings(heights []int) []int {
	var result []int

	for i := 0 ; len(heights) - 1 >= i ; i++ {
		add := true
		for a := i ; len(heights) - 1 >= a ; a++ {
			if heights[a] >= heights[i] && a != i {
				add = false
				break
			}
		} 

		if add {
			result = append(result, i)
		}
	}

    return result
}

func main(){
	var arr = []int{4,3,2,1}
	fmt.Println(findBuildings(arr))

	// var arr1 = []int{1}
	// fmt.Println(findBuildings(arr1))

	var arr2 = []int{2,2,2,2}
	fmt.Println(findBuildings(arr2))
}