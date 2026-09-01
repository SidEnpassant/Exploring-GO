package main

import "fmt"

type gasEngine struct {
	mpg     uint8
	gallons uint8
	// ownerInfo owner
}

type owner struct {
	name string
}

func (e gasEngine) milesLeft() uint8 {
	return e.gallons * e.mpg
}

type electricEngine struct {
	kwh         uint8
	milesPerKwh uint8
	ownerInfo   owner
}

func (e electricEngine) milesLeft() uint8 {
	return e.kwh * e.milesPerKwh
}

type engine interface {
	milesLeft() uint8 // called method signature
}

func canMakeIt(e engine, miles uint8) {
	if miles < e.milesLeft() {
		fmt.Println("You can make it!")
	} else {
		fmt.Println("You cannot make it! Need to fuel up")
	}
}
func main() {
	var myEngine gasEngine = gasEngine{25, 15}
	canMakeIt(myEngine, 50)
}
