package main

import "fmt"



func main() {

	message := "Help!"
	fmt.Println("Original message is: ", message)

	a, b, c := 2, false, 4.6 //testing unassiged variable int, bool, double

	var greeting *string = &message
	fmt.Println("memory location address  message: ", greeting)

	// greeting is the memory location of message; change string literal to which it points
	*greeting = "hi"
	fmt.Println(*greeting)

	fmt.Println("message: ", message,"greeting: ", *greeting)
	fmt.Println(a, b, c)

}
