package main

import (
	"fmt"
	"sort"
)

func groupAnagrams(strs []string) [][]string {
    mapS := make(map[string][]string, 0)
    
    for _, val := range strs {
        byteString := []byte(val)
        sort.Slice(byteString, func(i,j int) bool { return byteString[i] < byteString[j]})
        sortedString := string(byteString)
        
        mapS[sortedString] = append(mapS[sortedString], val)
    }
    
    ans := make([][]string,0)

    for _,val := range mapS {
        ans = append(ans, val)
    }
    return ans
}

func test(strs []string) []string {
    // Logic to group anagrams
	return strs
}


func main(){
	// myStrings := []string{"cat", "tac", "dog", "god", "act"}
	// fmt.Println(test(myStrings))

	// 如果有4會出現錯誤
	// arr1 := [4]string{"tan", "nat", "bac", "abc"}
	arr1 := []string{"tan", "nat", "bac", "abc"}
	fmt.Println(groupAnagrams(arr1))
}



