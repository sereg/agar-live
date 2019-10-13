package geom

import (
	"agar-life/math"
	"fmt"
	"testing"
)

func TestIntersect(t *testing.T) {

}

type lineByPoint struct {
	p1, p2 Point
	y, slope float64
}

func TestLineByPoint(t *testing.T) {
	testRes := []lineByPoint{
		lineByPoint{
			p1: NewPoint(1, 1), p2: NewPoint(3, 3), y: 0, slope: 1, //y = x
		},
		lineByPoint{
			p1: NewPoint(1, 1), p2: NewPoint(4, 2), y: 0.666667, slope: 0.333333, //y = x
		},
	}
	for _, v := range testRes {
		line := NewLine(v.p1, v.p2)
		if math.ToFixed(line.y, 6) != v.y || math.ToFixed(line.slope, 6) != v.slope {
			fmt.Println(line.y != v.y)
			fmt.Println(line.slope != v.slope)
			t.Errorf("p1 - %v, p2 - %v, recieved unexpected result, expected - '%f', '%f', got - '%f', '%f'", v.p1, v.p2, v.y, v.slope ,line.y ,line.slope)
		}
	}
}