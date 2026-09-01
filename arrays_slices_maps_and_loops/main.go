// package main

// import "fmt"

// func main() {
// 	var intArr [3]int32 = [3]int32{1, 2, 3} //in second bracket we specify no of elements
// 	// intArr[0] = 123
// 	// fmt.Println(intArr[0])
// 	// fmt.Println(intArr[1:3])
// 	fmt.Println(intArr)

// 	//slices are just arrays with additional functionality
// 	var intSlice []int32 = []int32{4, 5, 6}
// 	fmt.Printf("The length of %v with capacity %v", len(intSlice), cap(intSlice)) // cap = capacity
// 	fmt.Println(intSlice)
// 	intSlice = append(intSlice, 7)
// 	fmt.Printf("The length of %v with capacity %v", len(intSlice), cap(intSlice)) // cap = capacity
// 	fmt.Println(intSlice)

// 	var intSlice2 []int32 = []int32{8, 9}
// 	intSlice = append(intSlice, intSlice2...)
// 	fmt.Println(intSlice)

// 	var intSlice3 []int32 = make([]int32, 3, 8)
// 	fmt.Println(intSlice3)

// 	//map
// 	var myMap map[string]uint8 = make(map[string]uint8)
// 	fmt.Println(myMap)

// 	var myMap2 = map[string]uint8{"Adam": 23, "Sarah": 45}
// 	fmt.Println(myMap2["Adam"])
// 	fmt.Println(myMap2["Jason"]) // This will print the zero value for the type (0 in this case)
// 	// delete(myMap2, "Adam")
// 	var age, ok = myMap2["Jason"]
// 	if ok {
// 		fmt.Println("The age is %v", age)
// 	} else {
// 		fmt.Println("Invalid Name")
// 	}

// 	for name, age := range myMap2 {
// 		fmt.Println("Name: %v , Age: %v \n", name, age)
// 	}
// 	for i, v := range intArr {
// 		fmt.Printf("Index: %v , value: %v \n", i, v)
// 	}

// 	for i := 0; i < 10; i++ {
// 		fmt.Println(i)
// 	}
// }
