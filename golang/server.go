package main

import (
	"fmt"
	"time"
)

func main() {
	// Create and run a goroutine
	go printNumbers()

	// Create and run a goroutine with an anonymous function
	go func() {
		for i := 0; i < 5; i++ {
			fmt.Println("Anonymous Goroutine:", i)
			time.Sleep(time.Millisecond * 500)
		}
	}()

	// Give the goroutines some time to execute
	time.Sleep(time.Second)

	fmt.Println("Main function has finished.")
}

func printNumbers() {
	for i := 0; i < 5; i++ {
		fmt.Println("Goroutine:", i)
		time.Sleep(time.Millisecond * 500)
	}
}
