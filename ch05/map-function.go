package main

import (
	"fmt"
	"strconv"
)

func add(i int, j int) int { return i + j }
func sub(i int, j int) int { return i - j }
func mul(i int, j int) int { return i * j }
func div(i int, j int) int { return i / j }

func main() {
	var opMap = map[string]func(int, int) int{
		"+": add,
		"-": sub,
		"*": mul,
		"/": div,
	}

	var expressions = [][]string{
		{"2", "+", "3"},
		{"2", "-", "3"},
		{"2", "*", "3"},
		{"2", "/", "3"},
		{"2", "%", "3"},
		{"two", "+", "three"},
		{"2", "+", "three"},
		{"5"},
	}

	for _, expression := range expressions {
		if len(expression) != 3 {
			fmt.Print(expression, " -- 不正な式です\n")
			continue
		}
		p1, err := strconv.Atoi(expression[0])
		if err != nil {
			fmt.Print(expression, " -- ", err, "\n")
			continue
		}

		op := expression[1]
		opFunc, ok := opMap[op]
		if !ok {
			fmt.Print(expression, " -- 不正な演算子です: ", op, "\n")
			continue
		}
		p2, err := strconv.Atoi(expression[2])
		if err != nil {
			fmt.Print(expression, " -- ", err, "\n")
			continue
		}

		result := opFunc(p1, p2)
		fmt.Print(expression, " = ", result, "\n")
	}

	/**
	[2 + 3] = 5
	[2 - 3] = -1
	[2 * 3] = 6
	[2 / 3] = 0
	[2 % 3] -- 不正な演算子です: %
	[two + three] -- strconv.Atoi: parsing "two": invalid syntax
	[2 + three] -- strconv.Atoi: parsing "three": invalid syntax
	[5] -- 不正な式です
	*/
}
