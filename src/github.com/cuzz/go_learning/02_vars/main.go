package main

import (
	"fmt"
)

func main() {
	// MAIN TYPES
	// string
	// bool
	// int
	// int int8 int16 int32 int64
	// uint uint8 uint16 uint32 uint64 uintptr
	// byte - alias for uint8
	// rune - alias for int32
	// float32 float64
	// complex64 complex128

	// test01()
	test02()

}

func test02()  {
	// Shorthand
	name := "cuzz"
	size := 1.1
	var size2 float32 = 2.2
	fmt.Println(name, size)
	fmt.Printf("Type of name is a %T\n", name)
	fmt.Printf("Type of size is a %T\n", size)
	fmt.Printf("Type of size2 is a %T\n", size2)
}


func test01() {
	// Using var
	var name = "cuzz"	     // var name string = "cuzz"
	var age int32 = 20		 // var age int = 20
	var isCool = true

	sex, email := "male", "cuzz1234@163.com"

	fmt.Println(name, age, isCool, sex, email)
	fmt.Printf("Type of name is a %T, Type of age is a %T", name, age)
}



