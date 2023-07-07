package main

import (
	"context"
	"fmt"
)

type userKey int

const key userKey = 1

func main() {
	// Bodner uses this as an example in the book; defining a constant value
	// to serve as the key into the context for implicit value storage; the
	// part that seems to be missing is that the context is not shared across
	// requests, this should be obvious based on the manner in which the context
	// is used, but as I was reading this part in the book I was deeply confused
	// about the use of a constant value as a key that must support multiple values.

	ctx := context.Background()

	ctx = context.WithValue(ctx, key, "hello")
	ctx = context.WithValue(ctx, key, "world")

	value, ok := ctx.Value(key).(string)
	if !ok {
		panic("help")
	}
	fmt.Println(value)
}
