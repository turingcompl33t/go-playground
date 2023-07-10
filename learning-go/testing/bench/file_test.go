package bench

import (
	"fmt"
	"testing"
)

func TestFileLen(t *testing.T) {
	expected := 65204

	result, err := FileLen("testdata/data.txt", 1)
	if err != nil {
		t.Fatal(err)
	}
	if result != expected {
		t.Errorf("Expected %d, got %d", expected, result)
	}
}

// Ensure the compiler does not optimize away calls to FileLen
var blackhole int

func BenchmarkFileLen(b *testing.B) {
	for _, v := range []int{1, 10, 100, 1000, 10000, 100000} {
		b.Run(fmt.Sprintf("FileLen-%d", v), func(b *testing.B) {
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
