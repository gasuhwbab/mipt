package main

type Node struct {
	key, val int
	left     *Node
	right    *Node
}

func NewNode(key, val int) *Node {
	return &Node{key: key, val: val}
}

func (node *Node) Insert(key, val int) {
	if node == nil {
		node = NewNode(key, val)
		return
	}
	if key == node.key {
		return
	} else if key < node.key {
		node.left.Insert(key, val)
	} else {
		node.right.Insert(key, val)
	}
}

func (node *Node) Find(key int) *Node {
	if node == nil {
		return nil
	} else if node.key == key {
		return node
	} else if key < node.key {
		return node.left.Find(key)
	} else {
		return node.right.Find(key)
	}
}

func (node *Node) GetMin() *Node {
	if node == nil {
		return nil
	} else if node.left == nil {
		return node
	} else {
		return node.left.GetMin()
	}
}

func (node *Node) GetMax() *Node {
	if node == nil {
		return nil
	} else if node.right == nil {
		return node
	} else {
		return node.right.GetMax()
	}
}

func (node *Node) Erase(key int) {
	if node == nil {
		return
	} else if key < node.key {
		node.left.Erase(key)
	} else if key > node.key {
		node.right.Erase(key)
	} else {
		if node.left == nil || node.right == nil {
			if node.left != nil {
				node = node.left
			} else {
				node = node.right
			}
		} else {
			minInRight := node.right.GetMin()
			node.key = minInRight.key
			node.val = minInRight.val
			node.right.Erase(minInRight.key)
		}
	}
}
