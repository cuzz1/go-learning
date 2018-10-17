package main

import "fmt"

func main() {

	//test01()
	//test02()
	test03()

}

func test03()  {
	color := "green"
	switch color {
	case "red":
		fmt.Println("color is red")
	case "blue":
		fmt.Println("color is blue")
	default:
		fmt.Println("color is NOT blue or red")
	}
}

func test02() {
	// else if
	color := "red"
	if color == "red" {
		fmt.Println("color is red")
	} else if color == "blue"{
		fmt.Println("color is blue")
	} else {
		fmt.Println("color is NOT blue or red")
	}

}

func test01() {

	x := 5
	y := 10

	// If else
	if x <= y {
		fmt.Printf("%d is less than or equal to %d\n", x, y)
	} else {
		fmt.Printf("%d is less than %d\n", y, x)
	}
}

