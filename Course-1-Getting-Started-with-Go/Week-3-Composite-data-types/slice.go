package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)

	integer_slice := make([]int, 0, 3)

	// main-loop
	for {

		fmt.Print("Enter a integer or type X to leave the program: ")
		scanner.Scan()

		// get user input
		userInput := scanner.Text()

		if len(userInput) != 0 {

			if userInput == "X" {

				// leave the program
				break
			}

			// convert userInput from string to integer
			inputInteger, str_to_int_error := strconv.Atoi(userInput)

			if str_to_int_error != nil {
				fmt.Println("Invalid input, please try again")

			} else {
				integer_slice = append(integer_slice, inputInteger)
				sort.Ints(integer_slice)
				fmt.Println("Integer slice in sorted orrder: ", integer_slice)
			}

		} else {
			// Empty input
			fmt.Println("No input, please try again")
			continue
		}

	}

	if error := scanner.Err(); error != nil {

		// OS　IO Error hanlde
		log.Println(error)
	}

}
