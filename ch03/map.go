package main

import "fmt"

func main() {
	var nilMap map[string]int
	fmt.Println(len(nilMap))   // 0
	fmt.Println(nilMap["abc"]) // 0
	// nilMap["abc"] = 123
	totalWins := map[string]int{}
	fmt.Println(totalWins == nil) // false
	fmt.Println(totalWins["abc"]) // 0
	totalWins["abc"] = 3
	fmt.Println(totalWins["abc"]) // 3

	//teams := map[string][]string{
	//	"ライターズ":    []string{"夏目", "森", "国木田"},
	//	"ナイツ":      []string{"武田", "徳川", "明智"},
	//	"ミュージシャンズ": []string{"ラベル", "ベートーベン", "リスト"},
	//}

	//ages := make(map[string]int, 10)
}
