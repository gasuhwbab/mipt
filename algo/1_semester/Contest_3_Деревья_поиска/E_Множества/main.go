package main

import (
	"bufio"
	"fmt"
	"os"
)

var seed uint64 = 1469598103934665603

func rnd() uint32 {
	seed ^= seed << 7
	seed ^= seed >> 9
	seed ^= seed << 8
	return uint32(seed)
}

type Node[T int | int64] struct {
	key         T
	prio        uint32
	left, right *Node[T]
}

func NewNode[T int | int64](key T) *Node[T] {
	return &Node[T]{key: key, prio: rnd()}
}

func (node *Node[T]) rotRE() *Node[T] {
	l := node.left
	node.left = l.right
	l.right = node
	return l
}

func (node *Node[T]) rotLE() *Node[T] {
	r := node.right
	node.right = r.left
	r.left = node
	return r
}

func (node *Node[T]) Insert(key T) *Node[T] {
	if node == nil {
		return NewNode(key)
	}
	if key < node.key {
		node.left = node.left.Insert(key)
		if node.left.prio > node.prio {
			node = node.rotRE()
		}
	} else if key > node.key {
		node.right = node.right.Insert(key)
		if node.right.prio > node.prio {
			node = node.rotLE()
		}
	}
	return node
}

func (node *Node[T]) Erase(key T) *Node[T] {
	if node == nil {
		return node
	}
	if key < node.key {
		node.left = node.left.Erase(key)
	} else if key > node.key {
		node.right = node.right.Erase(key)
	} else {
		if node.left == nil {
			return node.right
		}
		if node.right == nil {
			return node.left
		}
		if node.left.prio > node.right.prio {
			node = node.rotRE()
			node.right = node.right.Erase(key)
		} else {
			node = node.rotLE()
			node.left = node.left.Erase(key)
		}
	}
	return node
}

func (node *Node[T]) Find(key T) *Node[T] {
	cur := node
	for cur != nil {
		if key < cur.key {
			cur = cur.left
		} else if key > cur.key {
			cur = cur.right
		} else {
			return cur
		}
	}
	return nil
}

func (node *Node[T]) Contains(key T) bool {
	return node.Find(key) != nil
}

func (node *Node[T]) Inorder(res *[]T) {
	if node == nil {
		return
	}
	node.left.Inorder(res)
	*res = append(*res, node.key)
	node.right.Inorder(res)
}

type OutNode struct {
	key         int64
	prio        uint32
	left, right *OutNode
	sets        *Node[int]
}

func (n *OutNode) rotLE() *OutNode {
	r := n.right
	n.right = r.left
	r.left = n
	return r
}

func (n *OutNode) rotRE() *OutNode {
	l := n.left
	n.left = l.right
	l.right = n
	return l
}

func (n *OutNode) Insert(key int64) (*OutNode, *OutNode) {
	if n == nil {
		t := &OutNode{key: key, prio: rnd()}
		return t, t
	}
	var got *OutNode
	if key < n.key {
		n.left, got = n.left.Insert(key)
		if n.left.prio > n.prio {
			n = n.rotRE()
		}
	} else if key > n.key {
		n.right, got = n.right.Insert(key)
		if n.right.prio > n.prio {
			n = n.rotLE()
		}
	} else {
		got = n
	}
	return n, got
}

func (n *OutNode) Erase(key int64) *OutNode {
	if n == nil {
		return nil
	}
	if key < n.key {
		n.left = n.left.Erase(key)
	} else if key > n.key {
		n.right = n.right.Erase(key)
	} else {
		if n.left == nil {
			return n.right
		}
		if n.right == nil {
			return n.left
		}
		if n.left.prio > n.right.prio {
			n = n.rotRE()
			n.right = n.right.Erase(key)
		} else {
			n = n.rotLE()
			n.left = n.left.Erase(key)
		}
	}
	return n
}

func (n *OutNode) Find(key int64) *OutNode {
	cur := n
	for cur != nil {
		if key < cur.key {
			cur = cur.left
		} else if key > cur.key {
			cur = cur.right
		} else {
			return cur
		}
	}
	return nil
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var N int64
	var M, K int
	fmt.Fscan(in, &N, &M, &K)

	sets := make([]*Node[int64], M+1)
	var rev *OutNode = nil
	for range K {
		var q string
		fmt.Fscan(in, &q)
		switch q {
		case "ADD":
			var e int64
			var s int
			fmt.Fscan(in, &e, &s)
			if !sets[s].Contains(e) {
				sets[s] = sets[s].Insert(e)
				var node *OutNode
				rev, node = rev.Insert(e)
				node.sets = node.sets.Insert(s)
			}
		case "DELETE":
			var e int64
			var s int
			fmt.Fscan(in, &e, &s)
			sets[s] = sets[s].Erase(e)
			if node := rev.Find(e); node != nil {
				node.sets = node.sets.Erase(s)
				if node.sets == nil {
					rev = rev.Erase(e)
				}
			}
		case "CLEAR":
			var s int
			fmt.Fscan(in, &s)
			var elems []int64
			sets[s].Inorder(&elems)
			for _, e := range elems {
				if node := rev.Find(e); node != nil {
					node.sets = node.sets.Erase(s)
					if node.sets == nil {
						rev = rev.Erase(e)
					}
				}
			}
			sets[s] = nil
		case "LISTSET":
			var s int
			fmt.Fscan(in, &s)
			var elems []int64
			sets[s].Inorder(&elems)
			if len(elems) == 0 {
				fmt.Fprintln(out, -1)
			} else {
				for i, x := range elems {
					if i > 0 {
						fmt.Fprint(out, " ")
					}
					fmt.Fprint(out, x)
				}
				fmt.Fprintln(out)
			}

		case "LISTSETSOF":
			var e int64
			fmt.Fscan(in, &e)
			node := rev.Find(e)
			if node == nil || node.sets == nil {
				fmt.Fprintln(out, -1)
				continue
			}
			var idxs []int
			node.sets.Inorder(&idxs)
			if len(idxs) == 0 {
				fmt.Fprintln(out, -1)
			} else {
				for i, v := range idxs {
					if i > 0 {
						fmt.Fprint(out, " ")
					}
					fmt.Fprint(out, v)
				}
				fmt.Fprintln(out)
			}
		}
	}
}
