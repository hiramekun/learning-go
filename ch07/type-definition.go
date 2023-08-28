package main

import "fmt"

type Score int
type HighScore Score
type Person struct {
	LastName  string
	FirstName string
	Age       int
}
type Employee Person

func main() {
	var i int = 300
	var s Score = 100
	var hs HighScore = 200
	//hs = s compile error
	//s = i compile error
	s = Score(i)
	hs = HighScore(s)
	fmt.Println(s, hs) // 300 300
	hhs := hs + 20
	fmt.Println(hhs) // 320
}
