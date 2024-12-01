package main

import "log"

type User struct {
	FirstName string
	LastName  string
}

func main() {
	myMap3 := make(map[string]User)
	me := User{
		FirstName: "Kons",
		LastName:  "Shi",
	}
	myMap3["me"] = me

	log.Println(myMap3["me"].FirstName)

}

func myMap() {
	// var myOtherMap map[string]string  // not whis way

	myMap := make(map[string]string) // this way si common

	myMap["dog"] = "Samson"
	myMap["other-dog"] = "Cassie"

	myMap["dog"] = "Fido"

	log.Println((myMap["dog"]))
	log.Println((myMap["other-dog"]))
	// log.Println((myMap))

	myMap2 := make(map[string]int)
	myMap2["First"] = 1
	myMap2["Second"] = 2

	log.Println(myMap2["First"], myMap2["Second"])
}
