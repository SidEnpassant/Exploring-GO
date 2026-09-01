package main

import "fmt"

func main() {
	var myString = "Resume"
	var indexed = myString[0]
	fmt.Println(myString)
	fmt.Printf("%v , %T \n", indexed, indexed)
	for i, v := range myString {
		fmt.Println(i, v)
	}
	fmt.Printf("\n The length of 'myString' is %v", len(myString))
	var myRune = 'a'
	fmt.Printf("\nmyRune = %v", myRune)

	var strSlice = []string{"s", "u", "b", "s", "c", "r", "i", "b", "e"}
	var catStr = ""
	for i := range strSlice {
		catStr += strSlice[i]
	}
	fmt.Printf("\n%v", catStr)
}
