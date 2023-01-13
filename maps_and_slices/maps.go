package maps_and_slices

import "log"

type User struct {
	FirstName string
	LastName  string
}

func main() {
	//Creating a map that links strings to "User" data structs
	myMap := make(map[string]User) //syntax for creating a map (analogous to a dictionary)

	me := User{ //creating a var me of type User
		FirstName: "Dhruva",
		LastName:  "Kashyap",
	}

	myMap["me"] = me //Here we have assigned the value of the contents of the User variable me to the string "me"

	log.Println(myMap["me"].FirstName) //accessing map contents and struct contents

}
