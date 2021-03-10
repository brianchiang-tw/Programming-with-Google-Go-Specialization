package main

import (
	"fmt"
	"sync"
	"time"
)

// Structure of chopstick
type ChopStick struct{ 
	sync.Mutex 
}

// Structure of Philosopher
type Philosopher struct {
	id                            int
	leftChopstick, rightChopstick *ChopStick
}

// Waitgroup for philosopher
var eatWgroup sync.WaitGroup

func main() {

	// total 5 philosophers and 5 chopsticks
	size := 5

	// create 5 chopsticks
	chopStickSlice := make([]*ChopStick, size)
	for i := 0; i < size; i++ {
		chopStickSlice[i] = new(ChopStick)
	}

	// create 5 philosophers with ID, left chopstick and right chopstick
	Philosophers := make([]*Philosopher, size)
	for i := 0; i < size; i++ {
		Philosophers[i] = &Philosopher{
			id: i, leftChopstick: chopStickSlice[i], rightChopstick: chopStickSlice[(i+1)%size]}
		eatWgroup.Add(1)
		go Philosophers[i].eat()

	}
	eatWgroup.Wait()
}
// End of function main


func (p Philosopher) eat() {

	// eat three times for each philosopher
	for j := 0; j < 3; j++ {

		// lock one pair of chopstick
		p.leftChopstick.Lock()
		p.rightChopstick.Lock()

		// eating
		fmt.Printf("Philosopher %d is eating\n", p.id+1)

		// simultation of eating time 
		time.Sleep(time.Millisecond * 500)

		// release one pair of chopstick
		p.rightChopstick.Unlock()
		p.leftChopstick.Unlock()

		fmt.Printf("Philosopher %d is finished eating\n", p.id+1)

	}
	eatWgroup.Done()
}
// End of Philosopher::member function Eat