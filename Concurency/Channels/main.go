package main

import (
	"fmt"
	"time"
)

// Channels are the pipes that connect concurrent goroutines.
// You can send values into channels from one goroutine and receive those values into another goroutine

func channels() {
	// Create a new channel with make(chan val-type).
	// channels are typed by the values they convey.
	message := make(chan string)

	// send a value into a channel using the channel <- syntax.
	// Here we send "ping" to the messages channel we made above, from a new goroutine.
	go func() { message <- "ping" }()

	// The <- channel syntax receives a value from the channel.
	// Here we'll receive the "ping" message we sent above and print it out.
	msg := <-message
	fmt.Println(msg)
	// When we run the program the "ping" message is successfully passed from the goroutine to another via our channel.

	// By default sends and receives block until both the sender and receiver are ready.
	// This property allowed us to wait at the end of our program for the "ping" message without having to use any other synchronization.
}

// By default channels are unbuffered, meaning that they will only accept sends(chan <-)
// if there is a corresponding receive ( <- chan) ready to receive the sent value.
// Buffered channels accept a limited number of values without a corresponding receiver for those values.
func channelBuffering() {
	// Here we make channel of strings buffering up to 2 values.
	messages := make(chan string, 2)

	// Because this channel is buffered, we can send these values into the channel without a corresponding concurrent receive.
	messages <- "buffered"
	messages <- "channel"

	// Later we can receive these two values as usual.
	fmt.Println(<-messages)
	fmt.Println(<-messages)
}

// We can use channels to synchronize execution across goroutines.
// Here's an example of using a blocking receive to wait for a goroutine to finish.
// When waiting for multiple goroutines to finish, you may prefer to use a WaitGroup
func channelSynchronization() {
	// Start a worker goroutine, giving it the channel to notify on.
	done := make(chan bool, 1)
	go worker(done)

	// Block until we receive a notification from the worker on the channel.
	<-done
	// If you removed the <- done line from this program, the program would exit before the worker even starred.
}

// this is the function we'll run in goroutine.
// The done channel will be used to notify another goroutine that this function's work is done.
func worker(done chan bool) {
	fmt.Print("working...")
	time.Sleep(time.Second)
	fmt.Println("done")

	done <- true
}

// When using channels as function parameters, you can specify if a channel is meant to only send or receive values.
// This specificity increases the type-safety of the program.

// This ping function only accepts a channel for sending values.
// it would be a compile-time errors to try to receive on this channel.
func ping(pings chan<- string, msg string) {
	pings <- msg
}

// The pong function accepts one channel for receives  (pings) and a second for sends (pongs)
func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings
	pongs <- msg
}

func channelDirections() {

	pings := make(chan string, 1)
	pongs := make(chan string, 1)
	ping(pings, "passed message")
	pong(pings, pongs)
	fmt.Println(<-pongs)
}

func main() {
	// channels()
	// channelBuffering()
	// channelSynchronization()
	channelDirections()
}
