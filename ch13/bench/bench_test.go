package bench

import (
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
