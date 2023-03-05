package main

import (
	"context"
	"fmt"
)

// context usage for controlling cancellation
// A context deadlines,cancellation signals and other required-scoped values across API boundaries and goroutine.

/*
*
When to use go context
-> To pass data to the downstream
-> When you want to halt the operation within a specified time from start.
-> When you want to halt an operation before a certain time.
*
*/
func enrichContext(ctx context.Context) context.Context {
	return context.WithValue(ctx, "request_id", "45756757")
}
func printSomething(ctx context.Context) {
	reqId := ctx.Value("request_id")
	fmt.Println(reqId)
}
func main() {

	//create new context
	rootCtx := context.Background()
	//context package function Background() return a empty context which implements the context interface
	// -> it has no values
	// -> it is never canceled
	// -> it has no deadline
	// -> context.Background() servers as the root of all context which will be derived from it.
	ctxEnrich := enrichContext(rootCtx)
	printSomething(ctxEnrich)
}
