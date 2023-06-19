// 宣告程式屬於哪個 package
package main

// 引入套件
import (
    "fmt"
)


// 程式執行入口
func main() {
	fmt.Println(romanToInt("III"))
    fmt.Println("---------------")
}


func romanToInt(s string) int {
	var v, lv, cv int
	h := map[uint8]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

    for i := len(s) - 1; i >= 0; i-- {
		cv = h[s[i]]
		if cv < lv {
			v -= cv
		} else {
			v += cv
		}
		lv = cv
	}

	return v
}
