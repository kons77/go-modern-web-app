package main

import (
	"log"
	"sort"
)

func main() {
	var mySlice []string

	mySlice = append(mySlice, "Trevor")
	mySlice = append(mySlice, "John")
	mySlice = append(mySlice, "Marry")

	log.Println(mySlice)

	var mySlice2 []int

	mySlice2 = append(mySlice2, 2)
	mySlice2 = append(mySlice2, 1)
	mySlice2 = append(mySlice2, 3)

	sort.Ints(mySlice2)

	log.Println(mySlice2)

	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	log.Println(numbers)
	log.Println(numbers[0:2])

	names := []string{"one", "seven", "cat", "fish"}
	log.Println(names[1:3])
}
