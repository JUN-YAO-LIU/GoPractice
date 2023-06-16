package main

import (
	"fmt"
	"strconv"
)

func validWordAbbreviation(word string, abbr string) bool {
    var strs []string
    tempStr:= ""

    for i:=0 ;i <= len(abbr) -1; i++{
		// fmt.Println("num:")
		// fmt.Println(i)
        if _, err := strconv.Atoi(string(abbr[i])); err == nil {
			// fmt.Println("----if----")
			// fmt.Println(string(abbr[i]))
            tempStr += string(abbr[i])
        }else{
			if tempStr != ""{
				strs = append(strs,tempStr)
            	tempStr = ""
			}
            
            strs = append(strs,string(abbr[i]))
			// fmt.Println("----else----")
			// fmt.Println(strs)
        }
    }

	fmt.Println(len(strs))
	for i:=0 ;i <= len(strs) -1; i++{
		fmt.Println("執行次數 : ",i)

		cNum, err := strconv.Atoi(string(strs[i]));
		if string(word[0]) != strs[i] && err != nil {
			return false
		}

		if err == nil {
			// fmt.Println(word[cNum:len(word)-1])
			word = word[cNum:len(word)]
		}else{
			// fmt.Println(word[1:len(word)])
			word = word[1:len(word)]
		}
	} 

	return true
}


func main(){
	w:= "substitution"
	abbr:= "s55n"
	fmt.Println(validWordAbbreviation(w,abbr))
	// validWordAbbreviation(s,s1)
}