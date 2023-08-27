package main

import "fmt"

type Person struct {
	LastName  string
	FirstName string
	Age       int
}

func (p Person) String() string {
	return fmt.Sprintf("%s %s (%d)", p.FirstName, p.LastName, p.Age)
}

func main() {
	p := Person{LastName: "山田", FirstName: "太郎", Age: 20}
	output := p.String()
	fmt.Println(output) // 太郎 山田 (20)
}
