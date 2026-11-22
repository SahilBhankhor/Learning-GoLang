package main

import "fmt"

func main() {

	i := 10
	for i >0{
		fmt.Println(i)
		i--
	}
	// CLASSIC FOR-LOOP IMPLEMENTATION

	fmt.Println("CLASSIC FOR-LOOP IMPLEMENTATION")
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	//  INFINITE LOOP
	for {
		fmt.Println("infinite Loop")
		break;
	}
	// GO---> RANGE BASED LOOP
	for num := range(5){
		if num%2==0 {continue}
		fmt.Println("5 x ",num,"is",num*5)
		
	}
}