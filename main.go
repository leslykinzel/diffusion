package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) != 2 {
		fmt.Println("Usage: diffusion <a> <b>")
		return
	}

	sublen := lcs(args[0], args[1], len(args[0]), len(args[1]))
	if sublen == len(args[0]) && sublen == len(args[1]) {
		fmt.Printf("Longest common subsequence is: %d, strings are identical.\n", sublen)
	} else {
		fmt.Printf("Longest common subsequence is: %d, strings are different.\n", sublen)
	}
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func lcs(a string, b string, m int, n int) int {
	if m == 0 || n == 0 {
		return 0
	}

	if a[m-1] == b[n-1] {
		return 1 + lcs(a, b, m-1, n-1)
	}

	return max(lcs(a, b, m, n-1), lcs(a, b, m-1, n))
}
