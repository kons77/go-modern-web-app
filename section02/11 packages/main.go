package main

import (
	"log"

	"github.com/kons77/newappp2/hhhelpers"
)

func main() {
	log.Println("Hello")

	var myVar hhhelpers.SomeType
	myVar.TypeName = "Some name"
	log.Println(myVar.TypeName)

}
