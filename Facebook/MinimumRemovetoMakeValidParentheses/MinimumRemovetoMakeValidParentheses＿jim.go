package main

import (
	"fmt"
)

// ())
// )()
// ))()
func minRemoveToMakeValid(s string) string {
	var left,right []int
	result := s


	for key,val := range s{
		if string(val) == "("{
			left = append(left,key)
		}

		if string(val) == ")"{
			right = append(right,key)
		}
	}


	if len(left) == 0 && len(right) == 0{
		return result
	}

	if len(left) > 0 && len(right) == 0{
		return result
	}

	if len(left) == 0 && len(right) > 0{
		return result
	}

    fmt.Println(left)
	fmt.Println(right)

	// 消除一個沒配對到的，如果都有就不用remove
	
	if len(left) != len(right){
		var remove int
		if len(left) > len(right){
			// left = append(left[1:len(left)])
			remove = left[0]
		}

		// 不用減1
		if len(left) < len(right){
			// right = append(right[1:len(right)])
			remove = right[0]
		}

		char := []rune(s)
    	res := delChar(char, remove)
		result = string(res)
	}

	

	// fmt.Println("----------")
	// fmt.Println(left)
	// fmt.Println(right)
	
	// 判斷有沒有pair，方向有沒有對
	if left[len(left)-1] > right[len(right)-1]{
		result = ""
	}
	

	return result
}


func delChar(s []rune, index int) []rune {
    return append(s[0:index], s[index+1:]...)
}


func main(){
	s :=")))))"
	
	fmt.Println("-----ANS-----")
	fmt.Println(minRemoveToMakeValid(s))

	fmt.Println("-----ANS-----")
	s1 :="lee(t(c)o)de)"
	fmt.Println(minRemoveToMakeValid(s1))


	fmt.Println("-----ANS-----")
	s2 :="a)b(c)d"
	fmt.Println(minRemoveToMakeValid(s2))

	fmt.Println("-----ANS-----")
	s3 :="))(("
	fmt.Println(minRemoveToMakeValid(s3))
	// minRemoveToMakeValid(s1)
}