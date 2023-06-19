// Note : 刪除最小數量的 patentheses然後回傳字串

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

	for key,val := range left{
		if right[key] < val{
			char := []rune(s)
    		result = delChar(char, val)
		}
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