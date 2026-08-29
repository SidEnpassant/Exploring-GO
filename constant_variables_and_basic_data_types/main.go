package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	var intNum int = 3000
	intNum = intNum + 1000
	fmt.Println(intNum)

	var floatNum float64 = 12345678.9
	fmt.Println(floatNum)

	var floatNum32 float32 = 10.1
	var intNum32 int32 = 2
	var result float32 = floatNum32 + float32(intNum32) // cant add two different , need to typecast and perform operation
	fmt.Println(result)

	var intNum1 int = 3
	var intNum2 int = 2
	fmt.Println(intNum1 / intNum2) // no type cast needed because of same datatype

	var myString string = "Hello \n World"
	fmt.Println(myString)
	fmt.Println(len(myString))                    // length of string , not the number of characters but the number of bytes int string
	fmt.Println(utf8.RuneCountInString(myString)) // this is used to count number of characters in string

	var myRune rune = 'a'
	fmt.Println(myRune)

	var myBooleanm bool = false
	fmt.Println(myBooleanm)

	const myConst string = "this is const" // cant change after created
	fmt.Println(myConst)
	const pi float64 = 3.14
	fmt.Println(pi)
}
