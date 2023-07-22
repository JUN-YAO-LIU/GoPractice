package main

import (
	"fmt"
	"sort"
	"strings"
)

func main(){
	s := "abc"
	s1 := "acb"
	strArr := []string{"abc","cba"}

	// 可以用byte比大小，因為可以直接把字串轉成數字陣列。
	a := []byte(s1)
	fmt.Println([]byte(s))
	fmt.Println(a)

	// 幫忙排序的用法。
	sort.Slice(a,func(i, j int) bool {return a[i] > a[j]})
	sort.Slice(strArr,func(i, j int) bool {return a[i] > a[j]})

	// fmt.Println(a)
	// fmt.Println(strArr)

	fmt.Println(string(a))


	// 換成取一位會變成uint8
	fmt.Println("秀出string，單字")
	fmt.Println(s[0])
	
	// for _, v := range a {
    // 	fmt.Println(v)
	// }

	fmt.Println("-----Cut-----")
	sCut := "abaasdf123c"
	sCut = sCut[1:len(sCut)-1]
	fmt.Println(sCut)

	// 包含strings 包含可以用
	fmt.Println("包含：",strings.ContainsAny(sCut,"0"))

	// 兩個變數，交換值
	swap1 := "swap1"
	swap2 := "swap2"

	swap1 , swap2 = swap2, swap1
	fmt.Println("swap1 :",swap1,"swap2 :",swap2)
}