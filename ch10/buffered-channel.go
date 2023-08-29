package main

import (
	"fmt"
	"math/rand"
	"time"
)

func processChannel(ch chan int) []int {
	const maxConc = 10
	results := make(chan int, maxConc)
	for i := 0; i < maxConc; i++ {
		go func() {
			v := <-ch
			results <- process(v)
		}()
	}
	fmt.Println("goroutineを起動しました")

	var out []int
	for i := 0; i < maxConc; i++ {
		out = append(out, <-results)
	}
	return out
}

func process(v int) int {
	returnVal := v * v
	rand.Seed(time.Now().UnixNano()) // シードの指定
	sleepSec := rand.Intn(3)         // 0以上3未満の整数を戻す

	fmt.Println("process:", v, returnVal, sleepSec)
	time.Sleep(time.Duration(sleepSec) * time.Second)
	return returnVal
}

func main() {
	ch := make(chan int)

	var result []int

	go func() { // 処理してもらう数値をchに入れる
		for i := 0; i < 100; i++ {
			ch <- i
		}
	}()

	result = processChannel(ch)

	fmt.Printf("result: %d\n", result)
	/**
	goroutineを起動しました
	process: 0 0 0
	process: 2 4 0
	process: 9 81 1
	process: 8 64 1
	process: 5 25 2
	process: 6 36 2
	process: 7 49 2
	process: 4 16 0
	process: 3 9 1
	process: 1 1 2
	result: [0 4 16 81 9 64 49 36 1 25]
	*/
}
