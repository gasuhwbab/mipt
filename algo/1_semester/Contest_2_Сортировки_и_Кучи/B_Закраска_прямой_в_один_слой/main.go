package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type point struct {
	x      int
	is_end int
}

func main() {
	var in *bufio.Reader
	var out *bufio.Writer
	in = bufio.NewReader(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var N int
	fmt.Fscan(in, &N)
	points := []point{}
	for ; N > 0; N-- {
		var l, r int
		fmt.Fscan(in, &l, &r)
		points = append(points, point{x: l, is_end: 1})
		points = append(points, point{x: r, is_end: -1})
	}
	sort.Slice(points, func(i, j int) bool {
		return points[i].x < points[j].x || (points[i].x == points[j].x && points[i].is_end > points[j].is_end)
	})
	cnt := 0
	len := 0
	prev := 0
	for _, p := range points {
		if cnt == 1 && prev != 0 {
			len += p.x - prev
		}
		cnt += p.is_end
		prev = p.x
	}
	fmt.Fprintln(out, len)
}
