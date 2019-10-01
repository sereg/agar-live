package grid

import (
	"testing"
)
//
func BenchmarkStruct(b *testing.B) {
	grid := func() Grid {
		return NewStruct(50, 6400)
	}
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		testGrid(grid)
	}
}

func BenchmarkString(b *testing.B) {
	grid := func() Grid {
		return NewString(50, 6400)
	}
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		testGrid(grid)
	}
}

func BenchmarkMultiply(b *testing.B) {
	grid := func() Grid {
		return NewMultiply(50, 6400)
	}
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		testGrid(grid)
	}
}

func BenchmarkArray(b *testing.B) {
	grid := func() Grid {
		return NewArray(50, 5000, 5000)
	}
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		testGrid(grid)
	}
}