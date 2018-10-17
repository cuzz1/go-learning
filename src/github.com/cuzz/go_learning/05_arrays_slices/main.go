package main

import "fmt"

func main() {
	//test01()
	//test02()
	test03()

}
func test03() {
	fruitArr := []string{"Apple", "Orange", "Grape"}
	fmt.Println(len(fruitArr))
	fmt.Println(fruitArr[0:len(fruitArr)])  // don‘t include the last

}

func test02() {
	// Declare and assign
	fruitArr := [2]string{"Apple", "Orange"}
	fmt.Println(fruitArr[0])
	fmt.Println(fruitArr)
}


func test01() {
	// Arrays
	var fruitArr [2]string
	// Assigin values
	fruitArr[0] = "Apple"
	fruitArr[1] = "Orange"
	fmt.Println(fruitArr[0])
	fmt.Println(fruitArr)
}

