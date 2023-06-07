package main

import(
	"fmt"
)

type Node struct{
	value int
	left *Node
	right *Node
}


func Insert(node *Node, val int) *Node{
	if node == nil{
		// node.value = val
		// node.left = nil
		// node.right = nil
		return &Node{value: val, left: nil, right: nil}
	}

	if node.left.value > val{
		Insert(node.right,val)
	}else{
		Insert(node.left,val)
	}

	return node
}

func main(){
	var node *Node

	node =Insert(node,20)
	fmt.Println(*node)
	fmt.Println(&node)
	fmt.Println(node)
}