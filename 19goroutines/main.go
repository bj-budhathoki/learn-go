package main

import (
	"fmt"
	"time"
)

// Go-routine is a lightweight thread managed by go runtime.
// -> has a separate independent execution
// -> goroutine helps achieve concurrency in golang
// usage special keyword 'go' for starting a goroutine.

func main() {
	go start()
	fmt.Println("Started")
	time.Sleep(1 * time.Second)
	fmt.Println("Finished")
}

func start() {
	fmt.Println("In Goroutine")
}
