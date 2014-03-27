package main

import (
	"fmt"
	"time"
)

func sum(n int) (out chan int) {
	out = make(chan int)
	go func() {
		answer := 0
		for i := n; i < 1000; i += n {
			answer += i
		}
		out <- answer
	}()
	return
}

func problem1() {
	x, y, z := sum(3), sum(5), sum(15)
	fmt.Println(<-x + <-y - <-z)
}

func main() {
	start := time.Now()
	problem1()
	fmt.Println(time.Now().Sub(start).Seconds())
}
