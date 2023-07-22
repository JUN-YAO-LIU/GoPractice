package main


import (
	"fmt"
)

func isValid(s string) bool {
	// 需要相互對到()
	// 不能{]
	// 不能換位置
	// ({)123} 這樣是對還是錯? 要false
	// 只有符號

	// 不能 left_1,left_2 :=0
	left_1 := 0
	left_2 := 0
	left_3 := 0

	// 解法：
    for _,val := range s{
		if string(val) == "("{
			left_1++
		}

		if string(val) == ")"{
			if left_1 <=0 {
				return false
			}
			left_1--
		}

		if string(val) == "{"{
			left_2++
		}

		if string(val) == "}"{
			if left_2 <=0{
				return false
			}
			left_2--
		}

		if string(val) == "["{
			left_3++
		}

		if string(val) == "]"{
			if left_3 <=0{
				return false
			}
			left_3--
		}
	}

	if left_1 > 0 || left_2 > 0 || left_3 > 0{
		return false
	}


	var arr []string
	arr = append(arr,string(s[0]))
	for i := 1 ; i < len(s) -1 ; i++{
		if string(s[i]) == ")"{
			c,b := getFirstIndex(arr,"(")

			if !b{
				return false
			}

			arr = remove(arr,c)
		}

		if string(s[i]) == "}"{
			c,b := getFirstIndex(arr,"{")

			if !b{
				return false
			}

			arr = remove(arr,c)
		}

		if string(s[i]) == "]"{
			c,b := getFirstIndex(arr,"[")

			if !b{
				return false
			}

			arr = remove(arr,c)
		}


		arr = append(arr,string(s[i]))
	}

	return true
}

func getFirstIndex(a []string,target string) (int,bool){
	first := 0
	for key, val := range a{
		if val == target{
			first = key
		}

		if key > first{
			return first,false
		}
	}

	return first,false
}

func remove(s []string, i int) []string {
    s[i] = s[len(s)-1]
    return s[:len(s)-1]
}

func main(){
	// str :="({[])}"
	// str :="({)}"
	// str :="()}"
	// str :="()"
	// str :="({})"
	str :="(){}"
	fmt.Println(isValid(str))
}
