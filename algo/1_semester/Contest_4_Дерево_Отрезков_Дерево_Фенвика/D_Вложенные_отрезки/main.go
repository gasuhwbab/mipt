package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type segment struct {
	l, r int
}

type fenwick struct {
	f []int
}

func newFenwick(n int) *fenwick { return &fenwick{f: make([]int, n+2)} }

func (t *fenwick) add(pos, val int) {
	for pos < len(t.f) {
		t.f[pos] += val
		pos = pos | (pos + 1)
	}
}

func (t *fenwick) sum(pos int) int {
	s := 0
	for pos > 0 {
		s += t.f[pos]
		pos = pos&(pos+1) - 1
	}
	return s
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n int
	fmt.Fscan(in, &n)

	a := make([]segment, n)
	rs := make([]int, n)
	for i := range n {
		fmt.Fscan(in, &a[i].l, &a[i].r)
		rs[i] = a[i].r
	}

	sort.Slice(a, func(i, j int) bool {
		return a[i].l < a[j].l || (a[i].l == a[j].l && a[i].r > a[j].r)
	})

	sort.Ints(rs)
	uniq := rs[:0]
	for i, v := range rs {
		if i == 0 || v != rs[i-1] {
			uniq = append(uniq, v)
		}
	}
	idx := make(map[int]int)
	for i, v := range uniq {
		idx[v] = i + 1
	}

	f := newFenwick(len(uniq) + 2)
	var ans int64 = 0
	total := 0
	for i := 0; i < n; {
		j := i + 1
		for j < n && a[j].l == a[i].l && a[j].r == a[i].r {
			j++
		}
		k := j - i // same segments amount
		r := idx[a[i].r]
		prev := total - f.sum(r-1)
		ans += int64(k * prev)
		f.add(r, k)
		total += k
		i = j
	}
	fmt.Fprintln(out, ans)
}
