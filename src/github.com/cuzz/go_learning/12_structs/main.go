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

// Greeting method (value receiver)
func (p Persion) greet() string {
	return "Hello, my name is " + p.firstName + " " + p.lastName
}

// hasBirthday method (pointer receiver)
func (p *Persion) hasBirthday()  {
	p.age++
}

// getMarried (pointer receiver)
func (p *Persion) getMarried(spouseLastName string)  {
	if p.gender == "m" {
		return
	} else {
		p.lastName = spouseLastName
	}

}

func main() {
	// Init person using struct
	persion1 := Persion{lastName:"Lee", city:"Wuhan", gender:"f", age:20}
	fmt.Println(persion1)

	// Alternative
	persion2 := Persion{"cuxx", "lee", "wuhan", "f", 20}
	fmt.Println(persion2)
	fmt.Println(persion2.firstName)

	// Gretting
	fmt.Println(persion2.greet())

	persion2.getMarried("xxx")
	fmt.Println(persion2.greet())




}


