// Channels are a way to communicate between goroutines.
// They allow you to send and receive values of a specified type between goroutines.
// Channels can be buffered or unbuffered,
// and they can be used to synchronize the execution of goroutines.
// FEATURES OF CHANNELS
// 1. Hold Data
// 2. Thread Safe - Avoid Data Race when reading and writing
// 3. Listen for Data - Block until data is available

package main

// LIKE THIS IT WILL THROW DEADLOCK ERROR

// func main() {
// 	var c = make(chan int) // create a channel of type int using make
// 	c <- 1
// 	var i = <-c
// 	fmt.Println(i)
// }

// import (
// 	"fmt"
// 	"time"
// )

// func main() {
// 	var c = make(chan int, 5) // create a buffered channel of type int using make
// 	go process(c)
// 	for i := range c {
// 		fmt.Println(i)
// 		time.Sleep(time.Second * 1)
// 	}
// }

// func process(c chan int) {
// 	defer close(c)
// 	for i := 0; i < 5; i++ {
// 		c <- i
// 	}
// 	fmt.Println("Existing Process")

// }

import (
	"fmt"
	"math/rand"
	"time"
)

var MAX_CHICKEN_PRICE float32 = 5

func main() {
	var chickenChannel = make(chan string)
	var websites = []string{"walmart.com", "costco.com", "wholefoods.com"}
	for i := range websites {
		go checkChickenPrices(websites[i], chickenChannel)
	}
	sendMessage(chickenChannel)
}

func checkChickenPrices(website string, chickenChannel chan string) {
	for {
		time.Sleep(time.Second * 1)
		var chickenPrice = rand.Float32() * 20
		if chickenPrice <= MAX_CHICKEN_PRICE {
			chickenChannel <- website
			break
		}
	}
}

func sendMessage(chickenChannel chan string) {
	fmt.Printf("\nFound a deal on chicken at %s", <-chickenChannel)
}
