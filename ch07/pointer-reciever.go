package main

import (
	"fmt"
	"time"
)

type Counter struct {
	total       int
	lastUpdated time.Time
}

func (c *Counter) Increment() {
	c.total++
	c.lastUpdated = time.Now()
}

func (c Counter) String() string {
	return fmt.Sprintf("total: %d, lastUpdated: %v", c.total, c.lastUpdated)
}

func main() {
	var c Counter
	fmt.Println(c.String()) // total: 0, lastUpdated: 0001-01-01 00:00:00 +0000 UTC
	c.Increment()
	fmt.Println(c.String()) // total: 1, lastUpdated: 2023-08-28 00:16:20.686966 +0900 JST m=+0.000321376
}
