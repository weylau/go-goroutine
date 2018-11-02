package main

import (
	"fmt"
)

type worker struct {
	in   chan int
	done chan bool
}

func doWorker(workid int, c worker) {
	for n := range c.in {
		fmt.Printf("recerve %d content %d\n", workid, n)
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

func genarate() chan int {
	out := make(chan int)
	go func() {
		for {
			i := 0
			for {
				//time.Sleep(
				//	time.Duration(rand.Intn(1500)) *
				//		time.Millisecond)
				out <- i
				i++
			}
		}
	}()
	return out
}

func main() {
	//生成待执行的任务数据
	var c1, c2 = genarate(), genarate()

	//创建执行任务的worker
	toWorker := createWorker(0)

	var valuesCache []int

	for {
		var activeWorker worker
		var activeValue int
		if len(valuesCache) > 0 {
			activeWorker = toWorker
			activeValue = valuesCache[0]
		}
		select {
		case n := <-c1:
			valuesCache = append(valuesCache, n)
		case n := <-c2:
			valuesCache = append(valuesCache, n)
		case activeWorker.in <- activeValue:
			valuesCache = valuesCache[1:]
		case <-toWorker.done:
			fmt.Println("任务执行完毕")
		}
	}

	fmt.Println("end.......")
}
