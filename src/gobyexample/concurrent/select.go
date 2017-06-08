package main

import (
	"fmt"
)

func main() {
	stopArray := make([]bool, 2)

	c1 := make(chan int, 3)
	c2 := make(chan int, 2)
	go func() {
		stopArray[0] = false
		for i := 0; i < 10; i++ {
			c1 <- i
		}
		stopArray[0] = true
	}()

	go func() {
		stopArray[1] = false
		for i := 0; i < 5; i++ {
			c2 <- i
		}
		stopArray[1] = true
	}()

L:
	for {
		select {
		case c1_element := <-c1:
			fmt.Println("c1", " ", c1_element)

		case c2_element := <-c2:
			fmt.Println("c2", " ", c2_element)

		default:
			if stopFunc(stopArray) {
				break L
			}
		}
	}

	fmt.Println("end")
}

func stopFunc(stopArray []bool) bool {
	stop := true
	for _, value := range stopArray {
		stop = stop && value
	}
	return stop
}
