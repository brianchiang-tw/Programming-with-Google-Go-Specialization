package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Min(x, y int) int {

	if x < y {
		return x
	} else {
		return y
	}
}

//end of function Min

//--------------------------------------------

func BubbleSort(slice []int, channel chan []int) {

	size := len(slice)

	fmt.Println("This goroutine sort this part: ")
	fmt.Println("before sorting:", slice)

	// Larger element moves to top gradually
	for i := 0; i < size; i++ {

		for j := 0; j < size-1; j++ {

			if slice[j] > slice[j+1] {

				// swap element
				slice[j], slice[j+1] = slice[j+1], slice[j]
			}
		}

	}
	fmt.Println("after sorting:", slice)
	fmt.Println("----------------------------------")
	// push sorted result into channel
	channel <- slice

	return
}

//end of function BubbleSort
//--------------------------------------------

// merge left and right into mergeResult in ascending sorted order
func Merge(left, right []int) []int {

	mergeResult := make([]int, len(left)+len(right))

	resultIndex := 0

	for len(left) > 0 && len(right) > 0 {
		if left[0] < right[0] {
			mergeResult[resultIndex] = left[0]
			left = left[1:]
		} else {
			mergeResult[resultIndex] = right[0]
			right = right[1:]
		}
		resultIndex += 1
	}

	// merge remaining numbers in left
	for i := 0; i < len(left); i++ {
		mergeResult[resultIndex] = left[i]
		resultIndex += 1
	}

	// merge remaining number in right
	for i := 0; i < len(right); i++ {
		mergeResult[resultIndex] = right[i]
		resultIndex += 1
	}

	//fmt.Println("dbg", mergeResult)
	return mergeResult
}

//end of function merge

//--------------------------------------------

func main() {

	userInput := bufio.NewScanner(os.Stdin)

	fmt.Println("Please input a sequence, up to 16 integers")

	userInput.Scan()

	inputString := userInput.Text()

	// parse input data by whitespace as delimiter
	tokens := strings.Fields(inputString)

	// take at most 16 integers
	sequence_len := Min(len(tokens), 16)

	// a integer slice to store input sequence
	var integer_slice []int

	// covert for string token to integer, one by one
	for i := 0; i < sequence_len; i++ {

		integer, error := strconv.Atoi(tokens[i])

		if error == nil {
			integer_slice = append(integer_slice, integer)

		} else {

			fmt.Println("There is invalid input in input sequence.")
			fmt.Println("Please rerun this program and input again")
			os.Exit(0)
		}

	}

	// size of input sequence
	size := len(integer_slice)
	// partition index
	cut_1 := size / 4
	cut_2 := cut_1 * 2
	cut_3 := cut_1 * 3

	// channel for data transfer with goroutine
	channel := make(chan []int, 4)

	// launch 4 golang goroutine to sort 4 difference pieces
	go BubbleSort(integer_slice[:cut_1], channel)
	go BubbleSort(integer_slice[cut_1:cut_2], channel)
	go BubbleSort(integer_slice[cut_2:cut_3], channel)
	go BubbleSort(integer_slice[cut_3:], channel)

	// pop sorted result of 4 difference pieces
	part_1 := <-channel
	part_2 := <-channel
	part_3 := <-channel
	part_4 := <-channel

	// merge 4 difference pieces
	sorted_slice := Merge((Merge(part_1, part_2)), Merge(part_3, part_4))

	// output result
	fmt.Println("Integer slice sorted in ascending order")
	fmt.Println(sorted_slice)

	return
}

//end of function Main
