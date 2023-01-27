package main

import "log"

// golang is statically typed
func main() {
	var myString string //all variables need to be used. Else program won't compile
	myString = "Hello"  //initializing a variable is not enough to be counted as being used
	log.Println(myString)
	
	myString = saySomething("Hello")
	log.Println(myString)

	myString, _ = saySomethingElse("Goodbye")
	log.Println(myString)
}

// func functionName(variableName variableType) returnType { }
func saySomething(s string) string {
	return s
}

func saySomethingElse(s string) (string, string) {
	return s, "World"
}
