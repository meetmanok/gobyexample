// _Channels_ are the pipes that connect concurrent
// goroutines. You can send values into 23channels from one
// goroutine and receive those values into another
// goroutine.

package main

import (
	"fmt"
)

func main() {

	// Create a new channel with `make(chan val-type)`.
	// Channels are typed by the values they convey.
	messages := make(chan string)

	// _Send_ a value into a channel using the `channel <-`
	// syntax. Here we send `"ping"`  to the `messages`
	// channel we made above, from a new goroutine.
	go func() { messages <- "ping" }()
	go func() { messages <- "color" }()


	// The `<-channel` syntax _receives_ a value from the
	// channel. Here we'll receive the `"ping"` message
	// we sent above and print it out.
	msg1, msg2 := <-messages, <-messages
	fmt.Println(msg1)
	fmt.Println(msg2)

	//By default sends and receives block until both the sender and receiver are ready.
	//This property allowed us to wait at the end of our program for the "ping" message without having to use any other synchronization.
}
