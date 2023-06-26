package main

import "fmt"

func main() {
	// Most basic loop structure (a single condition)
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i += 1
	}

	// A classic C-style loop
	for j := 1; j <= 3; j++ {
		fmt.Println(j)
	}

	// Infinite loop structure
	k := 1
	for {
		if k > 3 {
			break
		}
		fmt.Println(k)
		k += 1
	}
}
