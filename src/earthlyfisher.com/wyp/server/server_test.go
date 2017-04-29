package server

import (
	"../server"
	"fmt"
	"time"
)

func Example() {
	s := server.New()

	divideByZero := &server.Work{
		Op:    func(a, b int) int { return a / b },
		A:     100,
		B:     0,
		Reply: make(chan int),
	}
	s <- divideByZero

	select {
	case res := <-divideByZero.Reply:
		fmt.Println(res)
	case <-time.After(time.Second):
		fmt.Println("No result in one second.")
	}
	// Output: No result in one second.
}
