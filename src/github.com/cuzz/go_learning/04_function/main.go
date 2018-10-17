package main

import (
	"fmt"
)

func greeting(name string) string {
	return "Hello " + name
}

func getSum(a int, b int) int {
	return a + b
}

func main() {
	fmt.Println(greeting("cuzz"))
	fmt.Println(getSum(2, 4))
}


