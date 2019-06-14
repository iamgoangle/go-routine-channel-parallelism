package main

import (
	"fmt"
	"time"
)

func main() {
	register := make(chan bool)

	go func() {
		register <- true
	}()

	go func() {
		register <- true
	}()

	go func() {
		register <- true
	}()

	for v := range register {
		fmt.Println(v)
	}

	time.Sleep(10 * time.Second)
}
