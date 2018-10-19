package main

import "fmt"

// Define person struct
type Persion struct {
	firstName string
	lastName string
	city string
	gender string
	age int
}

func main() {
	// Init person using struct
	persion1 := Persion{lastName:"Lee", city:"Wuhan", gender:"f", age:20}
	fmt.Println(persion1)

	// Alternative
	persion2 := Persion{"cuxx", "lee", "wuhan", "f", 20}
	fmt.Println(persion2)
	fmt.Println(persion2.firstName)

}


