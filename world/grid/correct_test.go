package grid

import "testing"

func TestCorrectArray(t *testing.T) {
	testData := []struct{
		name string
		fn func() Grid
	}{
		{
			name: "array",
			fn: func() Grid {
				return NewArray(50, 5000, 5000)
			},
		},
		{
			name: "struct",
			fn: func() Grid {
				return NewStruct(50, 6400)
			},
		},		{
			name: "string",
			fn: func() Grid {
				return NewString(50, 6400)
			},
		},
		{
			name: "Multiply",
			fn: func() Grid {
				return NewMultiply(50, 6400)
			},
		},
	}
	expectedUsed, expectedFound := 6400, 207025
	for _, v := range testData{
		l, f := testGrid(v.fn)
		if l != expectedUsed {
			t.Errorf("%s used - %d, expected - %d", v.name, l, expectedUsed)
		}
		if f != expectedFound {
			t.Errorf("%s found - %d, expected - %d", v.name, f, expectedFound)
		}
	}
}

func testGrid(fn func() Grid) (l, f int) {
	found := []int{}
	grid := fn()
	for i:= 0.0; i < 100; i++ {
		for j := 0.0; j < 100; j++ {
			grid.Set(i * 40, j * 40, int(i))
		}
	}
	for i:= 0.0; i < 100; i++ {
		for j := 0.0; j < 100; j++ {
			found = append(found, grid.GetObjInRadius(i * 40, j * 40, 70)...)
		}
	}
	l = grid.Len()
	f = len(found)
	return
}