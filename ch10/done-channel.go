package main

import "fmt"

func searchData(s string, searchers []func(string) []string) []string {
	done := make(chan struct{})
	resultChan := make(chan []string)
	for _, searcher := range searchers {
		go func(f func(string) []string) {
			select {
			case resultChan <- f(s):
				fmt.Println("結果が戻ってきた")
			case <-done:
				fmt.Println("doneを選択")
			}
		}(searcher)
	}

	r := <-resultChan
	close(done)
	return r
}
