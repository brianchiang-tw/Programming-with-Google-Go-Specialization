package main

import "fmt"
import "math"



func GenDisplaceFn(a, v0, s0 float64) func(float64) float64 {

	DisplacemntOfT := func(t float64) float64 {
		//fmt.Println("Dbg message", ((0.5)*a*t*t + v0*t + s0))
		return ((0.5)*a*math.Pow(t, 2) + v0*t + s0)
	}
	return DisplacemntOfT
}

func main() {

	var a, v0, s0, t float64

	// get input value for a
	fmt.Println("Please input a floating number for acceleration:")
	fmt.Scanln(&a)

	// get input value for v0
	fmt.Println("Please input a floating number for initial velocity:")
	fmt.Scanln(&v0)

	// get input value fo s0
	fmt.Println("Please input a floating number for initial displacement")
	fmt.Scanln(&s0)

	// get input valie for t
	fmt.Println("Please input a floating number for elapsed time")
	fmt.Scanln(&t)

	f := GenDisplaceFn(a, v0, s0)
	fmt.Printf("Displacement at t= %f : %f\n", t, f(t))

	fmt.Printf("Displacement at t= %f : %f\n", 3.0, f(3.0))

	fmt.Printf("Displacement at t= %f : %f\n", 5.0, f(5.0))

}

//End of function
