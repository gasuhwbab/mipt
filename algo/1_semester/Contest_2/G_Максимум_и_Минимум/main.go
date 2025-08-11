package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Pair struct {
	ind int
	v   int
}

type maxHeap []Pair

func (h maxHeap) Len() int           { return len(h) }
func (h maxHeap) Less(i, j int) bool { return h[i].v > h[j].v }
func (h maxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *maxHeap) Push(x any)        { *h = append(*h, x.(Pair)) }
func (h *maxHeap) Pop() any {
	old := *h
	ans := old[len(old)-1]
	*h = old[:len(old)-1]
	return ans
}

type minHeap []Pair

func (h minHeap) Len() int           { return len(h) }
func (h minHeap) Less(i, j int) bool { return h[i].v < h[j].v }
func (h minHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *minHeap) Push(x any)        { *h = append(*h, x.(Pair)) }
func (h *minHeap) Pop() any {
	old := *h
	ans := old[len(old)-1]
	*h = old[:len(old)-1]
	return ans
}
func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n int
	fmt.Fscan(in, &n)

	maxH := &maxHeap{}
	heap.Init(maxH)
	minH := &minHeap{}
	heap.Init(minH)

	removed := make([]bool, n+1)
	nextId := 0

	for range n {
		var req string
		fmt.Fscan(in, &req)
		if strings.Contains(req, "Insert(") {
			x, _ := strconv.Atoi(req[7 : len(req)-1])
			id := nextId
			nextId++
			heap.Push(maxH, Pair{ind: id, v: x})
			heap.Push(minH, Pair{ind: id, v: x})
		} else if req == "GetMin" {
			for minH.Len() > 0 {
				x := heap.Pop(minH).(Pair)
				if removed[x.ind] {
					continue
				}
				removed[x.ind] = true
				fmt.Fprintln(out, x.v)
				break
			}
		} else {
			for maxH.Len() > 0 {
				x := heap.Pop(maxH).(Pair)
				if removed[x.ind] {
					continue
				}
				removed[x.ind] = true
				fmt.Fprintln(out, x.v)
				break
			}
		}
	}
}
