package binarytree

import "fmt"

type Node struct {
	Value int
	//parent *Node
	left  *Node
	right *Node
}

func (N *Node) insertData(content int) {

	if content < N.Value {
		if N.left == nil {
			N.left = &Node{Value: content}
		} else {
			N.left.insertData(content)
		}
	} else {
		if N.right == nil {
			N.right = &Node{Value: content}
		} else {
			N.right.insertData(content)
		}
	}

}

func preOrdertraversal(Root Node) {

}

func (N *Node) InOrder() {

	if N.left != nil {
		N.left.InOrder()
	}
	fmt.Print(N.Value, " ")
	if N.right != nil {
		N.right.InOrder()
	}
}

func ImplementBinaryTree() {
	var n, data int
	fmt.Println("Enter the number of elements to be inserted")
	fmt.Scan(&n)
	fmt.Println("Enter the ", n, " data:")

	var Root Node
	for i := 0; i < n; i++ {
		fmt.Scan(&data)
		if i == 0 {
			Root.Value = data
		} else {
			Root.insertData(data)
		}

	}

	fmt.Print("[")
	Root.InOrder()
	fmt.Print("]")
	//preOrdertraversal

}
