package main

// splay

import (
	"bufio"
	"fmt"
	"os"
)

type Node struct {
	key                 int64
	parent, left, right *Node
	size                int
}

func NewNode(key int64) *Node {
	return &Node{key: key, size: 1}
}

func (node *Node) recalc() {
	if node == nil {
		return
	}
	s := 1
	if node.left != nil {
		s += node.left.size
	}
	if node.right != nil {
		s += node.right.size
	} 
	node.size = s
}

func (node *Node) rotateRight() *Node {
	if node == nil || node.left == nil {
		return node
	}
	p := node.parent
	l := node.left
	if p != nil {
		if p.right == node {
			p.right = l
		} else {
			p.left = l
		}
	}
	tmp := l.right
	l.right = node
	node.left = tmp
	node.parent = l
	l.parent = p
	if node.left != nil {
		node.left.parent = node
	}
	node.recalc()
	l.recalc()
	return l
}

func (node *Node) rotateLeft() *Node {
	if node == nil || node.right == nil {
		return node
	}
	p := node.parent
	r := node.right
	if p != nil {
		if p.right == node {
			p.right = r
		} else {
			p.left = r
		}
	}
	tmp := r.left
	r.left = node
	node.right = tmp
	node.parent = r
	r.parent = p
	if node.right != nil {
		node.right.parent = node
	}
	node.recalc()
	r.recalc()
	return r
}

func (node *Node) splay() *Node {
	if node == nil {
		return nil
	}
	for node.parent != nil {
		p := node.parent
		g := p.parent
		if g == nil {
			if p.left == node {
				p.rotateRight()
			} else {
				p.rotateLeft()
			}
		} else if g.left == p && p.left == node {
			g.rotateRight()
			p.rotateRight()
		} else if g.right == p && p.right == node {
			g.rotateLeft()
			p.rotateLeft()
		} else if g.left == p && p.right == node {
			p.rotateLeft()
			g.rotateRight()
		} else {
			p.rotateRight()
			g.rotateLeft()
		}
	}
	return node
}

func (node *Node) getMax() *Node {
	if node == nil {
		return node
	}
	for node.right != nil {
		node = node.right
	}
	return node
}

func (node *Node) Find(key int64) *Node {
	cur := node
	var last *Node
	for cur != nil {
		last = cur
		if key < cur.key {
			cur = cur.left
		} else if key > cur.key {
			cur = cur.right
		} else {
			return cur.splay()
		}
	}
	if last != nil {
		return last.splay()
	}
	return last
}

func Merge(node1, node2 *Node) *Node {
	if node1 == nil {
		return node2
	}
	if node2 == nil {
		return node1
	}
	m := node1.getMax()
	node1 = m.splay()
	node1.right = node2
	node2.parent = node1
	node1.recalc()
	return node1
}

func (node *Node) Split(key int64) (*Node, *Node) {
	if node == nil {
		return nil, nil
	}
	node = node.Find(key)
	if node.key <= key {
		r := node.right
		if r != nil {
			r.parent = nil
		}
		node.right = nil
		node.recalc()
		return node, r
	} else {
		l := node.left
		if l != nil {
			l.parent = nil
		}
		node.left = nil
		node.recalc()
		return l, node
	}
}

func (node *Node) Insert(key int64) *Node {
	if node == nil {
		return NewNode(key)
	}
	l, r := node.Split(key)
	if l != nil && l.key == key {
		return Merge(l, r)
	}
	ans := NewNode(key)
	ans.left = l
	if l != nil {
		l.parent = ans
	}
	ans.right = r
	if r != nil {
		r.parent = ans
	}
	ans.recalc()
	return ans
}

func (node *Node) Erase(key int64) *Node {
	if node == nil {
		return node
	}
	node = node.Find(key)
	if node.key != key {
		return node
	}
	l, r := node.left, node.right
	if l != nil {
		l.parent = nil
	}
	if r != nil {
		r.parent = nil
	}
	node.left, node.right, node.parent = nil, nil, nil
	return Merge(l, r)
}

func (node *Node) kth(k int) *Node {
	if node == nil || k <= 0 || k > node.size {
		return node
	}
	cur := node
	for {
		lsize := 0
		if cur.left != nil {
			lsize = cur.left.size
		}
		if k == lsize+1 {
			return cur.splay()
		} else if k <= lsize {
			cur = cur.left
		} else {
			k -= (lsize + 1)
			cur = cur.right
		}
	}
	return nil
}

func (node *Node) KthMax(k int) *Node {
	if node == nil {
		return nil
	}
	i := node.size - k + 1
	return node.kth(i)
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var splay *Node
	var n int
	fmt.Fscan(in, &n)
	for i := 0; i < n; i++ {
		var c int
		var k int64
		fmt.Fscan(in, &c, &k)
		switch c {
		case 1:
			splay = splay.Insert(k)
		case 0:
			splay = splay.KthMax(int(k))
			fmt.Fprintln(out, splay.key)
		case -1:
			splay = splay.Erase(k)
		}
	}
}
