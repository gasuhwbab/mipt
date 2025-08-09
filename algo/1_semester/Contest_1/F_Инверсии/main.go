package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var in *bufio.Reader
	var out *bufio.Writer
	in = bufio.NewReader(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n int
	fmt.Fscan(in, &n)
	a := make([]int64, n)
	for i := range n {
		fmt.Fscan(in, &a[i])
	}
	var ans int64
	merge := func(l, r []int64) []int64 {
		to := make([]int64, len(l)+len(r))
		for i, j := 0, 0; i < len(l) || j < len(r); {
			if j == len(r) || (i < len(l) && l[i] < r[j]) {
				to[i+j] = l[i]
				i++
				ans += int64(j)
			} else {
				to[i+j] = r[j]
				j++
			}
		}
		return to
	}
	var mergeSort func(arr []int64) []int64
	mergeSort = func(arr []int64) []int64 {
		if len(arr) <= 1 {
			return arr
		}
		mid := len(arr) / 2
		l := mergeSort(arr[:mid])
		r := mergeSort(arr[mid:])
		return merge(l, r)
	}
	mergeSort(a)
	fmt.Fprintln(out, ans)
}
