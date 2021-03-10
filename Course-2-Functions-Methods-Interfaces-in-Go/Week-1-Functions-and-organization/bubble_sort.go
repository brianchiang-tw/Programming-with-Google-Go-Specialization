package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Swap(slice []int, idx int) {

	// swap value by tuple assignment
	slice[idx], slice[idx+1] = slice[idx+1], slice[idx]
	return
}

//End of function Swap

func BubbleSort(slice []int) {

	size := len(slice)

	// Larger element moves to top gradually
	for i := 0; i < size; i++ {

		for j := 0; j < size-1-i; j++ {

			if slice[j] > slice[j+1] {
				Swap(slice, j)
			}
		}
	}
}

//End of function BubbleSort

func Min(x, y int) int {
	if x < y {
		return x

	} else {
		return y
	}
}

//End of function Min

func main() {

	userInput := bufio.NewScanner(os.Stdin)

	fmt.Println("Please input a sequence separated by whitespace, up to 10 integers.\n> ")

	userInput.Scan()
	inputString := userInput.Text()

	tokens := strings.Fields(inputString)

	sequence_len := Min(len(tokens), 10)

	var integer_slice []int

	for i := 0; i < sequence_len; i++ {

		integer, error := strconv.Atoi(tokens[i])

		if error == nil {
			integer_slice = append(integer_slice, integer)
		} else {
			fmt.Println("There is invalid input in input sequence")
			fmt.Println("Please rerun this program and input again")
			os.Exit(0)
		}
	}

	BubbleSort(integer_slice)

	fmt.Println("Integer slice sorted in ascending order")
	fmt.Println(integer_slice)

	return
}

//End of function main
