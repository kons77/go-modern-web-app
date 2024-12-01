package main

import "log"

func main() {
	for i := 0; i <= 10; i++ {
		log.Println(i)
	}

	animals := []string{"dog", "fish", "horse", "cow", "cat"}

	for i, animal := range animals {
		log.Println(i, animal)
	}

	for _, animal := range animals {
		log.Println(animal)
	}

	mAnimals := make(map[string]string)
	mAnimals["dog"] = "Fido"
	mAnimals["cat"] = "Fluffy"
	mAnimals["horse"] = "Charlie"

	for mAnimalType, mAnimal := range mAnimals {
		log.Println(mAnimalType, mAnimal)
	}

	var firstLine = "Once upon a midnight dreary"

	// string is a slice of bytes
	for i, l := range firstLine {
		log.Println(i, ":", l)
	}

	type User struct {
		FirstName string
		LastName  string
		Email     string
		Age       int
	}

	var users []User // slice of User
	users = append(users, User{"John", "Smith", "john@smith.com", 30})
	users = append(users, User{"Mary", "Jones", "mary@jones.com", 20})
	users = append(users, User{"Sally", "Brown", "sally@brown.com", 45})
	users = append(users, User{"Alex", "Anderson", "alex@anderson.com", 17})

	for _, u := range users {
		log.Println(u.FirstName, u.LastName, u.Email, u.Age)
	}
}
