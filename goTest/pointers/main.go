package main

import "log"

func main() {
	var myColor string
	myColor = "Green"
	log.Println("My color is set to" + myColor)
	changeUsingPointer(&myColor) //passing by reference
	log.Println("My color is set to " + myColor)
}

func changeUsingPointer(s *string) { //This changes the value inside myString without returning any value similar to pointer in other languages
	newVar := "Red" //note := means declaration + assignment (it auto assigns it to a string type)
	*s = newVar
}
