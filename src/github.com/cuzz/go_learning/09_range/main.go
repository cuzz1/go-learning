package main

import "fmt"

func main() {
	ids := []int{22, 32, 46, 31, 04,15}

	// Loop through ids
	for i , id := range ids{
		fmt.Printf("%d - ID: %d\n", i, id)
	}

	// Not using index
	for _, id := range ids{
		fmt.Printf("ID: %d\n", id)
	}

	// Add ids together
	sum := 0
	for _, id := range ids {
		sum += id
	}

	fmt.Printf("Sum: %d\n", sum)

	// Range with map
	emails := map[string]string{"cuzz":"cuzz1234@163.com", "cuxx":"123456@qq.com"}

	for k, v := range emails {
		fmt.Printf("%s: %s\n", k, v)
	}


}


