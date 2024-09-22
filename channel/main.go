package main

import (
	"fmt"
	"time"
)

func main() {

	ch := make(chan int, 3)

	go func() {
		for {
			i := <-ch
			fmt.Println(i)
		}
	}()

	time.Sleep(1 * time.Second)

	ch <- 1

	return
}
