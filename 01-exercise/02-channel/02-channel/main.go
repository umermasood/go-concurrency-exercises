package main

import "fmt"

func main() {
	ch := make(chan int)
	go func() {
		defer close(ch)
		for i := 0; i < 6; i++ {
			// send iterator over channel
			fmt.Printf("Sending %d\n", i)
			ch <- i
		}
	}()

	// range over channel to receive values
	for i := range ch {
		fmt.Printf("Receiving %d\n", i)
	}
}
