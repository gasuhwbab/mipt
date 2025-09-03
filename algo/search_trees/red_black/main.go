package main

type Node struct {
	key, val            int
	left, right, parent *Node
	color               bool // true - red; false - black
}

func NewNode(key, val int) *Node {
	return &Node{key: key, val: val, color: true}
}

 