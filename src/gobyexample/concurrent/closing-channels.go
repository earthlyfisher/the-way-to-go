package main

import "fmt"

/*
 It repeatedly receives from jobs with j, more := <-jobs.
 In this special 2-value form of receive, the more value
 will be false if jobs has been closed and all values in
 the channel have already been received. We use this
 to notify on done when we’ve worked all our jobs.
 */
func main() {
	jobs := make(chan int, 5)
	done := make(chan bool)

	go func() {
		for {
			j, more := <-jobs
			if more {
				fmt.Println("received job", j)
			} else {
				fmt.Println("received all jobs")
				done <- true
				return
			}
		}
	}()

	for j := 1; j <= 3; j++ {
		jobs <- j
		fmt.Println("sent job", j)
	}

	//close channel,即使关闭了,channel中的元素还存在，需要等待接收
	close(jobs)
	fmt.Println("sent all jobs")

	<-done
}
