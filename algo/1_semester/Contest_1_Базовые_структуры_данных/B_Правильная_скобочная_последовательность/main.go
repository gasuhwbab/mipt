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

	var s string
	fmt.Fscan(in, &s)
	stack := make([]rune, 0)
	pairs := map[rune]rune{')': '(', ']': '[', '}': '{'}
	for _, r := range s {
		if r == '[' || r == '(' || r == '{' {
			stack = append(stack, r)
		} else {
			if len(stack) == 0 || pairs[r] != stack[len(stack)-1] {
				fmt.Fprintln(out, "no")
				return
			}
			stack = stack[:len(stack)-1]
		}
	}
	if len(stack) == 0 {
		fmt.Fprintln(out, "yes")
	} else {
		fmt.Fprintln(out, "no")
	}
}
