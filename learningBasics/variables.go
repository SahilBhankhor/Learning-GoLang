package main

import "fmt"

func main() {
	// var name string = "Sahil Bhankhor"
	// OR
	var name = "Sahil"
	fmt.Println(name)

	// multi value initialisation
	var num1 , num2 = 2,6
	fmt.Println(num1&num2) // 2
	// var userName string , var age int = "SahilBhankhor" , 10 	// ERROR -->> can only assign same type values
	//

	var e int
	e =10 
	fmt.Println(e)

	// shorthand
	// type-inferance
	fruit:="apple"
	fmt.Println(fruit)
}