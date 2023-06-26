package main

import (
	"fmt"
	"strconv"
)

func validWordAbbreviation(word string, abbr string) bool {
    i := 0
    j := 0
    for i < len(word) && j < len(abbr){
       if word[i] == abbr[j]{
           i++
           j++
           continue
       }
       if(string(abbr[j]) <= "0" || string(abbr[j]) > "9"){
           return false
       }
       start := j
       for j < len(abbr) && string(abbr[j]) >= "0" && string(abbr[j]) <= "9"{
           j++
       }
       temp, _ := strconv.Atoi(abbr[start:j])
       i += temp
    }
    return i == len(word) && j == len(abbr)
}


func main(){
	w:= "a"
	abbr:= "01"
	fmt.Println(validWordAbbreviation(w,abbr))
	// validWordAbbreviation(s,s1)
}