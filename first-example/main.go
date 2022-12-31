package main

import (
	"fmt"
	"sync"
)

// wg access to wait group to decrease the elements
func printSomething(s string, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Println(s)
}

func main() {
	var wg sync.WaitGroup

	words := []string{
		"alfa",
		"beta",
		"gamma",
		"pi",
		"zeta",
		"eta",
		"theta",
		"epsilon",
	}

	// waiting for 9 elements of words before going on
	wg.Add(len(words))

	for i, x := range words {
		go printSomething(fmt.Sprintf("%d: %s", i, x), &wg)
	}

	wg.Wait() // good solution

	/**
	go printSomething("First thing to be printed !!!")

	// bad solution
	time.Sleep(1 * time.Second)
	**/

	wg.Add(1)
	printSomething("Second thing to be printed !!!", &wg)
}
