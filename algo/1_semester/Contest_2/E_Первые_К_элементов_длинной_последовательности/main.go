package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
	"sort"
)

type Heap []uint32

func (h Heap) Len() int {
	return len(h)
}

func (h Heap) Less(i, j int) bool {
	return h[i] > h[j]
}

func (h Heap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *Heap) Push(x any) {
	*h = append(*h, x.(uint32))
}

func (h *Heap) Pop() any {
	old := *h
	n := len(old)
	ans := old[n-1]
	*h = old[:n-1]
	return ans
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n, k int
	fmt.Fscan(in, &n, &k)
	var a, x, y uint64
	fmt.Fscan(in, &a, &x, &y)
	mod := uint64((1 << 30) - 1)
	h := &Heap{}
	heap.Init(h)
	for range n {
		a = (a*x + y) & mod
		v := uint32(a)
		if h.Len() < k {
			heap.Push(h, v)
		} else if k > 0 && v < (*h)[0] {
			heap.Pop(h)
			heap.Push(h, v)
		}
	}
	res := make([]uint32, 0)
	for h.Len() > 0 {
		res = append(res, heap.Pop(h).(uint32))
	}
	sort.Slice(res, func(i, j int) bool {
		return res[i] < res[j]
	})
	for i, v := range res {
		if i > 0 {
			fmt.Fprint(out, " ")
		}
		fmt.Fprint(out, v)
	}
	fmt.Fprintln(out)
}
