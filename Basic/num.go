package main

import(
	"fmt"
)

func main(){
	nums := []int{-10,-3,-5,-1}

	fmt.Println(FindMaxOrMin(nums))

	// big int。num1是10進至，把字串轉成數字
	bigIntNum1, _ := new(big.Int).SetString(num1, 10)
	bigIntNum2, _ := new(big.Int).SetString(num2, 10)

	sum := new(big.Int).Add(bigIntNum1, bigIntNum2)
	// 轉成字串
	fmt.Println(sum.String())
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