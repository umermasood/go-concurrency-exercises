package main

import "fmt"

// Implement relaying of message with Channel Direction

func genMsg(ch1 chan<- string) {
	// send message on ch1
	ch1 <- "message"
}

func relayMsg(ch1 <-chan string, ch2 chan<- string) {
	//msg := <-ch1 // receive message on ch1
	//ch2 <- msg   // send it on ch2
	//OR
	ch2 <- <-ch1 // Alternative one liner
}

func main() {
	// create ch1 and ch2
	ch1 := make(chan string)
	ch2 := make(chan string)

	// spin goroutine genMsg and relayMsg
	go genMsg(ch1)
	go relayMsg(ch1, ch2)

	// receive message on ch2
	fmt.Printf("Recieved relayed message: %q\n", <-ch2)
}
