package world

import (
	_const "agar-life/world/const"
	"strconv"
	"testing"
)

func benchmarkCycle(size float64, b *testing.B) {
	_const.GridSize = size
	for n := 0; n < b.N; n++ {
		cycle()
	}
}

func BenchmarkCycle(b *testing.B) {
	for i:= 10; i < 150; i +=2{
		b.Run("Cycle" + strconv.Itoa(i), func(b *testing.B) {
			_const.GridSize = float64(i)
			for n := 0; n < b.N; n++ {
				cycle()
			}
		})
	}
}