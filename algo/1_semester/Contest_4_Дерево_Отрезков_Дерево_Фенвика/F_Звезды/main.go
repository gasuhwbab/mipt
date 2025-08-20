package main

import (
	"bufio"
	"fmt"
	"os"
)

type BIT struct {
	t [][][]int
}

func NewBIT(n int) *BIT {
	t := make([][][]int, n)
	for i := range t {
		t[i] = make([][]int, n)
		for j := range t[i] {
			t[i][j] = make([]int, n)
		}
	}
	return &BIT{t: t}
}

func (b *BIT) add(x, y, z, val int) {
	for xi := x; xi < len(b.t); xi = xi | (xi + 1) {
		for yi := y; yi < len(b.t[xi]); yi = yi | (yi + 1) {
			for zi := z; zi < len(b.t[xi][yi]); zi = zi | (zi + 1) {
				b.t[xi][yi][zi] += val
			}
		}
	}
}

func (b *BIT) sum(x, y, z int) int {
	s := 0
	for xi := x; xi >= 0; xi = (xi & (xi + 1)) - 1 {
		for yi := y; yi >= 0; yi = (yi & (yi + 1)) - 1 {
			for zi := z; zi >= 0; zi = (zi & (zi + 1)) - 1 {
				s += b.t[xi][yi][zi]
			}
		}
	}
	return s
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n int
	fmt.Fscan(in, &n)
	bit := NewBIT(n)
	var q int
	fmt.Fscan(in, &q)
	for q != 3 {
		if q == 1 {
			var x, y, z, k int
			fmt.Fscan(in, &x, &y, &z, &k)
			bit.add(x, y, z, k)
		} else if q == 2 {
			var x1, y1, z1, x2, y2, z2 int
			fmt.Fscan(in, &x1, &y1, &z1, &x2, &y2, &z2)
			ans := bit.sum(x2, y2, z2) - bit.sum(x1-1, y2, z2) - bit.sum(x2, y1-1, z2) - bit.sum(x2, y2, z1-1) +
				+bit.sum(x2, y1-1, z1-1) + bit.sum(x1-1, y2, z1-1) + bit.sum(x1-1, y1-1, z2) - bit.sum(x1-1, y1-1, z1-1)
			fmt.Fprintln(out, ans)
		}
		q = 0
		fmt.Fscan(in, &q)
	}
}
