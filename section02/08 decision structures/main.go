package main

import "log"

func main() {
	myNum := 99
	isTrue := false

	if myNum > 99 && !isTrue {
		log.Println("myNum is greater than 99")
	} else if myNum < 100 && isTrue {
		log.Println("myNum is less than 100")
	} else {
		log.Println("something else")
	}

	myVar := "bear"

	switch myVar {
	case "cat":
		log.Println("It's cat")

	case "dog":
		log.Println("It's dog")

	case "fish":
		log.Println("It's fish")

	default:
		log.Println("It's something else")
	}

}
