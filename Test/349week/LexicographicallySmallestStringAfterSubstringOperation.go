package main

import (
	"fmt"
)

func smallestString(s string) string {
    i := 0
	for i < len(s) && s[i] == 'a' {
		i++
	}

	arr := []byte(s)
	if i == len(arr) {
		arr[len(s)-1] = 'z'
	} else {
		for i < len(s) && s[i] != 'a' {
			arr[i]--
			i++
		}
	}

	return string(arr)
}


func main(){
	s := "leetcode"
	//fmt.Println(string(s[1]))
	fmt.Println(smallestString(s))
}