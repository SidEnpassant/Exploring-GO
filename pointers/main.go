package main

import "fmt"

// func main() {
// 	var p *int32 = new(int32)
// 	var i int32
// 	fmt.Printf("The value p points to is: %v", *p) // * = dereferencing the pointer
// 	fmt.Printf("\nThe address of i is: %v", i)

// 	*p = 10 // this means set the value of the variable that p points to to 10
// 	fmt.Printf("\nThe value p points to is now: %v", *p)
// 	p = &i // this means set p to point to the address of i
// 	*p = 1 // this means set the value of the variable that p points to (which is i) to 1

// 	var slice = []int32{1, 2, 3}
// 	var sliceCopy = slice
// 	sliceCopy[2] = 4
// 	fmt.Println(slice)
// 	fmt.Println(sliceCopy)
// }

func main() {
	var thing1 = [5]float64{1, 2, 3, 4, 5}
	fmt.Printf("\nThe memory location of the thing1 array is: %p", &thing1)
	var result [5]float64 = square(thing1)
	fmt.Println("\nThe result is: %v", result)
}

func square(thing2 [5]float64) [5]float64 {
	fmt.Printf("\nThe memory location of the thing2 array is: %p", &thing2)
	for i := range thing2 {
		thing2[i] = thing2[i] * thing2[i]
	}
	return thing2
}
