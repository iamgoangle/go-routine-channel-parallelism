package main

import (
	"time"
)

func main() {
	register := make(chan bool)

	func() {
		register <- true
	}()

	func() {
		register <- true
	}()

	func() {
		register <- true
	}()

	time.Sleep(10 * time.Second)
}
