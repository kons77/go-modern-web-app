package main

import (
	"log"

	"github.com/kons77/newappp/helpers"
)

const numPool = 1000

func CalculateValue(intChan chan int) {
	randomNumber := helpers.RandomNumber(numPool)
	intChan <- randomNumber
}

func main() {
	intChan := make(chan int)
	defer close(intChan) // good habbit to do taht

	go CalculateValue(intChan)

	num := <-intChan
	log.Println(num)
}
