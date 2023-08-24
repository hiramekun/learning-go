package main

import "fmt"

func main() {
	intSet := map[int]bool{}
	vals := []int{5, 10, 2, 5, 8, 7, 3, 9, 1, 2, 10}
	for _, v := range vals {
		intSet[v] = true
	}

	fmt.Println(len(vals), len(intSet)) // 11 8
	fmt.Println(intSet[5])              // true
	fmt.Println(intSet[500])            // false
	if intSet[100] {
		fmt.Println("100が含まれている") // 出力されない
	}

	if intSet[10] {
		fmt.Println("10が含まれている") // 出力される
	}

	intSet2 := map[int]struct{}{}
	for _, v := range vals {
		intSet2[v] = struct{}{}
	}
	if _, ok := intSet2[5]; ok {
		fmt.Println("5が含まれている") // 出力される
	}

	if _, ok := intSet2[500]; ok {
		fmt.Println("500が含まれている") // 出力されない
	}
}
