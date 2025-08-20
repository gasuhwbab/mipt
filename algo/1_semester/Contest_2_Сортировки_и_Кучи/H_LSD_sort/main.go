package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n int
	fmt.Fscan(in, &n)
	arr := make([]uint64, n)
	for i := range n {
		fmt.Fscan(in, &arr[i])
	}
	m := getMax(arr)
	radixSort(arr, m, len(arr))
	for i := range arr {
		if i > 0 {
			fmt.Fprint(out, " ")
		}
		fmt.Fprint(out, arr[i])
	}
	fmt.Fprintln(out)
}

func radixSort(arr []uint64, m uint64, n int) {
	var exp uint64 = 1
	for ; m/exp > 0; exp *= 10 {
		countSort(arr, n, exp)
	}
}

func countSort(arr []uint64, n int, exp uint64) {
	output := make([]uint64, n)
	cnt := make([]int, 10)
	for i := 0; i < n; i++ {
		cnt[(arr[i]/exp)%10]++
	}
	for i := 1; i < 10; i++ {
		cnt[i] += cnt[i-1]
	}
	for i := n - 1; i >= 0; i-- {
		output[cnt[(arr[i]/exp)%10]-1] = arr[i]
		cnt[(arr[i]/exp)%10]--
	}
	for i := range n {
		arr[i] = output[i]
	}
}

func getMax(a []uint64) uint64 {
	var maxx uint64
	for _, x := range a {
		maxx = max(maxx, x)
	}
	return maxx
}
