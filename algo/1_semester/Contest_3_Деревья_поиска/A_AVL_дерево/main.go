package main

import (
	"bufio"
	"fmt"
	"os"
)

type Node struct {
	key         int
	left, right *Node
	height      int
}

func NewNode(key int) *Node { return &Node{key: key} }

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func (node *Node) GetMin() *Node {
	if node == nil {
		return nil
	}
	if node.left == nil {
		return node
	}
	return node.left.GetMin()
}

func (node *Node) GetMax() *Node {
	if node == nil {
		return nil
	}
	if node.right == nil {
		return node
	}
	return node.right.GetMax()
}

func (node *Node) getHeight() int {
	if node == nil {
		return -1
	}
	return node.height
}

func (node *Node) updateHeight() {
	node.height = max(node.right.getHeight(), node.left.getHeight()) + 1
}

func (node *Node) leftRotate() *Node {
	if node == nil || node.right == nil {
		return node
	}
	r := node.right
	buf := r.left
	r.left = node
	node.right = buf
	node.updateHeight()
	r.updateHeight()
	return r
}

func (node *Node) rightRotate() *Node {
	if node == nil || node.left == nil {
		return node
	}
	l := node.left
	buf := l.right
	l.right = node
	node.left = buf
	node.updateHeight()
	l.updateHeight()
	return l
}

func (node *Node) getBalance() int {
	return node.right.getHeight() - node.left.getHeight()
}

func (node *Node) balance() *Node {
	if node == nil {
		return nil
	}
	node.updateHeight()
	delta := node.getBalance()
	if delta == 2 {
		if node.right != nil && node.right.getBalance() < 0 {
			node.right = node.right.rightRotate()
		}
		return node.leftRotate()
	} else if delta == -2 {
		if node.left != nil && node.left.getBalance() > 0 {
			node.left = node.left.leftRotate()
		}
		return node.rightRotate()
	}
	return node
}

func (node *Node) Insert(key int) *Node {
	if node == nil {
		return NewNode(key)
	}
	if node.key == key {
		return node
	} else if key < node.key {
		node.left = node.left.Insert(key)
	} else {
		node.right = node.right.Insert(key)
	}
	return node.balance()
}

func (node *Node) Erase(key int) *Node {
	if node == nil {
		return nil
	}
	if key > node.key {
		node.right = node.right.Erase(key)
	} else if key < node.key {
		node.left = node.left.Erase(key)
	} else {
		if node.right == nil {
			return node.left
		}
		if node.left == nil {
			return node.right
		}
		minInRight := node.right.GetMin()
		node.key = minInRight.key
		node.right = node.right.Erase(minInRight.key)
	}
	return node.balance()
}

func (node *Node) Exists(key int) bool {
	if node == nil {
		return false
	}
	if node.key == key {
		return true
	} else if key < node.key {
		return node.left.Exists(key)
	} else {
		return node.right.Exists(key)
	}
}

func (node *Node) Find(key int) *Node {
	if node == nil {
		return nil
	}
	if node.key == key {
		return node
	} else if key < node.key {
		return node.left.Find(key)
	} else {
		return node.right.Find(key)
	}
}

func (node *Node) Next(key int) *Node {
	var ans *Node
	cur := node
	for cur != nil {
		if cur.key > key {
			ans = cur
			cur = cur.left
		} else {
			cur = cur.right
		}
	}
	return ans
}

func (node *Node) Prev(key int) *Node {
	var ans *Node
	cur := node
	for cur != nil {
		if cur.key < key {
			ans = cur
			cur = cur.right
		} else {
			cur = cur.left
		}
	}
	return ans
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var avl *Node

	for {
		var q string
		if _, err := fmt.Fscan(in, &q); err != nil {
			break
		}

		switch q {
		case "insert":
			var x int
			fmt.Fscan(in, &x)
			avl = avl.Insert(x)
		case "delete":
			var x int
			fmt.Fscan(in, &x)
			avl = avl.Erase(x)
		case "exists":
			var x int
			fmt.Fscan(in, &x)
			if avl.Exists(x) {
				fmt.Fprintln(out, "true")
			} else {
				fmt.Fprintln(out, "false")
			}
		case "next":
			var x int
			fmt.Fscan(in, &x)
			ans := avl.Next(x)
			if ans == nil {
				fmt.Fprintln(out, "none")
			} else {
				fmt.Fprintln(out, ans.key)
			}

		case "prev":
			var x int
			fmt.Fscan(in, &x)
			ans := avl.Prev(x)
			if ans == nil {
				fmt.Fprintln(out, "none")
			} else {
				fmt.Fprintln(out, ans.key)
			}
		}
	}
}

// 2 5 3
