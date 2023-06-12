package main

import (
	"fmt"
	"sort"
)

func findNonMinOrMax(nums []int) int {
	r := -1

	if len(nums) > 2{
		sort.Slice(nums,func (i,j int) bool { return nums[i] > nums[j]})
		return nums[1]
	}

    return r
}


func main(){
	// arr1 := []int{3,1,5,7}
	arr1 := []int{3,1,4}
	fmt.Println(findNonMinOrMax(arr1))
}