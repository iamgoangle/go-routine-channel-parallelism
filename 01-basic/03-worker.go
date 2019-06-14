package main

import (
	"fmt"
	"time"
)

func result(register chan int) {
	for {
		select {
		case r := <-register:
			fmt.Println(r)
		default:
			time.Sleep(time.Second)
			fmt.Println("end of result")
		}
	}
}

func main() {
	register := make(chan int)

	// register 10 clients in the same times
	for i := 1; i <= 10; i++ {
		go func(i int) {
			register <- i
		}(i)
	}

	result(register)
}
