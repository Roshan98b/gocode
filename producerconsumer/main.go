package main

import (
	"fmt"
	"strconv"
	"sync"
)

func main() {
	var channel = make(chan string)
	defer close(channel)

	var wgp sync.WaitGroup
	var wgc sync.WaitGroup
	var pmutex sync.Mutex
	var cmutex sync.Mutex

	for i := 1; i<=5; i++ {
		wgp.Add(1)
		go producer(channel, "item"+strconv.Itoa(i), &wgp, &pmutex)		
	}

	
	for i := 1; i<=5; i++ {
		wgc.Add(1)
		go consumer(channel, &wgc, &cmutex)
	}
	
	wgp.Wait()
	wgc.Wait()
	fmt.Println("Complete")
}


func producer(channel chan<- string, item string, wgp *sync.WaitGroup, mutex *sync.Mutex) {
	defer wgp.Done()
	mutex.Lock()
	channel <- item
	mutex.Unlock()
}

func consumer(channel <-chan string, wgc *sync.WaitGroup, mutex *sync.Mutex) {
	defer wgc.Done()
	mutex.Lock()
	fmt.Println("Recieved item in channel: ", <-channel)
	mutex.Unlock()
}