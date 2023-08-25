package main

import (
	"fmt"
	"sort"
)

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

func main() {
	people := []Person{
		{"Pat", "Patterson", 37},
		{"Tracy", "Bobbert", 23},
		{"Fred", "Fredson", 18},
	}

	fmt.Println("初期データ")
	fmt.Println(people)

	sort.Slice(people, func(i, j int) bool {
		return people[i].LastName < people[j].LastName
	})
	fmt.Println("姓でソート")
	fmt.Println(people)
}

func makeMult(base int) func(int) int {
	return func(factor int) int {
		return base * factor
	}
}
