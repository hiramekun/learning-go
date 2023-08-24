package main

import "fmt"

func main() {
	var s string = "Hello ☀︎"
	var s2 string = s[4:7]
	var s3 string = s[:5]
	var s4 string = s[6:]

	fmt.Println(s)      // Hello ☀︎
	fmt.Println(s2)     // o �
	fmt.Println(s3)     // Hello
	fmt.Println(s4)     // ☀︎
	fmt.Println(len(s)) // 12
}
