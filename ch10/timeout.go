package main

import (
	"errors"
	"time"
)

func timeLimit() (int, error) {
	var result int
	var err error
	done := make(chan struct{})
	go func() {
		result, err = doSomeWork()
		close(done)
	}()

	select {
	case <-done:
		return result, err
	case <-time.After(2 * time.Second):
		return 0, errors.New("timeout")
	}
}

func doSomeWork() (int, error) {
	return 1, nil
}
