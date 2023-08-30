package main

import (
	"encoding/json"
	"strings"
)

const data = `
{"name": "フレッド", "age": 40}
{"name": "ジョン", "age": 35}
{"name": "アン", "age": 27}
`

func main() {
	var t struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}
	dec := json.NewDecoder(strings.NewReader(data))

	for dec.More() {
		err := dec.Decode(&t)
		if err != nil {
			panic(err)
		}
		// tを処理する
	}
}
