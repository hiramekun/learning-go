package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	words := []string{"山", "sun", "微笑み", "人類学者", "モグラの穴", "mountain", "タコの足とイカの足", "antholopologist", "タコのあしは8本でイカの足は10本"}
	for _, word := range words {
		switch size := utf8.RuneCountInString(word); size {
		case 1, 2, 3, 4:
			fmt.Printf(" [%s] の文字数は%dで、短い単語だ\n", word, size)
		case 5:
			fmt.Printf(" [%s] の文字数は%dで、これはちょうど良い長さだ\n", word, size)
		case 6, 7, 8, 9:
		default:
			fmt.Printf(" [%s] の文字数は%dで、とても長い\n", word, size)
			if n := len(word); size < n {
				fmt.Printf(" %dバイトもある\n", n)
			} else {
				fmt.Println()
			}
		}
	}
	/**
	[山] の文字数は1で、短い単語だ
	[sun] の文字数は3で、短い単語だ
	[微笑み] の文字数は3で、短い単語だ
	[人類学者] の文字数は4で、短い単語だ
	[モグラの穴] の文字数は5で、これはちょうど良い長さだ
	[antholopologist] の文字数は15で、とても長い

	[タコのあしは8本でイカの足は10本] の文字数は17で、とても長い
	45バイトもある

	*/
}
