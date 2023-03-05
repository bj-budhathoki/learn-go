package main

import "context"

// context usage for controlling cancellation
// A context deadlines,cancellation signals and other required-scoped values across API boundaries and goroutine.

/**
When to use go context
-> To pass data to the downstream
-> When you want to halt the operation within a specified time from start.
-> When you want to halt an operation before a certain time.
**/

func main() {

	//create new context
	rootCtx := context.Background()

}
