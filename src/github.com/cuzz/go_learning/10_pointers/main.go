package main

import "fmt"

func main() {
	a := 5
	b := &a
	fmt.Println(a, b)                 // 5 0xc00004a080
	fmt.Printf("%T %T\n", a, b) // int *int

	// Use * to read val from address
	fmt.Println(*b)     // 5
	fmt.Println(*&a) // 5

	// Change val with pointer
	*b = 10
	fmt.Println(a)    // 10
}


