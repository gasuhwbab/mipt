package main

import (
	"bufio"
	"fmt"
	"os"
)

type SegTree struct {
	t     []int
	n, sz int
}

func NewTree(a []int) *SegTree {
	n := len(a)
	sz := 1
	for sz < n {
		sz <<= 1
	}
	t := make([]int, 2*sz)
	for i := range t {
		t[i] = -1
	}
	for i := range a {
		t[sz+i] = a[i]
	}
	for i := sz - 1; i >= 1; i-- {
		t[i] = max(t[i*2], t[i*2+1])
	}
	return &SegTree{t: t, n: n, sz: sz}
}

func (st *SegTree) Set(i, x int) {
	j := st.sz + i - 1
	st.t[j] = x
	for j >>= 1; j >= 1; j >>= 1 {
		st.t[j] = max(st.t[j*2], st.t[j*2+1])
		if j == 1 {
			break
		}
	}
}

func (st *SegTree) Get(i, x int) int {
	if i > st.n {
		return -1
	}
	j := st.sz + i - 1
	if st.t[j] >= x {
		return i
	}
	for j > 1 {
		if j%2 == 0 {
			right := j + 1
			if st.t[right] >= x {
				j = right
				for j < st.sz {
					left := 2 * j
					if st.t[left] >= x {
						j = left
					} else {
						j = left + 1
					}
				}
				res := j - st.sz + 1
				if res <= st.n {
					return res
				}
				return -1
			}
		}
		j /= 2
	}
	return -1
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
	st := NewTree(a)
	for ; m > 0; m-- {
		var op, i, x int
		fmt.Fscan(in, &op, &i, &x)
		if op == 0 {
			st.Set(i, x)
		} else {
			fmt.Fprintln(out, st.Get(i, x))
		}
	}
}
