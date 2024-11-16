package main

import (
	"fmt"
	"os"

	diff "github.com/leslykinzel/diffusion/pkg/difftools"
)

func main() {
	args := os.Args[1:]
	if len(args) != 2 {
		fmt.Println("Usage: diffusion <a> <b>")
		os.Exit(1)
	}

	sublen := diff.LCS(args[1], args[1], len(args[0]), len(args[1]))
	if sublen == len(args[0]) && sublen == len(args[1]) {
		fmt.Printf("Longest common subsequence is: %d, strings are identical.\n", sublen)
	} else {
		fmt.Printf("Longest common subsequence is: %d, strings are different.\n", sublen)
	}
}
