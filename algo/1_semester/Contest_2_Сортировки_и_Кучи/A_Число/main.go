package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {
	var in *bufio.Reader
	var out *bufio.Writer
	in = bufio.NewReader(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
	defer out.Flush()

	nums := []string{}
	var num string
	fmt.Fscan(in, &num)
	for num != "" {
		nums = append(nums, num)
		num = ""
		fmt.Fscan(in, &num)
	}
	sort.Slice(nums, func(i, j int) bool {
		return nums[i]+nums[j] > nums[j]+nums[i]
	})
	fmt.Fprintln(out, strings.Join(nums, ""))
}
