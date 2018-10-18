package main

import "fmt"

func main() {
	// Define map
	//emails := make(map[string]string)

	// Assign kv
	//emails["cuzz"] = "cuzz1234@163.com"
	//emails["cuxx"] = "123456@qq.com"

	// Declare mao and add kv
	emails := map[string]string{"cuzz":"cuzz1234@163.com", "cuxx":"123456@qq.com"}

	fmt.Println(emails)
	fmt.Println(len(emails))
	fmt.Println(emails["cuzz"])

	// Delete from map
	delete(emails, "cuxx")



}


