package second

import (
	"sort"
)

type Point struct {
	pivot  int
	is_end int
}

type Segment struct {
	l, r int
}

func TaskOne(arr []int) int {
	ans := 0
	merge := func(a, b []int) []int {
		to := make([]int, len(a)+len(b))
		for i, j := 0, 0; i < len(a) || j < len(b); {
			if j == len(b) || (i < len(a) && a[i] <= b[j]) {
				to[i+j] = a[i]
				i++
				ans += j
			} else {
				to[i+j] = b[j]
				j++
			}
		}
		return to
	}
	var mergeSort func(a []int) []int
	mergeSort = func(a []int) []int {
		if len(a) <= 1 {
			return a
		}
		k := len(a) / 2
		l := mergeSort(a[:k])
		r := mergeSort(a[k:])
		return merge(l, r)
	}
	mergeSort(arr)
	return ans
}

func TaskTwoA(points []Point) int {
	sort.Slice(points, func(i, j int) bool {
		return points[i].pivot < points[j].pivot || (points[i].pivot == points[j].pivot &&
			points[i].is_end > points[j].is_end)
	})
	sum := 0
	cnt := 0
	prev := -1
	for _, point := range points {
		if prev != -1 && cnt > 0 {
			sum += point.pivot - prev
		}
		cnt += point.is_end
		prev = point.pivot
	}
	return sum
}

func TaskTwoB(points []Point) int {
	minb := 1000
	maxa := -1
	for _, p := range points {
		if p.is_end == 1 {
			maxa = max(maxa, p.pivot)
		} else {
			minb = min(minb, p.pivot)
		}
	}
	return minb - maxa
}

func TaskThreeC(segments []Segment) int {
	sort.Slice(segments, func(i, j int) bool {
		return segments[i].r <= segments[j].r ||
			(segments[i].r == segments[j].r && segments[i].l < segments[j].l)
	})
	last, cnt := 0, 0
	for _, s := range segments {
		if s.l > last {
			last = s.r
			cnt++
		}
	}
	return cnt
}

func TaskFour(a []int, x int) {
	pivot := a[x]
	i := 0
	j := len(a) - 1
	for i < j {
		for a[i] <= pivot {
			i++
		}
		for a[j] > pivot {
			j--
		}
		a[i], a[j] = a[j], a[i]
	}
}

func TaskFive(a []int, b []int) []int {
	ans := make([]int, 0)
	for _, x := range a {
		l, r := 0, len(b)
		for r-l > 1 {
			mid := (r + l) / 2
			if x < b[mid] {
				r = mid
			} else {
				l = mid
			}
		}
		ans = append(ans, l)
	}
	return ans
}
