package main

import "fmt"

func main() {
	words := []string{"hi", "salutations", "hello"}
	for _, word := range words {
		switch wordLen := len(word); {
		case wordLen < 5:
			fmt.Println(word, "は短い単語です")
		case wordLen > 10:
			fmt.Println(word, "は長い単語です")
		default:
			fmt.Println(word, "はちょうど良い長さです")
		}
	}
	/**
	hi は短い単語です
	salutations は長い単語です
	hello はちょうど良い長さです
	*/
}
