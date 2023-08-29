package main

import (
	"example.co.jp/package_example/formatter"
	"example.co.jp/package_example/math"
	"fmt"
)

func main() {
	num := math.Double(2)
	output := print.Format(num)
	fmt.Println(output)
	/**
	The number is 4
	*/
}
