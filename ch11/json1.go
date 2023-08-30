package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	f := struct {
		Name string
		Age  int
	}{}
	err := json.Unmarshal([]byte(`{"name":"Alice", "occupation":"engineer", "aage":20}`), &f)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%+v\n", f)
	/**
	{Name:Alice Age:0}
	*/
}
