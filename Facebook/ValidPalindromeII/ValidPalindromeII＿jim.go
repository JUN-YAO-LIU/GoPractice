package main

import (
	"fmt"
	"sort"
)

// palindrome回文，兩個單或是三個同
func validPalindrome(s string) bool {
	var p,c,o,e uint8
	var m []int
 
	 if len(s) == 1{
		 return true
	 }
 
	 arr := []byte(s)
	 sort.Slice(arr,func(i,j int)bool{return arr[i]> arr[j]})
 
	 for i:=0; 0 <= len(arr) -1;i++{
 
		 if i == 0{
			m = append(m, 1)
			//  m[c] +=1
		 }
 
		 if i >0 && p == arr[i]{
			 m[c] +=1
		 }
 
		 if i >0 && p != arr[i]{
			m = append(m, 1)
			c++
			//  m[c] +=1
		 }
 
		 p = arr[i]
	 }
 
	 for i:=0; 0 <= len(m) -1;i++{
		 if m[i] % 2 ==1{
			 o++
		 }else{
			 e++
		 }
	 }
 
	 if e == 0 || (e > 0 && o >e){
		 return false
	 }
 
	 return true
 }


 func main(){

	a :="ab"
	fmt.Println(validPalindrome(a));

 }