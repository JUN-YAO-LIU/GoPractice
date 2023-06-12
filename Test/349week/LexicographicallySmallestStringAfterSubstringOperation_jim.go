package main

import (
	"fmt"
)

// 問題 如何字串 轉 char 、如何取得array value or key、spilt、join用法


func getIndex(s string) int{
	h := []string{"a","b","c","d","e","f","g","h","i","j","k","l","m","n","o","p","q","r","s","t","u","v","w","x","y","z"}
	var r  = -1

	// array 本來就是0開始
	for key,val := range h{
		if s == val{
			r = key
		}
	}
	return r
}
    
func smallestString(s string) string {
	h := []string{"a","b","c","d","e","f","g","h","i","j","k","l","m","n","o","p","q","r","s","t","u","v","w","x","y","z"}
	var	r = ""
	for i := 0; i <= len(s)-1; i++{
		// get value
		c := getIndex(string(s[i]))
		// fmt.Println("--------")
		// fmt.Println(c)
		if c == 0 {
			// fmt.Println(string(s[i]))
			// r[i] += string(s[i]) error 寫法
			r += string(s[i])
		}else{

			// r[i] += h[c - 1] error 寫法
			r += h[c - 1]
			// fmt.Println(h[c - 1])
		}
	}
	return r
}


func main(){
	s := "leetcode"
	//fmt.Println(string(s[1]))
	fmt.Println(smallestString(s))
}