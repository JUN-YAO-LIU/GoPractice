package main

import(
	"fmt"
)

func main(){
	nums := []int{-10,-3,-5,-1}

	fmt.Println(FindMaxOrMin(nums))
}


func FindMaxOrMin(nums []int) (int,int){
	max := nums[0]
	min := nums[0]

	for _,val := range nums{
		if val > max{
			max = val
		}
	}

	for _,val := range nums{
		if val < min{
			min = val
		}
	}

	return max,min
}