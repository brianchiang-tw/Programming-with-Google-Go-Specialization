package main

import (
	"fmt"
	"os"
)

type Animal struct {
	food, locomotion, sound       string
}

// Animal::Member function Eat

func (v *Animal) Eat() {
	// print food
	fmt.Println(v.food)
}

// Animal::Member function Move

func (v *Animal) Move() {
	// print locomotion
	fmt.Println(v.locomotion)
}

// Animal::Member function Speak

func (v *Animal) Speak() {
	// print sound
	fmt.Println(v.sound)
}

// Main function
func main() {

	animal_action := map[string]Animal{
		"cow":   Animal{"grass", "walk", "moo"},
		"bird":  Animal{"worms", "fly", "peep"},
		"snake": Animal{"mice", "slither", "hsss"},
	}

	for {

		// prompt message
		fmt.Print(">")

		var animal, action string

		// get user input: animal, action
		fmt.Scan(&animal, &action)
		
		// get animal and its corresponding pointer
		cur_animal := (animal_action[animal])
		cur_ptr := &cur_animal

		// do the specified action
		switch action {
		case "eat":
			cur_ptr.Eat()
		case "move":
			cur_ptr.Move()
		case "speak":
			cur_ptr.Speak()
		default:
			os.Exit(0)
		}

	}
}

//End of function main
