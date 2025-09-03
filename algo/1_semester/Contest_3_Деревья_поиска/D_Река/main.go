package main

import (
	"bufio"
	"fmt"
	"os"
)

var seed uint64 = 17848981204124578281

func rnd() uint32 {
	seed ^= seed >> 7
	seed ^= seed << 9
	seed ^= seed >> 8
	return uint32(seed)
}

type Node struct {
	key         int64
	sz          int
	prio        uint32
	left, right *Node
}

func NewNode(key int64) *Node {
	return &Node{key: key, sz: 1, prio: rnd()}
}

func size(node *Node) int {
	if node == nil {
		return 0
	}
	return node.sz
}

func upd(node *Node) {
	if node != nil {
		node.sz = 1 + size(node.left) + size(node.right)
	}
}

func split(node *Node, k int) (*Node, *Node) {
	if node == nil {
		return nil, nil
	}
	if size(node.left) >= k {
		l, r := split(node.left, k)
		node.left = r
		upd(node)
		return l, node
	} else {
		l, r := split(node.right, k-size(node.left)-1)
		node.right = l
		upd(node)
		return node, r
	}
}

func merge(a, b *Node) *Node {
	if a == nil {
		return b
	}
	if b == nil {
		return a
	}
	if a.prio > b.prio {
		a.right = merge(a.right, b)
		upd(a)
		return a
	} else {
		b.left = merge(a, b.left)
		upd(b)
		return b
	}
}

func leftMost(node *Node) *Node {
	if node == nil {
		return nil
	}
	for node.left != nil {
		node = node.left
	}
	return node
}

func rightMost(node *Node) *Node {
	if node == nil {
		return nil
	}
	for node.right != nil {
		node = node.right
	}
	return node
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n, p int
	fmt.Fscan(in, &n, &p)
	var root *Node
	var sumXX int64 = 0
	for range n {
		var x int64
		fmt.Fscan(in, &x)
		root = merge(root, NewNode(x))
		sumXX += x * x
	}
	fmt.Fprintln(out, sumXX)
	var k int
	fmt.Fscan(in, &k)
	for range k {
		var e, v int
		fmt.Fscan(in, &e, &v)
		a, bc := split(root, v-1)
		b, c := split(bc, 1)
		l := b.key
		if e == 1 {
			if a != nil && c != nil {
				l1 := l / 2
				l2 := l - l1
				aRight := rightMost(a)
				sumXX -= aRight.key * aRight.key
				aRight.key += l1
				sumXX += aRight.key * aRight.key

				cLeft := leftMost(c)
				sumXX -= cLeft.key * cLeft.key
				cLeft.key += l2
				sumXX += cLeft.key * cLeft.key
			} else if a != nil {
				aRight := rightMost(a)
				sumXX -= aRight.key * aRight.key
				aRight.key += l
				sumXX += aRight.key * aRight.key
			} else {
				cLeft := leftMost(c)
				sumXX -= cLeft.key * cLeft.key
				cLeft.key += l
				sumXX += cLeft.key * cLeft.key
			}
			sumXX -= l * l
			root = merge(a, c)
		} else {
			l1 := l / 2
			l2 := l - l1
			sumXX += l1*l1 + l2*l2 - l*l
			root = merge(a, merge(NewNode(l1), merge(NewNode(l2), c)))
		}
		fmt.Fprintln(out, sumXX)
	}
}
