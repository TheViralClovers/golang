package types_and_structs

import (
	"log"
	"time"
)

type User struct { //This is the syntax for declaration of a structure
	FirstName   string
	LastName    string
	PhoneNumber string
	Age         int
	Date        time.Time
}

//When a variable name is capitalized it becomes a global variable and can be accessed anywhere

func main() {
	user := User{
		FirstName: "Trevor",
		LastName:  "Swelor",
	}

	log.Println(user.FirstName, user.LastName, "Birthdate:", user.Date) //Gives default values if not initialized
}
