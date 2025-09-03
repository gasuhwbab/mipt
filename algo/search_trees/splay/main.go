// // bottom-up
// package main

// import (
// 	"bufio"
// 	"fmt"
// 	"os"
// )

// type Node struct {
// 	key                 int
// 	sum                 int
// 	parent, left, right *Node
// }

// func NewNode(key int) *Node { return &Node{key: key, sum: key} }

// func (node *Node) recalc() {
// 	if node == nil {
// 		return
// 	}
// 	s := node.key
// 	if node.left != nil {
// 		s += node.left.sum
// 	}
// 	if node.right != nil {
// 		s += node.right.sum
// 	}
// 	node.sum = s
// }

// func (node *Node) rotateLeft() *Node {
// 	if node == nil || node.right == nil {
// 		return node
// 	}
// 	p := node.parent
// 	r := node.right
// 	if p != nil {
// 		if p.left == node {
// 			p.left = r
// 		} else {
// 			p.right = r
// 		}
// 	}
// 	tmp := r.left
// 	r.left = node
// 	node.right = tmp
// 	node.parent = r
// 	r.parent = p
// 	if node.right != nil {
// 		node.right.parent = node
// 	}
// 	node.recalc()
// 	r.recalc()
// 	return r
// }

// func (node *Node) rotateRight() *Node {
// 	if node == nil || node.left == nil {
// 		return node
// 	}
// 	p := node.parent
// 	l := node.left
// 	if p != nil {
// 		if p.right == node {
// 			p.right = l
// 		} else {
// 			p.left = l
// 		}
// 	}
// 	tmp := l.right
// 	l.right = node
// 	node.left = tmp
// 	node.parent = l
// 	l.parent = p
// 	if node.left != nil {
// 		node.left.parent = node
// 	}
// 	node.recalc()
// 	l.recalc()
// 	return l
// }

// func (node *Node) splay() *Node {
// 	if node == nil {
// 		return nil
// 	}
// 	for node.parent != nil {
// 		p := node.parent
// 		g := node.parent.parent
// 		if g == nil {
// 			//zig
// 			if node == p.right {
// 				p.rotateLeft()
// 			} else {
// 				p.rotateRight()
// 			}
// 		} else if node == p.left && p == g.left {
// 			// zig-zig
// 			g.rotateRight()
// 			p.rotateRight()
// 		} else if node == p.right && p == g.right {
// 			//zig-zig
// 			g.rotateLeft()
// 			p.rotateLeft()
// 		} else if node == p.right && p == g.left {
// 			//zig-zag
// 			p.rotateLeft()
// 			g.rotateRight()
// 		} else {
// 			//zig-zag
// 			p.rotateRight()
// 			g.rotateLeft()
// 		}
// 	}
// 	return node
// }

// func (node *Node) getMax() *Node {
// 	if node == nil {
// 		return nil
// 	}
// 	for node.right != nil {
// 		node = node.right
// 	}
// 	return node
// }

// func (node *Node) getMin() *Node {
// 	if node == nil {
// 		return nil
// 	}
// 	for node.left != nil {
// 		node = node.left
// 	}
// 	return node
// }

// func (node *Node) Find(key int) *Node {
// 	cur := node
// 	var last *Node
// 	for cur != nil {
// 		last = cur
// 		if key < cur.key {
// 			cur = cur.left
// 		} else if key > cur.key {
// 			cur = cur.right
// 		} else {
// 			return cur.splay()
// 		}
// 	}
// 	if last != nil {
// 		return last.splay()
// 	}
// 	return last
// }

// func (node *Node) Split(key int) (*Node, *Node) {
// 	if node == nil {
// 		return nil, nil
// 	}
// 	node = node.Find(key)
// 	if node.key <= key {
// 		r := node.right
// 		if r != nil {
// 			r.parent = nil
// 		}
// 		node.right = nil
// 		node.recalc()
// 		return node, r
// 	} else {
// 		l := node.left
// 		if l != nil {
// 			l.parent = nil
// 		}
// 		node.left = nil
// 		node.recalc()
// 		return l, node

// 	}
// }

// func Merge(node1 *Node, node2 *Node) *Node {
// 	if node1 == nil {
// 		return node2
// 	}
// 	if node2 == nil {
// 		return node1
// 	}
// 	m := node1.getMax()
// 	node1 = m.splay()
// 	node1.right = node2
// 	node2.parent = node1
// 	node1.recalc()
// 	return node1
// }

// func (node *Node) Insert(key int) *Node {
// 	if node == nil {
// 		return NewNode(key)
// 	}
// 	l, r := node.Split(key)
// 	if l != nil && node.key == key {
// 		return Merge(l, r)
// 	}
// 	ans := NewNode(key)
// 	ans.left = l
// 	if l != nil {
// 		l.parent = ans
// 	}
// 	ans.right = r
// 	if r != nil {
// 		r.parent = ans
// 	}
// 	ans.recalc()
// 	return ans
// }

