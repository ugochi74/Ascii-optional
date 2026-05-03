package main

//import "fmt"

import (
	"fmt"
)

func PrintCombN(n int) {
	if n <= 0 || n >= 10 {
		return
	}

	// Create a slice to act as our "dynamic loops"
	comb := make([]int, n)
	for i := 0; i < n; i++ {
		comb[i] = i
	}

	for {
		// 1. Print the current combination
		for i := 0; i < n; i++ {
			fmt.Printf("%c", comb[i]+'0')
		}

		// 2. Check if we just printed the absolute last combination (e.g., 789)
		if comb[0] == 10-n {
			fmt.Println()
			break
		}

		// 3. Print the separator
		fmt.Print(", ")
		//fmt.Print(' ')

		// 4. Find the rightmost digit that can still be incremented
		i := n - 1
		for i >= 0 && comb[i] == 9-(n-1-i) {
			i--
		}

		// 5. Increment that digit and reset everything to its right
		comb[i]++
		for j := i + 1; j < n; j++ {
			comb[j] = comb[j-1] + 1
		}
	}
}

func main() {
	PrintCombN(1)
	PrintCombN(3)
	PrintCombN(9)
	fmt.Println()
}
