package main

import "fmt"
import "time"

func main() {
	naturals := make(chan int)
	squares := make(chan int)

	// Counter
	go func() {
		for x := 0; x < 10; x++ {
			naturals <- x
		}
		close(naturals)
	}()

	// Squarer
	go func() {
		for {
			x, ok := <-naturals
			if !ok {
				break
			}
			squares <- x * x
		}
		close(squares)
	}()

	// Printer (in main goroutine)
	for {
		x, ok := <-squares
		if !ok {
			break
		}
		fmt.Println(x)
		time.Sleep(time.Millisecond * 200)
	}
	fmt.Println("Exit")
}
