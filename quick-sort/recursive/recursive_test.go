package recursive_test

import (
	"testing"

	"github.com/cfabrica46/fundamental-algorithms/recursive"
)

func BenchmarkQuicksortRecursive(b *testing.B) {
	slice := []int8{6, 4, 2, 1, 7, 10, 20}
	for n := 0; n < b.N; n++ {
		recursive.QuicksortRecursive(slice, 0, int8(len(slice)-1))
	}
}
