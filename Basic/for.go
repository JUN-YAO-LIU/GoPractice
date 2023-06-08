package main

import "fmt"

var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

func main() {

   for i, v := range pow {
      fmt.Printf("2**%d = %d\n", i, v)
   }

    map1 := make(map[int]float32)
    map1[1] = 7.0
    map1[2] = 8.0
    map1[3] = 9.0
    map1[4] = 41.0
   
    for key, value := range map1 {
      fmt.Printf("key is: %d - value is: %f\n", key, value)
    }

    for key := range map1 {
      fmt.Printf("key is: %d\n", key)
    }

    for _, value := range map1 {
      fmt.Printf("value is: %f\n", value)
    }

	// 這樣有順序性，如果用map range可能會沒有順序
	for i := 1; i <= len(map1);i++{
		fmt.Printf("value is: %f\n", map1[i])
	}
}