package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	runtime.GOMAXPROCS(2)
	//ABCDEFGHIJKLMNOPQRSTUVWXYZABCDEFGHIJKLMNOPQRSTUVWXabcdefghYZABCDEFGijklmnopqrstuvwxyzaHIJKLMNOPQRSTUVWXYZbcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyz
	var wg sync.WaitGroup
	wg.Add(2)
	fmt.Println("start GORoutine")

	go func() {
		defer wg.Done()
		for count := 0; count < 3; count++ {
			for char := 'a'; char < 'a'+26; char++ {
				fmt.Printf("%c", char)
			}
		}
	}()

	go func() {
		defer wg.Done()
		for count := 0; count < 3; count++ {
			for char := 'A'; char < 'A'+26; char++ {
				fmt.Printf("%c", char)
			}
		}
	}()

	fmt.Println("waiting to finshing")
	wg.Wait()

	fmt.Println("\n terminating program")
}
