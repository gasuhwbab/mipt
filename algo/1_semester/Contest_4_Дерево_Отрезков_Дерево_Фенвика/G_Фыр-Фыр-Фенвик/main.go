package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type BIT struct {
	xs  []int
	ys  [][]int
	bit [][]int
	n   int
}

func lowerBound(a []int, x int) int {
	return sort.Search(len(a), func(i int) bool { return a[i] >= x })
}

func upperBound(a []int, x int) int {
	return sort.Search(len(a), func(i int) bool { return a[i] > x })
}

func unique(a []int) []int {
	if len(a) == 0 {
		return a
	}
	sort.Ints(a)
	j := 1
	for i := 1; i < len(a); i++ {
		if a[i] != a[i-1] {
			a[j] = a[i]
			j++
		}
	}
	return a[:j]
}

func NewBIT(pointsX, pointsY []int) *BIT {
	xs := make([]int, 0)
	xs = append(xs, pointsX...)
	sort.Ints(xs)
	xs = unique(xs)
	n := len(xs)

	ys := make([][]int, n)
	for i := range pointsX {
		xi := lowerBound(xs, pointsX[i])
		for j := xi; j < n; j = j | (j + 1) {
			ys[j] = append(ys[j], pointsY[i])
		}
	}

	bit := make([][]int, n)
	for i := range n {
		ys[i] = unique(ys[i])
		bit[i] = make([]int, len(ys[i]))
	}
	return &BIT{xs: xs, ys: ys, bit: bit, n: n}
}

func (t *BIT) add(x, y, val int) {
	for xi := x; xi < t.n; xi = xi | (xi + 1) {
		ys := t.ys[xi]
		pos := lowerBound(ys, y)
		b := t.bit[xi]
		for j := pos; j < len(b); j = j | (j + 1) {
			b[j] += val
		}
	}
}

func (t *BIT) sum(rx, ry int) int {
	xi := upperBound(t.xs, rx) - 1
	if xi < 0 {
		return 0
	}
	res := 0
	for ; xi >= 0; xi = xi&(xi+1) - 1 {
		yi := upperBound(t.ys[xi], ry) - 1
		if yi < 0 {
			continue
		}
		b := t.bit[xi]
		for j := yi; j >= 0; j = j&(j+1) - 1 {
			res += b[j]
		}
	}
	return res
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n int
	fmt.Fscan(in, &n)
	x := make([]int, n)
	y := make([]int, n)
	w := make([]int, n)
	for i := range n {
		fmt.Fscan(in, &x[i], &y[i], &w[i])
	}

	b := NewBIT(x, y)

	for i := range n {
		xi := lowerBound(b.xs, x[i])
		b.add(xi, y[i], w[i])
	}

	var m int
	fmt.Fscan(in, &m)
	for ; m > 0; m-- {
		var query string
		fmt.Fscan(in, &query)
		if query == "get" {
			var rx, ry int
			fmt.Fscan(in, &rx, &ry)
			fmt.Fprintln(out, b.sum(rx, ry))
		} else {
			var i, z int
			fmt.Fscan(in, &i, &z)
			i--
			diff := z - w[i]
			w[i] = z
			xi := lowerBound(b.xs, x[i])
			b.add(xi, y[i], diff)
		}
	}
}
