package difftools

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func LCS(a, b string) int {
	m := len(a)
	n := len(b)

	if m == 0 || n == 0 {
		return 0
	}

	if a[m-1] == b[n-1] {
		return 1 + LCS(a[:m-1], b[:n-1])
	}

	return max(LCS(a, b[:n-1]), LCS(a[:m-1], b))
}
