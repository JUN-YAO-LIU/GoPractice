package main

import(
	"fmt"
	"sort"
)

// 檢討
// string to byte array 排列字串裡面的字元
// []byte array
// sort Slice 用法
// insert node 時候比較要填入哪邊節點怎麼寫
// 1. constuctor 是兩個function 2. 結構可以.function 3. string 不行

// bst insert and struct
// constructor
// for
// sort

// bst insert and struct
type Node struct{
	Value int
	Left *Node
	Right *Node
}

func InsertNode(root *Node,val int) *Node{
	// 錯
	// var node *Node

	if root ==nil{
		return &Node{Value:val,Left:nil,Right:nil}
	}

	// 錯
	// if root.Left.Value > val{
	// 	node = InsertNode(root.Left,val)
	// }else{
	// 	node = InsertNode(root.Right,val)
	// }

	if root.Value > val{
		root.Left = InsertNode(root.Left,val)
	}else{
		root.Right = InsertNode(root.Right ,val)
	}

	return root
}

func OrderNode(root *Node){
	if root != nil {
		fmt.Printf("%d ", root.Value)
        OrderNode(root.Left)
        OrderNode(root.Right)
    }
}

// constructor
type Init struct{
	value int
	name string
}

func Contructor(n string,p int) Init{
	return Init{name:n,value:p}
}

// 錯的
// func Contructor(this *Init) Get(n1 string, add int) Init{
// 	this.name = name
// 	this.value = add
// 	return this
// }

func (this *Init) Get(n1 string, add int) *Init{
	this.name = n1
	this.value = add
	return this
}


func main(){
	// node
	var node *Node

	node = InsertNode(node,10)
	InsertNode(node,20)
	InsertNode(node,30)
	OrderNode(node)
	fmt.Println(node)

	// sort
	var s_sort = "bac"
	// byteArr := byte(strArr)
	byteArr := []byte(s_sort)
	// 忘記
	sort.Slice(byteArr,func(i,j int) bool { return byteArr[i] > byteArr[j]})
	fmt.Println(string(byteArr))

	// for
	var strArr = []string{"asv","df","bb"}
	for key,value := range strArr{
		fmt.Println(key,value)
		fmt.Println("-------f------")
		fmt.Printf("%d %s",key,value)
	}

	fmt.Println("-------------")

	// constructor
	// 初始化錯誤
	// var testInit *Init
	var testInit = Init{name:"Jim",value:123}
	fmt.Println(testInit.Get("JIm",100))
	fmt.Println(s_sort.ChangeStr("JIm"))
}