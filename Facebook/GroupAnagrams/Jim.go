package main

import (
	"fmt"
	"sort"
)

// 第一個存入第一層array
// 新的string 與全部的array比較一次 一樣就加入，不一樣就再一層
// 比較方式是 迴圈執行這一層裡面的每一第一個string，判斷長度、字母是不是都有
func groupAnagrams(strs []string) [][]string {
	var result = [][]string{}
	x := 0
	y := 0


    for i := 0 ;i <= len(strs) - 1 ; i++ {

		if i == 0 {
			result[x][y] = strs[i]
		}

		if i > 0 {
			for c := 0;c <= x; c++{
				temp := result[c][0]
				if temp == strs[i]{
					// 加上去，不知道有沒有方法
					result[c]
				}else{

				}
			}
		}

	}

	return result
}