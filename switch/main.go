package main

import (
	"fmt"
)

func main() {

	//simple switch statement
	// i := 3
	// switch i {
	// case 1:
	// 	fmt.Println("One")
	// case 2:
	// 	fmt.Println("Two")
	// case 3:
	// 	fmt.Println("Three")
	// case 4:
	// 	fmt.Println("Four")
	// default:
	// 	fmt.Println("other")
	// }

	//multiple condition switch
	// switch time.Now().Weekday() {
	// case time.Saturday, time.Sunday:
	// 	fmt.Println("It is weekend")
	// default:
	// 	fmt.Println("Its workday")
	// }

	//type switch
	whoAmI := func(i interface{}) {
		switch t := i.(type) {
		case int:
			fmt.Println("Its an integer")
		case string:
			fmt.Println("Its a string")
		case bool:
			fmt.Println("Its a boolean")
		default:
			fmt.Println("Its something else", t)
		}
	}

	whoAmI("goLang")
}
