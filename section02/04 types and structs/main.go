package main

import (
	"log"
	"time"
)

/* no way because of long and ugly  -  use truct instead
var firstName string
var lastName string
var phoneNumber string
var age int
var birthDate time.Time
*/

// Type can be use as a variable
type User struct {
	// with capital letters because might need this type to be available to another package
	FirstName   string
	LastName    string
	PhoneNumber string
	Age         int
	BirthDate   time.Time
}

var special string // available only from this package
var Special string // available outside from this package
// same thing with functions

func main() {
	user := User{
		FirstName: "Trevor",
		LastName:  "Sawler",
	}

	log.Println(user.FirstName, user.LastName, "BirthDate", user.BirthDate)
}
