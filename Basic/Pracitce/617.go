package main 

import(
	"fmt"
)

//Next ： 樹的其他建法

// 題目：
// map用法
// node
// constructor
// string trim
// string to char
// array
// leetcode題目：2 4 6 9


// fail : 不用,
type Node struct{
	value int
	left *Node
	right *Node
}

// node 忘記
func InsertNode(root *Node,val int) *Node{

	// 沒有根了，自己就是頭
	if root == nil{
		return &Node{value: val,left:nil,right:nil}
	}

	// node 節點比較
	if root.value < val{
		root.left = InsertNode(root.left,val)
	}else{
		root.right = InsertNode(root.right,val)
	}

	return root
}

func TravelNode(root *Node){
	// fail : 少判斷 root == nil
	if root != nil{
		// 其他root往下走
		TravelNode(root.left)
		// fmt.Println("left :",root.left)
		fmt.Println("node value :",root.value)

		TravelNode(root.right)
		// fmt.Println("rignt :",root.right)
	}
	
}

// constructor
type StructValue struct{
	value int
}

// fail : 完全忘記 
// func contructor (n value) (StructValue){
// 	return Init{value:n}
// }

// 2 Dot Product of two sparse vectors
type SparseVector struct {
	arr []int
}

func Constructor(nums []int) SparseVector {
	var a SparseVector
	a.arr = nums
	return a
}

// Return the dotProduct of two sparse vectors
func (this *SparseVector) dotProduct(vec SparseVector) int {
	var result int
	for key,value := range this.arr{
		result += value * vec.arr[key]
	}

	return result
}

/**
 * Your SparseVector object will be instantiated and called as such:
 * v1 := Constructor(nums1);
 * v2 := Constructor(nums2);
 * ans := v1.dotProduct(v2);
 */


func main(){

	fmt.Println("-----Map-----")
	// fail : 需要{}
	map1 := map[string]int{}
	map1["test"] = 1
	fmt.Println("map1 :",map1)

	// fail : 每一個都需要,就算是最後一個也一樣
	map2 := map[string]string{
		"test":"testStr",
		"test01":"test01str",
	}
	fmt.Println("map2 :",map2)

	map3 := make(map[int]string,3)
	map3[1] = "01"
	fmt.Println("map3 :",map3)

	fmt.Println("-----Node-----")
	var node *Node
	node = InsertNode(node,100)
	InsertNode(node,20)
	InsertNode(node,30)
	TravelNode(node)


	fmt.Println("-----Leet Code 1570-----")
	leetcodeNum1570 := []int{0,9,2}
	leetcodeNum1570_1 := []int{3,1,2}
	v1 := Constructor(leetcodeNum1570);
	v2 := Constructor(leetcodeNum1570_1);
	fmt.Println(v1.dotProduct(v2))

	fmt.Println("-----Leet Code 1570-----")
}