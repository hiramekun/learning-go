package main

import (
	"errors"
	"fmt"
	"os"
)

func addTo(base int, vals ...int) []int {
	out := make([]int, 0, len(vals))
	for _, v := range vals {
		out = append(out, base+v)
	}
	return out
}

func divAndRemainder(numerator, denominator int) (int, int, error) {
	if denominator == 0 {
		return 0, 0, errors.New("division by zero")
	}
	return numerator / denominator, numerator % denominator, nil
}

func main() {
	result, remainder, err := divAndRemainder(5, 2)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(result, remainder)
}
