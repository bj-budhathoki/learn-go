package main

import (
	"fmt"
	"time"
)

// select statement lets a goroutine wait on multiple communications operation
// A select block until one of its case can run,then it executes that case.
// Select statement is just like switch statement,but in the select statement,case statement refer to communication.
func main() {
	s1 := make(chan string)
	s2 := make(chan string)
	go func() {
		time.Sleep(1 * time.Second)
		s1 <- "one"
	}()
	go func() {
		time.Sleep(2 * time.Second)
		s2 <- "two"
	}()
	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-s1:
			fmt.Println("received", msg1)
		case msg2 := <-s2:
			fmt.Println("received", msg2)
		}
	}
}
