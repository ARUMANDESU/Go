package main

import (
	"learningGolang/mocking"
	"os"
	"time"
)

func main() {
	sleeper := &mocking.ConfigurableSleeper{Duration: 2 * time.Second, SleepFunc: time.Sleep}
	mocking.Countdown(os.Stdout, sleeper)
}
