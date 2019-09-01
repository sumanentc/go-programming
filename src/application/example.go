package main

import (
	"fmt"
	"sync"
)

func main() {
	oddEven()
}

//Priniting odd even numbers using multiple go routines
func oddEven() {
	var wg sync.WaitGroup
	wg.Add(2)
	num := 10
	oddChannel := make(chan bool)
	evenChannel := make(chan bool)
	fmt.Println("Length and Capacity of Odd Channels !!! ", len(oddChannel), " ", cap(oddChannel))
	fmt.Println("Length and Capacity of Even Channels !!! ", len(evenChannel), " ", cap(evenChannel))
	go printOddNumbers(num, oddChannel, evenChannel, &wg)
	oddChannel <- true
	go printEvenNumbers(num, oddChannel, evenChannel, &wg)
	fmt.Println(" Waiting to be Done !!!! ")
	wg.Wait()
	fmt.Println(" It's Done !!!! ")
}

func printOddNumbers(num int, odd chan bool, even chan bool, wg *sync.WaitGroup) {
	i := 1
	defer close(odd)
	for i <= num {
		//fmt.Println("Inside odd ", i)
		if i%2 != 0 {
			b := <-odd
			if b {
				fmt.Println("Printing odd ", i)
				even <- true
			}
		}
		i++
	}
	fmt.Println("Odd Numbers are Done !!!")
	wg.Done()
}

func printEvenNumbers(num int, odd, even chan bool, wg *sync.WaitGroup) {
	i := 2
	defer close(even)
	for i <= num {
		//fmt.Println("Inside even ", i)
		if i%2 == 0 {
			b := <-even
			if b {
				fmt.Println("Printing even ", i)
				if i < num {
					odd <- true
				}

			}
		}
		i++

	}
	fmt.Println("Even Numbers are Done !!!")
	wg.Done()
}
