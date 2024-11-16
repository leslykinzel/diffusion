package difftools

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func LCS(a, b string, m, n int) int {
	if m == 0 || n == 0 {
		return 0
	}

	if a[m-1] == b[n-1] {
		return 1 + LCS(a, b, m-1, n-1)
	}

	return max(LCS(a, b, m, n-1), LCS(a, b, m-1, n))
}
