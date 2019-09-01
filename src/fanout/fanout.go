package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	c1 := make(chan int)
	c2 := make(chan int)

	go populate(c1)
	go fanOut(c1, c2)

	for v := range c2 {
		fmt.Println(v)
	}
	fmt.Println("Fan out Done !!!")
}

func populate(c chan int) {
	defer close(c)
	for i := 0; i < 100; i++ {
		c <- i
	}

}

func fanOut(c1, c2 chan int) {
	var wg sync.WaitGroup
	for v := range c1 {
		wg.Add(1)
		go func(v2 int) {
			c2 <- timeConsumingWork(v2)
			wg.Done()
		}(v)
	}
	wg.Wait()
	close(c2)

}

func timeConsumingWork(n int) int {
	time.Sleep(time.Millisecond * 1000)
	return n + rand.Intn(1000)

}
