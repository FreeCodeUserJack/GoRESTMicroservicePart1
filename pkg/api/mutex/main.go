package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var (
	counter int64
	mutex   sync.Mutex
	atomicValue AtomicInt
)

type AtomicInt struct {
	value int64
	lock sync.Mutex
}

func (a *AtomicInt) Increase() {
	a.lock.Lock()
	defer a.lock.Unlock()
	a.value++
}

func main() {

	var wg sync.WaitGroup

	for i := 0; i < 10000; i++ {
		wg.Add(1)
		go func() {
			mutex.Lock()
			counter++
			mutex.Unlock()
			wg.Done()

			atomicValue.Increase()
		}()
	}

	for i := 0; i < 10000; i++ {
		wg.Add(1)
		go func() {
			atomic.AddInt64(&counter, 1)

			atomicValue.Increase()
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println(counter)
	fmt.Println(atomicValue.value)
}