// func (node *Node) Erase(key int) *Node {
// 	if node == nil {
// 		return nil
// 	}
// 	node = node.Find(key)
// 	if node.key != key {
// 		return node
// 	}
// 	l, r := node.left, node.right
// 	if l != nil {
// 		l.parent = nil
// 	}
// 	if r != nil {
// 		r.parent = nil
// 	}
// 	node.left, node.right, node.parent = nil, nil, nil
// 	return Merge(l, r)
// }

// func (root *Node) sumLE(x int) (int, *Node) {
// 	if root == nil {
// 		return 0, nil
// 	}
// 	root = root.Find(x) // splay ближайшего к x в корень
// 	res := 0
// 	if root.left != nil {
// 		res += root.left.sum
// 	}
// 	if root.key <= x {
// 		res += root.key
// 	}
// 	return res, root
// }

// func (root *Node) Sum(l, r int) (int, *Node) {
// 	if root == nil || l > r {
// 		return 0, root
// 	}
// 	s1, root := root.sumLE(r)
// 	s0, root := root.sumLE(l - 1)
// 	return s1 - s0, root
// }

// func main() {
// 	in := bufio.NewReader(os.Stdin)
// 	out := bufio.NewWriter(os.Stdout)
// 	defer out.Flush()

// 	const MOD int = 1_000_000_000

// 	var splay *Node
// 	isOk := false
// 	last := 0

//		var n int
//		fmt.Fscan(in, &n)
//		for i := 0; i < n; i++ {
//			var q string
//			fmt.Fscan(in, &q)
//			if q == "+" {
//				var x int
//				fmt.Fscan(in, &x)
//				if isOk {
//					x = (x + last) % MOD
//				}
//				splay = splay.Insert(x)
//				isOk = false
//			} else {
//				var l, r int
//				fmt.Fscan(in, &l, &r)
//				sum, root := splay.Sum(l, r)
//				splay = root
//				fmt.Fprintln(out, sum)
//				last = sum
//				isOk = true
//			}
//		}
//	}
package main

import (
	"bufio"
	"fmt"
	"os"
)

type Node struct {
	key                 int64
	sum                 int64
	parent, left, right *Node
}

func NewNode(key int64) *Node { return &Node{key: key, sum: key} }

func (node *Node) recalc() {
	if node == nil {
		return
	}
	s := node.key
	if node.left != nil {
		s += node.left.sum
	}
	if node.right != nil {
		s += node.right.sum
	}
	node.sum = s
}

func (node *Node) rotateLeft() *Node {
	if node == nil || node.right == nil {
		return node
	}
	p := node.parent
	r := node.right
	if p != nil {
		if p.left == node {
			p.left = r
		} else {
			p.right = r
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

func (node *Node) splay() *Node {
	if node == nil {
		return nil
	}
	for node.parent != nil {
		p := node.parent
		g := p.parent
		if g == nil {
			if node == p.right {
				p.rotateLeft()
			} else {
				p.rotateRight()
			}
		} else if node == p.left && p == g.left {
			g.rotateRight()
			p.rotateRight()
		} else if node == p.right && p == g.right {
			g.rotateLeft()
			p.rotateLeft()
		} else if node == p.right && p == g.left {
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
	for node != nil && node.right != nil {
		node = node.right
	}
	return node
}

func (node *Node) Find(key int64) *Node {
	cur, last := node, (*Node)(nil)
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
	return nil
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
	}
	l := node.left
	if l != nil {
		l.parent = nil
	}
	node.left = nil
	node.recalc()
	return l, node
}

func Merge(a, b *Node) *Node {
	if a == nil {
		return b
	}
	if b == nil {
		return a
	}
	m := a.getMax()
	a = m.splay()
	a.right = b
	b.parent = a
	a.recalc()
	return a
}

// ВАЖНО: множество — дубликаты НЕ вставляем
func (node *Node) Insert(key int64) *Node {
	if node == nil {
		return NewNode(key)
	}
	l, r := node.Split(key)
	if l != nil && l.key == key { // уже есть ключ
		return Merge(l, r) // просто склеиваем обратно
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
		return nil
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

func (root *Node) sumLE(x int64) (int64, *Node) {
	if root == nil {
		return 0, nil
	}
	root = root.Find(x)
	res := int64(0)
	if root.left != nil {
		res += root.left.sum
	}
	if root.key <= x {
		res += root.key
	}
	return res, root
}

func (root *Node) Sum(l, r int64) (int64, *Node) {
	if root == nil || l > r {
		return 0, root
	}
	s1, root := root.sumLE(r)
	s0, root := root.sumLE(l - 1)
	return s1 - s0, root
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	const MOD int64 = 1_000_000_000

	var splay *Node
	var n int
	fmt.Fscan(in, &n)

	var last int64 = 0
	prevWasSum := false

	for i := 0; i < n; i++ {
		var q string
		fmt.Fscan(in, &q)
		if q == "+" {
			var x int64
			fmt.Fscan(in, &x)
			if prevWasSum {
				x = (x + last) % MOD
			}
			splay = splay.Insert(x)
			prevWasSum = false
		} else { // "? l r"
			var l, r int64
			fmt.Fscan(in, &l, &r)
			var sum int64
			sum, splay = splay.Sum(l, r)
			fmt.Fprintln(out, sum)
			last = sum
			prevWasSum = true
		}
	}
}
