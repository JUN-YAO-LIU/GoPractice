// Note : 要做到 type.function，這個type就要是struct
// func (this *sturct) Dosomething() (int){
//
//}

package main

import (
	"fmt"
)

// go做出constuctor的效果
type SparseVector struct {
    nums []int
}

// factory
func Constructor(nums []int) SparseVector {
	var s SparseVector;
	s.nums = nums
    return s
}

// Return the dotProduct of two sparse vectors
func (this *SparseVector) dotProduct(vec SparseVector) int {
    
}

/**
 * Your SparseVector object will be instantiated and called as such:
 * v1 := Constructor(nums1);
 * v2 := Constructor(nums2);
 * ans := v1.dotProduct(v2);
 */

 // sparse 疏、稀
 // vectors 載體
 // sparse vector
 func main(){
	fmt.Printf("ass")
 }