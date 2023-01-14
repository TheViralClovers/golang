package main

import (
	"log"
	"sort"
)

func main() {
	var mySlice []string //declaring a slice of strings

	//Syntax to append 2 strings to a slice
	mySlice = append(mySlice, "Hello", "thing2")
	mySlice = append(mySlice, "Morty")

	log.Println(mySlice)

	var myIntSlice []int

	myIntSlice = append(myIntSlice, 1, 23, 5, 345, 63, 2)
	sort.Ints(myIntSlice) //sorts the slice

	log.Println(myIntSlice)
}
