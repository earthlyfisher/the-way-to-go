package main

import (
	"time"
	"fmt"
)

/*
   timer:定时器；
   ticker:重复器(以固定的间隔重复某事)
 */

func main() {
	timer1 := time.NewTimer(time.Second * 2)

	<-timer1.C
	fmt.Println("Timer 1 expired")

	timer2 := time.NewTimer(time.Second)
	go func() {
		<-timer2.C
		fmt.Println("Timer 2 expired")
	}()
	stop2 := timer2.Stop()
	if stop2 {
		fmt.Println("Timer 2 stopped")
	}

	tickers()
}

func tickers() {
	ticker1 := time.NewTicker(time.Second * 2)

	go func() {
		for t := range ticker1.C {
			fmt.Println("Tick at", t)
		}
	}()

	time.Sleep(time.Second * 10)
	ticker1.Stop()
	fmt.Println("Ticker stopped")
}
