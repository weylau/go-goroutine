package main

import (
	"fmt"
	"time"
)

func main() {
	channelDemo()
}

func worker(workid int, c chan int) {
	for n := range c {
		fmt.Printf("recerve %d content %c\n", workid, n)
	}
}

func createWorker(id int) chan int {
	c := make(chan int)
	go worker(id, c)
	return c
}
func channelDemo() {
	var channels [10]chan int
	for i := 0; i < 10; i++ {
		channels[i] = createWorker(i)
	}

	for i := 0; i < 10; i++ {
		channels[i] <- 'a' + i
	}

	time.Sleep(time.Microsecond)
}
