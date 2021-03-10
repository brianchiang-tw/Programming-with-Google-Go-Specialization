
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	// show prompt message
	fmt.Printf("Please input a string: ")

	// read from standard input
	inputStream := bufio.NewReader(os.Stdin)

	// get input string
	userInputString, error := inputStream.ReadString('\n')


	if error == nil {

		// trim carriage return symbol
		userInputString = strings.Trim(userInputString, "\r\n")

		// force transform to lowercase
		inputStringLowercase := strings.ToLower(userInputString)

		// prefix check
		prefix_i := inputStringLowercase[0] == 'i'

		// content check
		contains_a := strings.Contains(inputStringLowercase, "a")

		// postfix check
		postfix_n := inputStringLowercase[len(inputStringLowercase)-1] == 'n'


		ok := prefix_i && contains_a && postfix_n
		if ok {
			fmt.Printf("Found!\n")

		} else {
			fmt.Printf("Not Found!\n")
		}

	} else {
		fmt.Printf("Error occurred during input")
	}

	return
}
