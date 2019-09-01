package main

import (
	"fmt"
	"sync"
)

func main() {
	even := make(chan int)
	odd := make(chan int)
	fanin := make(chan int)

	go send(even, odd)
	go receive(even, odd, fanin)
	for v := range fanin {
		fmt.Println(v)
	}
	fmt.Println(" We are done !!!")

}

// send channel fanin taken values from multiple channel and adding in single channel
func send(even, odd chan<- int) {
	defer close(odd)
	defer close(even)
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			even <- i
		} else {
			odd <- i
		}
	}

}

//receive channel
func receive(even <-chan int, odd <-chan int, fanin chan<- int) {
	var wg sync.WaitGroup
	defer close(fanin)
	wg.Add(2)

	go func() {
		for v := range even {
			fanin <- v
		}
		wg.Done()
	}()

	go func() {
		for v := range odd {
			fanin <- v
		}
		wg.Done()
	}()

	wg.Wait()
}
