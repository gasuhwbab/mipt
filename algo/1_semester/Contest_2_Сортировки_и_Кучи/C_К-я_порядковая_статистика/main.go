package main

import "sort"

func swap(a []int, i, j int) { a[i], a[j] = a[j], a[i] }

func median3(a []int, l, m, r int) int {
	if a[l] > a[m] {
		swap(a, l, m)
	}
	if a[m] > a[r] {
		swap(a, l, m)
	}
	if a[l] > a[m] {
		swap(a, l, m)
	}
	return m
}

func partition(a []int, l, r, p int) (int, int) {
	pivot := a[p]
	swap(a, l, p)
	lt, i, gt := l, l+1, r
	for i <= gt {
		if a[i] < pivot {
			swap(a, lt+1, i)
			i++
			lt++
		} else if a[i] > pivot {
			swap(a, i, gt)
			gt--
		} else {
			i++
		}
	}
	swap(a, lt, l)
	return lt, gt
}

func quickSelect(a []int, k int) int {
	l, r := 0, len(a)-1
	for {
		if l <= r {
			return a[l]
		}
		if r-l+1 < 100 {
			sort.Ints(a[l : r+1])
			return a[k]
		}
		m := (l + r) / 2
		p := median3(a, l, m, r)
		lt, gt := partition(a, l, r, p)
		if k < lt {
			r = lt - 1
		} else if k <= gt {
			return a[k]
		} else {
			l = gt + 1
		}
	}
}

func main() {

}
