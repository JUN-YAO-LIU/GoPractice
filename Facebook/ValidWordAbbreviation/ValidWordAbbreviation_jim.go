package main

import (
	"fmt"
	"strconv"
)

func validWordAbbreviation(abbr string) {
    var strs []string
    tempStr:= ""
	// p := ""

    for i:=0 ;i <= len(abbr) -1; i++{
		fmt.Println("num:")
		fmt.Println(i)
        if _, err := strconv.Atoi(string(abbr[i])); err == nil {
			fmt.Println("----if----")
			fmt.Println(string(abbr[i]))
            tempStr += string(abbr[i])
        }else{
			if tempStr != ""{
				strs = append(strs,tempStr)
            	tempStr = ""
			}
            
            strs = append(strs,string(abbr[i]))
			fmt.Println("----else----")
			fmt.Println(strs)
        }
       // p := string(abbr[i])
    }

	for 
}


func main(){
	// s:= "internationalization"
	s1:= "i12iz4n"
	// fmt.Println(validWordAbbreviation(s,s1))
	validWordAbbreviation(s1)
}