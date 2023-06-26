package main

import (
	"fmt"
	"strconv"
)

func validWordAbbreviation(word string, abbr string) bool {
    var strs []string

	// 前一個數字
    pStr:= ""

	// 字元 現在數字 前一個數字
    for i:=0 ;i <= len(abbr) -1; i++{
		_, err := strconv.Atoi(string(abbr[i]));

		if err == nil {
			pStr += string(abbr[i])
        }else{
			_, errp := strconv.Atoi(pStr);
			if errp == nil{
				strs = append(strs,pStr)
				pStr = ""
			}
			strs = append(strs,string(abbr[i]))
		}
    }

	if len(pStr) >0{
		strs = append(strs,pStr)
		pStr = ""
	}

	fmt.Println(strs)
	// fmt.Println(len(strs))
	for i:=0 ;i <= len(strs) -1; i++{
		fmt.Println("執行次數 : ",i)

		cNum, err := strconv.Atoi(string(strs[i]));
		if string(word[0]) != strs[i] && err != nil {
			return false
		}

		if err == nil {
			// fmt.Println(word[cNum:len(word)-1])
			if cNum > len(word){
				return false
			}
			word = word[cNum:len(word)]
		}else{
			// fmt.Println(word[1:len(word)])
			word = word[1:len(word)]
		}
	} 

	return true
}


func main(){
	w:= "a"
	abbr:= "01"
	fmt.Println(validWordAbbreviation(w,abbr))
	// validWordAbbreviation(s,s1)
}