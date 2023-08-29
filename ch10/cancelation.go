package main

import "fmt"

func countTo(max int) (<-chan int, func()) {
	ch := make(chan int)
	done := make(chan struct{})
	cancelFunc := func() {
		close(done)
	}

	go func() {
		for i := 0; i < max; i++ {
			select {
			case <-done:
				return
			case ch <- i:
			}
		}
		close(ch)
	}()
	return ch, cancelFunc
}

func main() {
	ch, cancelFunc := countTo(10)
	for i := range ch {
		if i >= 5 {
			break
		}
		fmt.Print(i, " ")
	}
	fmt.Println()
	cancelFunc()

	// 0 1 2 3 4
}
