package main

import "fmt"

func main() {
	var x = 10
	var y = 30.2
	var z = float64(x) + y
	var d = x + int(y)
	fmt.Println(z, d)
}
