package main

import (
	"errors"
	"fmt"
)

func main() {
	var printValue string = "Hello"
	var numerator int = 10
	var denominator int = 2
	var result, remainder, err = intDivision(numerator, denominator)
	// if err != nil {
	// 	fmt.Printf(err.Error())
	// } else if remainder == 0 {
	// 	fmt.Printf("The result of the integer division is %v", result)
	// } else {
	// 	fmt.Printf("The result of the integer division is %v with remainder %v", result, remainder)
	// }
	switch {
	case err != nil:
		fmt.Printf(err.Error())
	case remainder == 0:
		fmt.Printf("The result of the integer division is %v", result)
	default:
		fmt.Printf("The result of the integer division is %v with remainder %v", result, remainder)
	}
	switch remainder {
	case 0:
		fmt.Printf("The division was exact")
	case 1, 2:
		fmt.Printf("The result of the division was close")
	default:
		fmt.Printf("The division was not close")
	}
	fmt.Println(result)
	fmt.Println(remainder)
	fmt.Printf("The result of the integer division is %v with remainder %v", result, remainder)
	printMe(printValue)
}

func printMe(printValue string) {
	fmt.Println(printValue)
}

func intDivision(numerator int, denominator int) (int, int, error) {
	var err error
	if denominator == 0 {
		err = errors.New("cannot divided by zero")
		return 0, 0, err
	}
	var result int = numerator / denominator
	var remainder int = numerator % denominator

	return result, remainder, err
}
