package main

import (
	"log"
)

// var s string
var s = "seven"

func main() {
	// Variable shadowing
	var s2 = 6

	//var s string
	s := "eight"

	log.Println("s is", s)
	log.Println("s2 is", s2)

	saySomething("xxx")
}

func saySomething(s3 string) (string, string) {
	log.Println("s from the saySomething", s)
	return s3, "World"
}
