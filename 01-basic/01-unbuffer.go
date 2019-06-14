package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	register := make(chan bool)

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go registerClient(&register, &wg)
	}

	go func() {
		wg.Wait()
		close(register)
	}()

	for v := range register {
		fmt.Println(v)
	}
}

func registerClient(r *chan bool, wg *sync.WaitGroup) {
	defer wg.Done()
	*r <- true
}
