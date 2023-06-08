package main

import (
	"fmt"
)

// 此樓往前走，有沒有比自己高? 沒有才記下自己位置。
func findBuildings(heights []int) []int {
	var result []int

	// result = append(result, 10)

	if len(heights) == 1{
		result = append(result, 0)
		return result
	}

	for i := 0 ; len(heights) - 1 >= i ; i++ {
		for a := i ; len(heights) - 1 >= a ; a++ {
			if heights[a] > heights[i] {
				break
			}
			
			if heights[a] < heights[i] && len(heights) - 1 == a{
				result = append(result, i)
			}
		} 
	}

    return result
}

func main(){
	var arr = []int{4,2,3,1}
	fmt.Println(findBuildings(arr))

	// var arr1 = []int{1}
	// fmt.Println(findBuildings(arr1))

	var arr2 = []int{1,3}
	fmt.Println(findBuildings(arr2))
}