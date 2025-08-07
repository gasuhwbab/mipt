package first

func TaskOne(a, b []int, n, m int) []int {
	ans := make([]int, n+m)
	for i, j := 0, 0; i < n || j < m; {
		if i == n {
			ans[i+j] = b[j]
			j++
			continue
		}
		if j == m || a[i] <= b[j] {
			ans[i+j] = a[i]
			i++
		} else {
			ans[i+j] = b[j]
			j++
		}
	}
	return ans
}

func TaskTwo(a, b []int, n, m int) bool {
	i, j := 0, 0
	for i < n && j < m {
		if a[i] == b[j] {
			i++
		}
		j++
	}
	return i == n
}

func TaskThree(a []int, l, r, x int) {
	if r > len(a) {
		return
	}
	for i := l; i <= r; i++ {
		a[i] += x
	}
}

func TaskFour(a []int, k int) int {
	n := len(a)
	maxLength := 0
	left := 0
	currSum := 0
	for right := range n {
		currSum += a[right]
		for currSum > k && left < right {
			currSum -= a[left]
			left++
		}
		if currSum <= k {
			maxLength = max(maxLength, right-left+1)
		}
	}
	return maxLength
}

func TaskFive(a []int, m int) int {
	l, r := a[0], a[len(a)-1]
	for r-l > 1 {
		mid := (l + r) / 2
		currSum := 0
		for _, x := range a {
			currSum += x / mid
			if x%mid != 0 {
				currSum += 1
			}
		}
		if currSum > m {
			l = mid
		} else {
			r = mid
		}
	}
	return r
}
