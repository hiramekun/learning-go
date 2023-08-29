package main

import "fmt"

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		v := 1
		ch1 <- v // ch1 からの読み込みを待つ
		v2 := <-ch2
		fmt.Println(v, v2)
	}()
	v := 2
	ch2 <- v // ch2 からの読み込みを待つ
	v2 := <-ch1
	fmt.Println(v, v2)
	// fatal error: all goroutines are asleep - deadlock!
}
