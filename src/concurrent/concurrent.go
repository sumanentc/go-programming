package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

func main() {
	testMutex()
	testAtomic()
}

func testAtomic() {
	var wg sync.WaitGroup
	fmt.Println(runtime.NumCPU)
	fmt.Println(runtime.NumGoroutine)

	var counter int64
	gs := 10
	wg.Add(gs)
	for i := 0; i < gs; i++ {
		go func() {
			atomic.AddInt64(&counter, 1)
			fmt.Println(atomic.LoadInt64(&counter))
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println("End Value Atomic !!! ", counter)

}

//concurrent lock example using mutex
func testMutex() {

	var wg sync.WaitGroup
	fmt.Println(runtime.NumCPU)
	fmt.Println(runtime.NumGoroutine)

	counter := 0
	gs := 10
	wg.Add(gs)
	var m sync.Mutex
	for i := 0; i < gs; i++ {
		go func() {
			m.Lock()
			v := counter
			runtime.Gosched()
			v++
			counter = v
			fmt.Println(counter)
			m.Unlock()
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println("End Value Mutex !!! ", counter)

}
