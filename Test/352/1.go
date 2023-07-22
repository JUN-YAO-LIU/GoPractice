package main

import (
	"fmt"
	//"sort"
)

func longestAlternatingSubarray(nums []int, threshold int) int {
    
    // l第一個數 一定是偶數
    lindex := 0
    rindex := 0
    max := 0
    lenght := len(nums) -1
    var arr []int
    
    for i:=0 ;lenght >= i;i++{
        
        if nums[i] % 2 ==0{
            lindex = i
			break
        }
        
     
    }

	for i:=0 ;lenght >= i;i++{
        if nums[i] <= threshold && nums[i] > max{
            max = nums[i]
            rindex = i
        }
    }
    
	fmt.Println("r",rindex)
	fmt.Println("l",lindex)
    if lindex == 0 || rindex == 0{
        return 0
    }
    

    // index [l ,r-1] 
    // [l,r]
    for i:=lindex ;rindex >= i;i++{
		
        
        if i + 1 < lenght && lindex <= rindex - 1 && nums[i] % 2 != nums[i + 1] % 2{
           arr = append(arr,nums[i])
        }else if lindex <= rindex && nums[i] <= threshold{
           arr = append(arr,nums[i])
        }
    }
    fmt.Println(arr)
    return len(arr)
}


func main(){
	// arr1 := []int{3,1,5,7} 5
	// [1,2]
	// 2
	// [2,3,4,5]
	// 4

	arr1 := []int{3,2,5,4}
	// arr1 := []int{1,2}
	fmt.Println(longestAlternatingSubarray(arr1,5))
}