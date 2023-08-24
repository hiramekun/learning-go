package main

import "fmt"

func main() {
	x := []int{1, 2, 3, 4}
	y := make([]int, 4)
	num := copy(y, x)
	fmt.Println(y, num) // [1 2 3 4] 4

	y2 := make([]int, 2)
	num2 := copy(y2, x)
	fmt.Println(y2, num2) // [1 2] 2

	y3 := make([]int, 2)
	copy(y3, x[1:])
	fmt.Println(y3) // [2 3]

	num3 := copy(x[:3], x[1:])
	fmt.Println(x, num3) // [2 3 4 4] 3
}
