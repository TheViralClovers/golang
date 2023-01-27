package reciever_structs_and_functions

import "log"

type myStruct struct {
	FirstName string
}

func (m *myStruct) printFirstName() string { //this is a receiver function, which receives a structure myStruct as m
	return m.FirstName
}

func main() {
	var myVar myStruct //variable myVar is set to type myStruct
	myVar.FirstName = "John"

	myVar2 := myStruct{ //another way to declare structs (shorthand)
		FirstName: "Mary",
	}

	log.Println(myVar.FirstName)  //accessing a variable inside a structure, by calling it directly
	log.Println(myVar2.FirstName) //same

	//Now let us access it using the "receiver function created above
	log.Println(myVar.printFirstName()) //This return the FirstName and also gets printed
	log.Println(myVar.printFirstName())
}