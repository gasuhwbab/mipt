package main

import (
	"bufio"
	"fmt"
	"os"
)

var seed uint64 = 14768731904924204997

func rnd() uint32 {
	seed ^= seed << 7
	seed ^= seed >> 9
	seed ^= seed << 8
	return uint32(seed)
}

type Node struct {
	key         int
	cnt         int
	prio        uint32
	left, right *Node
}

type Treap struct {
	root *Node
}

func NewNode(key int) *Node { return &Node{key: key, prio: rnd(), cnt: 1} }

func rotLE(node *Node) *Node {
	if node == nil || node.right == nil {
		return node
	}
	r := node.right
	node.right = r.left
	r.left = node
	return r
}

func rotRE(node *Node) *Node {
	if node == nil || node.left == nil {
		return node
	}
	l := node.left
	node.left = l.right
	l.right = node
	return l
}

func insert(node *Node, key int) *Node {
	if node == nil {
		return NewNode(key)
	}
	if key < node.key {
		node.left = insert(node.left, key)
		if node.left.prio > node.prio {
			return rotRE(node)
		}
	} else if key > node.key {
		node.right = insert(node.right, key)
		if node.right.prio > node.prio {
			return rotLE(node)
		}
	} else {
		node.cnt++
		return node
	}
	return node
}
func (t *Treap) Insert(key int) { t.root = insert(t.root, key) }

func eraseOne(node *Node, key int) *Node {
	if node == nil {
		return nil
	}
	if key < node.key {
		node.left = eraseOne(node.left, key)
		return node
	}
	if key > node.key {
		node.right = eraseOne(node.right, key)
		return node
	}
	if node.cnt > 1 {
		node.cnt--
		return node
	}
	if node.left == nil {
		return node.right
	}
	if node.right == nil {
		return node.left
	}
	if node.left.prio > node.right.prio {
		node = rotRE(node)
		node.right = eraseOne(node.right, key)
	} else {
		node = rotLE(node)
		node.left = eraseOne(node.left, key)
	}
	return node
}
func (t *Treap) EraseOne(key int) { t.root = eraseOne(t.root, key) }

func (t *Treap) Max() int {
	c := t.root
	if c == nil {
		return 0
	}
	for c.right != nil {
		c = c.right
	}
	return c.key
}

func (t *Treap) Prev(key int) (int, bool) {
	c := t.root
	ok := false
	ans := 0
	for c != nil {
		if key <= c.key {
			c = c.left
		} else {
			ok = true
			ans = c.key
			c = c.right
		}
	}
	return ans, ok
}

func (t *Treap) Next(key int) (int, bool) {
	c := t.root
	ok := false
	ans := 0
	for c != nil {
		if key >= c.key {
			c = c.right
		} else {
			ok = true
			ans = c.key
			c = c.left
		}
	}
	return ans, ok
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var N, Y, X int
	fmt.Fscan(in, &N, &Y, &X)

	adds := make([][]int, Y+2)
	dels := make([][]int, Y+2)
	for range N {
		var x, y1, y2 int
		fmt.Fscan(in, &x, &y1, &y2)
		adds[y1] = append(adds[y1], x)
		dels[y2+1] = append(dels[y2+1], x)
	}

	s := &Treap{}
	g := &Treap{}

	freq := make(map[int]int)

	s.Insert(0)
	s.Insert(X)
	freq[0] = 1
	freq[X] = 1
	g.Insert(X)

	ans := make([]int, 0)
	for y := range Y + 1 {
		for _, x := range dels[y] {
			freq[x]--
			if freq[x] == 0 {
				p, _ := s.Prev(x)
				n, _ := s.Next(x)
				g.EraseOne(x - p)
				g.EraseOne(n - x)
				g.Insert(n - p)
				s.EraseOne(x)
			}
		}
		for _, x := range adds[y] {
			if freq[x] == 0 {
				p, _ := s.Prev(x)
				n, _ := s.Next(x)
				g.EraseOne(n - p)
				g.Insert(x - p)
				g.Insert(n - x)
				s.Insert(x)
			}
			freq[x]++
		}
		ans = append(ans, g.Max())
	}
	for i := range ans {
		fmt.Fprintln(out, ans[i])
	}
}
