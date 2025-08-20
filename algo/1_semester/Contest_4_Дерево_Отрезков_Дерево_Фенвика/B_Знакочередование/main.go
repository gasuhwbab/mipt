package main

import (
	"bufio"
	"fmt"
	"os"
)

type tree struct {
	a []int
	n int
}

func NewTree(n int) *tree {
	return &tree{n: n, a: make([]int, 4*n)}
}

func (t *tree) Build(arr []int) { t.build(arr, 1, 0, t.n-1) }
func (t *tree) build(arr []int, v, tl, tr int) {
	if tl == tr {
		t.a[v] = arr[tl]
		return
	}
	tm := (tl + tr) / 2
	t.build(arr, 2*v, tl, tm)
	t.build(arr, 2*v+1, tm+1, tr)
	t.a[v] = t.a[2*v] + t.a[2*v+1]
}

func (t *tree) Update(pos, val int) { t.update(1, 0, t.n-1, pos, val) }
func (t *tree) update(v, tl, tr, pos, val int) {
	if tl == tr {
		t.a[v] = val
		return
	}
	tm := (tl + tr) / 2
	if pos <= tm {
		t.update(2*v, tl, tm, pos, val)
	} else {
		t.update(2*v+1, tm+1, tr, pos, val)
	}
	t.a[v] = t.a[2*v] + t.a[2*v+1]
}

func (t *tree) GetSum(l, r int) int { return t.getSum(1, 0, t.n-1, l, r) }
func (t *tree) getSum(v, tl, tr, l, r int) int {
	if l > r {
		return 0
	}
	if tl == l && tr == r {
		return t.a[v]
	}
	tm := (tl + tr) / 2
	return t.getSum(2*v, tl, tm, l, min(r, tm)) + t.getSum(2*v+1, tm+1, tr, max(l, tm+1), r)
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n int
	fmt.Fscan(in, &n)
	a := make([]int, n)
	for i := range n {
		var x int
		fmt.Fscan(in, &x)
		if (i+1)%2 == 1 {
			a[i] = x
		} else {
			a[i] = -x
		}
	}

	t := NewTree(n)
	t.Build(a)
	var m int
	fmt.Fscan(in, &m)
	for ; m > 0; m-- {
		var op int
		fmt.Fscan(in, &op)
		if op == 0 {
			var i, j int
			fmt.Fscan(in, &i, &j)
			if i%2 == 1 {
				t.Update(i-1, j)
			} else {
				t.Update(i-1, -j)
			}
		}
		if op == 1 {
			var l, r int
			fmt.Fscan(in, &l, &r)
			sum := t.GetSum(l-1, r-1)
			if l%2 == 0 {
				sum = -sum
			}
			fmt.Fprintln(out, sum)
		}
	}
}
