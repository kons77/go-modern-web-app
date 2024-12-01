package main

import "log"

func main() {
	var myString string
	myString = "Green"

	log.Println("myString is set to", myString)

	changeUsingPointer(&myString) // to point to the location in memory
	log.Println("after fucn call myString is set to", myString)

}

func changeUsingPointer(s *string) {
	log.Println("s is set to", s) // memory address
	newValue := "Red"
	*s = newValue
}
