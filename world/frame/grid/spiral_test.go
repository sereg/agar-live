package grid

import "testing"

type spiralData struct{
	x, y int
}

func TestSpiral(t *testing.T) {
	testData := []spiralData{
		spiralData{9,10},
		spiralData{9,11},
		spiralData{10,11},
		spiralData{11,11},
		spiralData{11,10},
		spiralData{11,9},
		spiralData{10,9},
		spiralData{9,9},
		spiralData{8,9},
		spiralData{8,10},
	}
	x, y := 10, 10
	sp := spiral(x, y)
	for i:=0; i < 10; i++ {
		x, y := sp()
		if x != testData[i].x || y != testData[i].y {
			t.Errorf("spet %d found - %d, %d, expected - %d, %d", i, x, y, testData[i].x, testData[i].y)
		}
	}
}
