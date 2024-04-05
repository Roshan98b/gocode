package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup
var rlock sync.Mutex
var wlock sync.Mutex

func main() {
	val := 0

	for i := 1; i <= 1000; i++ {
		wg.Add(2)
		rlock.Lock()
		go write(&val)
		wlock.Lock()
		go increment(&val)
	}
	wg.Wait()
	fmt.Println("This is main: ", val)
}

func write(p *int) {
	defer wg.Done()
	fmt.Printf("Value : %v\n", *p)
	rlock.Unlock()
}

func increment(p *int) {
	defer wg.Done()
	*p = *p + 1
	wlock.Unlock()
}
