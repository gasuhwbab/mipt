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
	heights := []int{}
	for range n {
		h := 0
		fmt.Fscan(in, &h)
		heights = append(heights, h)
	}
	stack := []int{}
	maxS := 0
	for i := 0; i <= n; i++ {
		currH := 0
		if i < n {
			currH = heights[i]
		}
		for len(stack) > 0 && heights[stack[len(stack)-1]] > currH {
			h := heights[stack[len(stack)-1]]
			stack = stack[:len(stack)-1]
			w := i
			if len(stack) > 0 {
				w = i - stack[len(stack)-1] - 1
			}
			maxS = max(maxS, h*w)
		}
		stack = append(stack, i)
	}
	fmt.Fprintln(out, maxS)
}
