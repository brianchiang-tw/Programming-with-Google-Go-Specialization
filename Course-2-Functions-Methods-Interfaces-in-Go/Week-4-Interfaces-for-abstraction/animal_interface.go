package main

import (
	"fmt"
)

// Struct: Animal
type Animal struct {
	food, locomotion, sound string
}

// function interface

type animalInterface interface {
	Eat()
	Move()
	Speak()
}

// member function Eat

func (ani Animal) Eat() {
	fmt.Println(ani.food)
	return
}

// member function Move

func (ani Animal) Move() {
	fmt.Println(ani.locomotion)
	return
}

// member function Speak

func (ani Animal) Speak() {
	fmt.Println(ani.sound)
	return
}

// Main function

func main() {

	// key: string
	// value: Animal with corresponding action
	animalAction := make(map[string]Animal)

	// create default Animal
	animalAction["cow"] = Animal{"grass", "walk", "moo"}
	animalAction["bird"] = Animal{"worms", "fly", "peep"}
	animalAction["snake"] = Animal{"mice", "slither", "hsss"}

	var currentAnimal animalInterface

	// main loop
	for {
		var command, requestAnimal, requestAction string
		fmt.Print(">")
		fmt.Scan(&command, &requestAnimal, &requestAction)
		if command == "query" {
			currentAnimal = animalAction[requestAnimal]
			switch requestAction {
			case "eat":
				currentAnimal.Eat()
			case "move":
				currentAnimal.Move()
			case "speak":
				currentAnimal.Speak()
			}
		}

		if command == "newanimal" {
			animalAction[requestAnimal] = animalAction[requestAction]
			fmt.Println("Created it!")
		}
	}
}

//End of main function
