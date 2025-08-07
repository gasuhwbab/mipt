package third

import "sort"

type Pair struct {
	a, b int
}

// 0 1 2 3 4 5

func TaskOne(pairs []Pair, k int) int {
	sum := 0
	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].a-pairs[i].b >= pairs[j].a-pairs[j].b
	})
	for i := range pairs {
		sum += pairs[i].b
		if k > 0 {
			sum += pairs[i].a - pairs[i].b
			k--
		}
	}
	return sum
}
