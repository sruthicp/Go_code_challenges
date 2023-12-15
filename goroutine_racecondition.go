/*
Assume there are M goroutines reading from a shared buffer (such as a byte slice) and N goroutines
writing into it. How to avoid deadlock and race condition? The goroutines are always running and there
is no wait group. Using wait groups is not allowed.
Solve for M = 8 and N = 2
Solve for M = 8 and N = 8
Solve for M = 8 and N = 16
Solve for M = 2 and N = 8
*/

package main

import (
	"fmt"
	"sync"
	"time"
)

const (
	N = 8 // change N here
	M = 8 // change M here
)

// declaring as global instead of passing to each functions
var (
	buffer    []string
	readIndex int
	mutex     sync.Mutex
)

func writer(id int) {
	mutex.Lock()
	message := fmt.Sprintf("Hi %d", id)
	buffer = append(buffer, message)

	fmt.Printf("Writer %d wrote: %s\n", id, message)
	mutex.Unlock()
}

func reader(id int) {
	mutex.Lock()
	message := buffer[readIndex]
	readIndex = (readIndex + 1) % len(buffer)
	fmt.Printf("Reader %d read: %s\n", id, message)
	mutex.Unlock()
}

func main() {
	for i := 0; i < N; i++ {
		go writer(i + 1)
	}
	for i := 0; i < M; i++ {
		go reader(i + 1)
	}
	time.Sleep(3 * time.Second) // waiting for sometime
	fmt.Println("Done!")
}
