package main

import (
	"bufio"
	"fmt"
	"os"
)

const (
	INF = int(^uint(0) >> 1)
)

type Node struct {
	val int
	ind int
}

type SpaseTable struct {
	st [][]Node
	lg []int
	n  int
}

func NewSparceTable(a []int) *SpaseTable {
	n := len(a)
	lg := make([]int, n+1)
	for i := 2; i <= n; i++ {
		lg[i] = lg[i/2] + 1
	}
	K := lg[n]
	st := make([][]Node, K+1)
	st[0] = make([]Node, n)
	for i := range a {
		st[0][i] = Node{a[i], i}
	}
	for k := 1; k <= K; k++ {
		w := 1 << k
		half := w >> 1
		st[k] = make([]Node, n-w+1)
		for i := 0; i+w <= n; i++ {
			st[k][i] = merge(st[k-1][i], st[k-1][i+half])
		}
	}
	return &SpaseTable{st: st, lg: lg, n: n}
}

func (sparse *SpaseTable) query(l, r int) Node {
	if l > r {
		return Node{val: INF, ind: -1}
	}
	k := sparse.lg[r-l+1]
	return merge(sparse.st[k][l], sparse.st[k][r-(1<<k)+1])
}

func merge(a, b Node) Node {
	if a.val < b.val {
		return a
	}
	if b.val < a.val {
		return b
	}
	if a.ind <= b.ind {
		return a
	}
	return b
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n, m int
	fmt.Fscan(in, &n, &m)
	a := make([]int, n)
	for i := range n {
		fmt.Fscan(in, &a[i])
	}
	st := NewSparceTable(a)
	for ; m > 0; m-- {
		var l, r int
		fmt.Fscan(in, &l, &r)
		l--
		r--
		minNode := st.query(l, r)
		left := st.query(l, minNode.ind-1).val
		right := st.query(minNode.ind+1, r).val
		ans := min(left, right)
		if ans == INF {
			fmt.Fprintln(out, minNode.val)
		} else {
			fmt.Fprintln(out, ans)
		}
	}
}
