package concurrency

import "fmt"

func ExamplePlusOne() {
	c := make(chan int)
	go func() {
		defer close(c)
		c <- 5
		c <- 3
		c <- 8
		c <- 10
	}()

	for num := range PlusOne(PlusOne(c)) {
		fmt.Println(num)
	}
	// Output:
	// 7
	// 5
	// 10
	// 12
}

func ExampleChain() {
	c := make(chan int)
	go func() {
		defer close(c)
		c <- 1
		c <- 3
		c <- 8
		c <- 10
	}()
	// Higher order functions
	PlusTwo := Chain(PlusOne, PlusOne)
	for num := range PlusTwo(c) {
		fmt.Println(num)
	}
	// Output:
	// 3
	// 5
	// 10
	// 12
}
