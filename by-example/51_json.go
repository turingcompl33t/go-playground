package main

import (
	"encoding/json"
	"fmt"
)

type response1 struct {
	Page   int
	Fruits []string
}

type response2 struct {
	Page   int      `json:"page"`
	Fruits []string `json:"fruits"`
}

func main() {
	// Marshalling for basic (atomic) types
	bolB, _ := json.Marshal(true)
	fmt.Println(string(bolB))

	intB, _ := json.Marshal(3)
	fmt.Println(string(intB))

	fltB, _ := json.Marshal(3.14)
	fmt.Println(string(fltB))

	strB, _ := json.Marshal("gopher")
	fmt.Println(string(strB))

	// Marshalling for slices and maps
	slcD := []string{"apple", "peach", "pear"}
	slcB, _ := json.Marshal(slcD)
	fmt.Println(string(slcB))

	mapD := map[string]int{"apple": 3, "peach": 5, "pear": 7}
	mapB, _ := json.Marshal(mapD)
	fmt.Println(string(mapB))

	// Automatic marshalling of custom types (without annotations)
	res1D := &response1{
		Page:   1,
		Fruits: []string{"apple", "pear", "peach"},
	}
	res1B, _ := json.Marshal(res1D)
	fmt.Println(string(res1B))

	// Automatic marshalling of custom types, with annotations
	res2D := &response2{
		Page:   1,
		Fruits: []string{"apple", "pear", "peach"},
	}
	res2B, _ := json.Marshal(res2D)
	fmt.Println(string(res2B))
}
