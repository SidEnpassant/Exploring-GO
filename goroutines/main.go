// Goroutines are the way to launch multiple functions concurrently in Go.
// They are lightweight threads managed by the Go runtime.
// Concurrency is not same as parallelism

package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// var mutex = sync.Mutex{} // mutex is used to lock the results slice so that only one goroutine can access it at a time
var readWriteMutex = sync.RWMutex{} // readWriteMutex is used to lock the results slice so that only one goroutine can access it at a time
var waitGroup = sync.WaitGroup{}
var dbData = []string{"id1", "id2", "id3", "id4", "id5"}
var results = []string{}

func main() {
	t0 := time.Now()
	for i := 0; i < len(dbData); i++ {
		waitGroup.Add(1)
		go dbCall(i) // used go keyword to launch a goroutine , run the function concurrently
	}
	waitGroup.Wait() // wait for all the goroutines to finish before continuing
	fmt.Printf("\nTotal execution time: %v", time.Since(t0))
	fmt.Printf("\nResults: %v", results)
}

func dbCall(i int) {
	var delay float32 = rand.Float32() * 2000
	time.Sleep(time.Duration(delay) * time.Millisecond)
	fmt.Println("The result from the database is:", dbData[i])
	//mutex.Lock() // lock the results slice so that only one goroutine can access it at a time
	save(dbData[i])
	log()
	//mutex.Unlock() // unlock the results slice so that other goroutines can access it
	waitGroup.Done()
}

func save(result string) {
	readWriteMutex.Lock() // lock the results slice
	results = append(results, result)
	readWriteMutex.Unlock() // unlock the results slice so that other goroutines can access it

}
func log() {
	readWriteMutex.RLock()
	fmt.Println("The results are: ", results)
	readWriteMutex.RUnlock()
}
