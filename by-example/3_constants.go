package main

import (
	"fmt"
	"math"
)

func main() {
	const name string = "katy"
	fmt.Println(name)

	const n = 500
	fmt.Println(int64(n))

	const m = 500.0
	fmt.Println(float64(m))

	const x = 3.14
	fmt.Println(math.Sin((x)))
}
