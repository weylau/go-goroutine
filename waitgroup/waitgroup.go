package main

import (
	"fmt"
	"sync"
)

type worker struct {
	in   chan int
	done func()
}

func main() {
	channelDemo()
}

func doWorker(workid int, c worker) {
	for n := range c.in {
		fmt.Printf("recerve %d content %c\n", workid, n)
		c.done()
	}
}

func createWorker(id int, wg *sync.WaitGroup) worker {
	c := worker{
		in: make(chan int),
		done: func() {
			wg.Done()
		},
	}
	go doWorker(id, c)
	return c
}
func channelDemo() {
	var workers [10]worker
	var wg sync.WaitGroup

	for i, _ := range workers {
		workers[i] = createWorker(i, &wg)
	}

	for i := 0; i < 1; i++ {
		for j, _ := range workers {
			wg.Add(1)
			workers[j].in <- 'a' + j
		}
	}
	wg.Wait()

}
