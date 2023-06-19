package main

import (
	"fmt"
	"sort"
)


// 邏輯寫錯了，不能排列
// palindrome回文，兩個單或是三個同
func validPalindrome(s string)bool {
	var p,c,o,e uint8
	var m []int

	if len(s) == 1{
		return true
	}
 
	 arr := []byte(s)
	 sort.Slice(arr,func(i,j int)bool{return arr[i]> arr[j]})

	// fmt.Println(arr);
 
	 for i:=0; i <= len(arr) -1;i++{
 
		 if i == 0{
			m = append(m, 1)
			
			//  m[c] +=1
		 }

		//  fmt.Println(m);
		//  fmt.Println(p);
 
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
 
	 for i:=0; i <= len(m) -1;i++{
		 if m[i] % 2 ==1{
			 o++
		 }else{
			 e++
		 }
	 }
	 
	fmt.Println(o);
	fmt.Println(e);

	r1 := true
	r2 := true

	if e < 0 || (e > 0 && o >= e){
		r1 = false
	}

	fmt.Println(r1);

	if o > 0{
		o--
	}else{
		e--
	}

	if e < 0 || (e > 0 && o >e) || (e ==0 && o >0){
		r2 = false
	}

	fmt.Println(r2);

	if r1 != r2{
		return false
	}else if r1{
		return true
	}else{
		return false
	}
 
	return true
 }


 func main(){

	a :="tebbem"
	fmt.Println(validPalindrome(a));

	// var p = 'C'
	// fmt.Println(p);

 }