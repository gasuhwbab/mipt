package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type query struct {
	typ byte
	x   int
}

type fenwick struct {
	t []int
}

func NewFenwick(n int) *fenwick { return &fenwick{t: make([]int, n)} }

func (f *fenwick) add(i, x int) {
	for i < len(f.t) {
		f.t[i] += x
		i = i | (i + 1)
	}
}

func (f *fenwick) sum(i int) int {
	s := 0
	for i >= 0 {
		s += f.t[i]
		i = i&(i+1) - 1
	}
	return s
}

func upperBound(a []int, x int) int {
	return sort.Search(len(a), func(i int) bool { return a[i] > x })
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var q int
	fmt.Fscan(in, &q)
	queries := make([]query, q)
	vals := make([]int, q)
	for i := range q {
		var typ string
		var x int
		fmt.Fscan(in, &typ, &x)
		queries[i].typ, queries[i].x = typ[0], x
		vals = append(vals, x)
	}
	sort.Ints(vals)
	uniq := vals[:0]
	for i, v := range vals {
		if i == 0 || v != vals[i-1] {
			uniq = append(uniq, v)
		}
	}
	ind := make(map[int]int)
	for i, v := range uniq {
		ind[v] = i
	}

	f := NewFenwick(len(uniq))
	for _, q := range queries {
		if q.typ == '+' {
			f.add(ind[q.x], q.x)
		} else {
			k := upperBound(uniq, q.x) - 1
			if k == 0 {
				fmt.Fprintln(out, 0)
			} else {
				fmt.Fprintln(out, f.sum(k))
			}
		}
	}
}
