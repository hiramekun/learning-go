package bench

import (
	"fmt"
	"math/rand"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	makeData()
	exitVal := m.Run()
	os.Remove("testdata/data.txt")
	os.Exit(exitVal)
}

// makeData makes our data file for us. Rather than checking in a large file, we recreate it for the test.
// By setting the random seed to the same value every time, we ensure that we generate the same file every time.
// This random seed generates a file that's 65,204 bytes long.
func makeData() {
	file, err := os.Create("testdata/data.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	rand.Seed(1)
	for i := 0; i < 10000; i++ {
		data := makeWord(rand.Intn(10) + 1)
		file.Write(data)
	}
}

func makeWord(l int) []byte {
	out := make([]byte, l+1)
	for i := 0; i < l; i++ {
		out[i] = 'a' + byte(rand.Intn(26))
	}
	out[l] = '\n'
	return out
}

/*
➜ go test  -v
=== RUN   TestFileLen
--- PASS: TestFileLen (0.04s)
PASS
ok      test_examples/bench     0.153s
*/
func TestFileLen(t *testing.T) {
	result, err := FileLen("testdata/data.txt", 1)
	if err != nil {
		t.Error("FileLen should not return an error, but got", err)
	}
	if result != 65204 {
		t.Error("FileLen should be 65204, but got", result)
	}
}

var blackhole int

/*
➜ go test -bench=. -benchmem
goos: darwin
goarch: arm64
pkg: test_examples/bench
BenchmarkFileLen1-10                  34          31781127 ns/op           65333 B/op      65208 allocs/op
PASS
ok      test_examples/bench     1.473s
*/
func BenchmarkFileLen1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		result, err := FileLen("testdata/data.txt", 1)
		if err != nil {
			b.Fatal(err)
		}
		blackhole = result
	}
}

/*
➜ go test -bench=. -benchmem
goos: darwin
goarch: arm64
pkg: test_examples/bench
BenchmarkFileLen1-10                  33          31028091 ns/op           65333 B/op      65208 allocs/op
BenchmarkFileLen/FileLen1-10                  38          30989535 ns/op           65333 B/op      65208 allocs/op
BenchmarkFileLen/FileLen10-10                375           3178990 ns/op          104480 B/op       6525 allocs/op
BenchmarkFileLen/FileLen100-10              3457            335977 ns/op           73377 B/op        657 allocs/op
BenchmarkFileLen/FileLen1000-10            23635             52555 ns/op           68736 B/op         70 allocs/op
BenchmarkFileLen/FileLen10000-10           60751             20307 ns/op           82048 B/op         11 allocs/op
BenchmarkFileLen/FileLen100000-10          53779             22461 ns/op          213121 B/op          5 allocs/op
PASS
ok      test_examples/bench     9.926s
*/
func BenchmarkFileLen(b *testing.B) {
	for _, v := range []int{1, 10, 100, 1_000, 10_000, 100_000} {
		b.Run(fmt.Sprintf("FileLen%d", v), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				result, err := FileLen("testdata/data.txt", v)
				if err != nil {
					b.Fatal(err)
				}
				blackhole = result
			}
		})
	}
}
