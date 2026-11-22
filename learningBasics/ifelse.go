package main

import "fmt"

func main() {

	// WE DON'T HAVE TERNARY if-operator IN GoLang


	if age:=10;age>=18{
		fmt.Println("elegible to vote")
	}else{
		fmt.Println("not elegible")
	}

	if (7%2==0) {
		fmt.Println("7 is even")
	}else{
		fmt.Println("7 is odd")
	}

	if num := 9; num < 0 {
        fmt.Println(num, "is negative")
    } else if num < 10 {
        fmt.Println(num, "has 1 digit")
    } else {
        fmt.Println(num, "is not a single digit number")
    }
// Note that you don’t need parentheses around conditions in Go, but that the braces are required.
	name:="sahil" 
	age:=21

	if name=="Sahil" || age ==21{
		fmt.Println(name , age)  // sahil , 21
	}


}