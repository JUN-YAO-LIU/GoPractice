// 宣告程式屬於哪個 package
package main

// 引入套件
import (
    "fmt"
)

// 程式執行入口
func main() {
    // 使用 fmt 套件印出字串 word，使用 := 簡化原本變數宣告 var word string = "Hello, World!"
    word := "Hello, World!"
    fmt.Println(word)

    // 但輸入到function裡面就不行
    arr1 := [4]string{"Volvo", "BMW", "Ford", "Mazda"}
    fmt.Println(arr1)
}