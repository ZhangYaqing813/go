package main

import (
	"fmt"
)
import "runtime"
import "sync"

var wg sync.WaitGroup

func main() {
	runtime.GOMAXPROCS(1)
	wg.Add(2)

	fmt.Println("create Goroutines")
	go printPrime("A")
	go printPrime("B")

	wg.Wait()

	fmt.Println("over")
}

func printPrime(prefix string) {
	defer wg.Done()

next:
	for outer := 2; outer < 50000; outer++ {
		for innter := 2; innter < outer; innter++ {
			if outer%innter == 0 {
				continue next
			}
		}
		fmt.Printf("%s:%d\n", prefix, outer)
	}
	fmt.Println("completed", prefix)
}
