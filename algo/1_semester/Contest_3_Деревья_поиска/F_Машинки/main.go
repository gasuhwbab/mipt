package main

import (
	"bufio"
	"fmt"
	"os"
)

type Key struct {
	next int
	id   int
}

func isEqual(a, b Key) bool {
	return a.id == b.id && a.next == b.next
}

func less(a, b Key) bool {
	return a.next < b.next || (a.next == b.next && a.id < b.id)
}

var seed uint64 = 17748612893471908218

func rnd() uint32 {
	seed ^= seed << 7
	seed ^= seed >> 9
	seed ^= seed << 8
	return uint32(seed)
}

type Node struct {
	key         Key
	prio        uint32
	left, right *Node
}

func NewNode(key Key) *Node {
	return &Node{key: key, prio: rnd()}
}

func rotLE(node *Node) *Node {
	r := node.right
	node.right = r.left
	r.left = node
	return r
}

func rotRE(node *Node) *Node {
	l := node.left
	node.left = l.right
	l.right = node
	return l
}

func insert(node *Node, key Key) *Node {
	if node == nil {
		return NewNode(key)
	}
	if less(key, node.key) {
		node.left = insert(node.left, key)
		if node.left.prio > node.prio {
			node = rotRE(node)
		}
	} else {
		node.right = insert(node.right, key)
		if node.right.prio > node.prio {
			node = rotLE(node)
		}
	}
	return node
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
		return a
	} else {
		b.left = merge(a, b.left)
		return b
	}
}

func erase(node *Node, key Key) *Node {
	if node == nil {
		return node
	}
	if isEqual(node.key, key) {
		return merge(node.left, node.right)
	}
	if less(key, node.key) {
		node.left = erase(node.left, key)
	} else {
		node.right = erase(node.right, key)
	}
	return node
}

func extractMax(node *Node) (*Node, Key) {
	if node.right == nil {
		return node.left, node.key
	}
	var key Key
	node.right, key = extractMax(node.right)
	return node, key
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var N, K, P int
	fmt.Fscan(in, &N, &K, &P)
	sequence := make([]int, P)
	pos := make([][]int, N+1)
	for i := range P {
		fmt.Fscan(in, &sequence[i])
		pos[sequence[i]] = append(pos[sequence[i]], i)
	}

	ptr := make([]int, N+1)

	const INF = int(1e9)

	nextAfterUse := func(id int) int {
		ptr[id]++
		if ptr[id] < len(pos[id]) {
			return pos[id][ptr[id]]
		}
		return INF
	}

	currNext := make([]int, N+1)
	onFloor := make([]bool, N+1)

	var node *Node
	floorCnt := 0
	cnt := 0

	for i := range P {
		x := sequence[i]

		if onFloor[x] {
			oldKey := Key{next: currNext[x], id: x}
			node = erase(node, oldKey)
			nxt := nextAfterUse(x)
			currNext[x] = nxt
			node = insert(node, Key{next: nxt, id: x})
		} else {
			cnt++
			if floorCnt == K {
				var key Key
				node, key = extractMax(node)
				onFloor[key.id] = false
				floorCnt--
			}
			nxt := nextAfterUse(x)
			currNext[x] = nxt
			node = insert(node, Key{next: nxt, id: x})
			onFloor[x] = true
			floorCnt++
		}
	}
	fmt.Fprintln(out, cnt)
}
