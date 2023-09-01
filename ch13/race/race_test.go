package race

import "testing"

/*
➜ go test -race -trimpath
==================
WARNING: DATA RACE
Read at 0x00c00018e148 by goroutine 11:

	race.getCounter.func1()
	    race/race.go:12 +0x40

Previous write at 0x00c00018e148 by goroutine 7:

	race.getCounter.func1()
	    race/race.go:12 +0x50

Goroutine 11 (running) created at:

	race.getCounter()
	    race/race.go:10 +0x74
	race.TestGetCounter()
	    race/race_test.go:44 +0x24
	testing.tRunner()
	    testing/testing.go:1595 +0x194
	testing.(*T).Run.func1()
	    testing/testing.go:1648 +0x40

Goroutine 7 (running) created at:

	race.getCounter()
	    race/race.go:10 +0x74
	race.TestGetCounter()
	    race/race_test.go:44 +0x24
	testing.tRunner()
	    testing/testing.go:1595 +0x194
	testing.(*T).Run.func1()
	    testing/testing.go:1648 +0x40

==================
--- FAIL: TestGetCounter (0.00s)

	testing.go:1465: race detected during execution of test

FAIL
exit status 1
FAIL    race    0.232s
*/
func TestGetCounter(t *testing.T) {
	counter := getCounter()
	if counter != 5000 {
		t.Error("想定外のカウンタ値: ", counter)
	}
}
