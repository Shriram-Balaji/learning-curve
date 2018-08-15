package main

import "fmt"

func longRun(limit int) {
	for index := 0; index < limit; index++ {
		fmt.Println(index)
	}
	go wait("Done. waiting for you to press Enter ↵")
}

func wait(message string) {
	fmt.Println(message)
}

func main() {

	// go-routines are called concurrent to the current function exec
	go longRun(10)
	fmt.Scanln()

	// channels - pipes connecting concurrent go routines
	// we can send and receive from go-routines using channels
	// ie. message passing between concurret threads of execution is possible through channels
	// syntax: make(chan val-type)
	messages := make(chan string)
	// send ping into the `messages` channel
	go func() {
		messages <- "Hello!"
	}()

	messages <- "Buffered"
	// receive ping from the `messages` channel and store it in var receivedMessage
	prefix := "Channel Says: "
	receivedMessage := prefix + <-messages
	fmt.Println(receivedMessage)
	fmt.Println(<-messages)

	//

}
