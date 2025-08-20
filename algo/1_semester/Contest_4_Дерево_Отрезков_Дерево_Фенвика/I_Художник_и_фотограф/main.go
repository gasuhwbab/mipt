package main

import (
	"bufio"
	"fmt"
	"os"
)

const (
	NONE = -1
	INF  = 1 << 32
)

type SegTree struct {
	mn   []int
	lazy []int
	n    int
}

func NewSegTree(a []int) *SegTree {
	n := len(a)
	sgt := &SegTree{mn: make([]int, 4*n), lazy: make([]int, 4*n), n: n}
	for i := range sgt.lazy {
		sgt.lazy[i] = NONE
	}
	sgt.build(a, 1, 0, sgt.n-1)
	return sgt
}

func (sgt *SegTree) build(a []int, v int, tl, tr int) {
	if tl == tr {
		sgt.mn[v] = a[tl]
		return
	}
	tm := (tl + tr) / 2
	sgt.build(a, 2*v, tl, tm)
	sgt.build(a, 2*v+1, tm+1, tr)
	sgt.mn[v] = min(sgt.mn[2*v], sgt.mn[2*v+1])
}

func (sgt *SegTree) push(v int) {
	if sgt.lazy[v] != NONE {
		val := sgt.lazy[v]
		sgt.lazy[2*v] = val
		sgt.lazy[2*v+1] = val
		sgt.mn[2*v] = val
		sgt.mn[2*v+1] = val
		sgt.lazy[v] = NONE
	}
}

func (sgt *SegTree) Update(l, r, val int) { sgt.update(1, 0, sgt.n-1, l, r, val) }
func (sgt *SegTree) update(v, tl, tr, l, r, val int) {
	if l > r {
		return
	}
	if l == tl && r == tr {
		sgt.mn[v] = val
		sgt.lazy[v] = val
		return
	}
	sgt.push(v)
	tm := (tl + tr) / 2
	if r <= tm {
		sgt.update(2*v, tl, tm, l, r, val)
	} else if l > tm {
		sgt.update(2*v+1, tm+1, tr, l, r, val)
	} else {
		sgt.update(2*v, tl, tm, l, tm, val)
		sgt.update(2*v+1, tm+1, tr, tm+1, r, val)
	}
	sgt.mn[v] = min(sgt.mn[2*v], sgt.mn[2*v+1])
}

func (sgt *SegTree) Query(l, r int) int { return sgt.query(1, 0, sgt.n-1, l, r) }
func (sgt *SegTree) query(v, tl, tr, l, r int) int {
	if l > r {
		return INF
	}
	if l == tl && r == tr {
		return sgt.mn[v]
	}
	sgt.push(v)
	tm := (tl + tr) / 2
	if r <= tm {
		return sgt.query(2*v, tl, tm, l, r)
	} else if l > tm {
		return sgt.query(2*v+1, tm+1, tr, l, r)
	} else {
		return min(sgt.query(2*v, tl, tm, l, tm), sgt.query(2*v+1, tm+1, tr, tm+1, r))
	}
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n int
	fmt.Fscan(in, &n)
	a := make([]int, n)
	for i := range n {
		var R, G, B int
		fmt.Fscan(in, &R, &G, &B)
		a[i] = R + G + B
	}
	sgt := NewSegTree(a)

	var k int
	fmt.Fscan(in, &k)
	for i := range k {
		if i != 0 {
			fmt.Fprint(out, " ")
		}
		var c, d, r, g, b, e, f int
		fmt.Fscan(in, &c, &d, &r, &g, &b, &e, &f)
		val := r + g + b
		sgt.Update(c, d, val)
		fmt.Fprint(out, sgt.Query(e, f))
	}
}
