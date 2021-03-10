package main

import (
	"fmt"
)

/*
Raca condition is when multiple threads are trying to access and manipulat the same variable.
In the code below, main, incrementByOne and decrementByOne are all accessing and changing the value of x.
Due to the uncertainty of Goroutine scheduling mechanism, the results of the following program is unpredictable.
*/

func incrementByOne(pt *int) {
	(*pt)++
}

func decrementByOne(pt *int) {
	(*pt)--
}
func main() {

	fmt.Printf("final value of i is nondeterministic due to race condition\n")

	// demo 10 times
	for demoTimes := 0; demoTimes < 10; demoTimes++ {

		i := 0
		
		// lunach two golnag coroutine, repeat 100 times
		for repeat := 0; repeat < 100; repeat++ {

			// i = i + 1
			go incrementByOne(&i)

			// i = i - 1
			go decrementByOne(&i)

		}

		fmt.Printf("final value of i = %d \n", i)
	}

	return
}
