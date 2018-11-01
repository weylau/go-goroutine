package main

import (
	"fmt"
)

type worker struct {
	in   chan int
	done chan bool
}

func main() {
	channelDemo()
}

func doWorker(workid int, c worker) {
	for n := range c.in {
		fmt.Printf("recerve %d content %c\n", workid, n)
		c.done <- true
	}
}

func createWorker(id int) worker {
	c := worker{
		in:   make(chan int),
		done: make(chan bool),
	}
	go doWorker(id, c)
	return c
}
func channelDemo() {
	var workers [10]worker
	for i, _ := range workers {
		workers[i] = createWorker(i)
	}

	for i := 0; i < 1; i++ {
		for j, _ := range workers {
			workers[j].in <- 'a' + j
		}
	}

	for _, worker := range workers {
		<-worker.done
	}

}
