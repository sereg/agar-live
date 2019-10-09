package vector

import (
	"math"
	"testing"
)

func TestGetVec(t *testing.T) {

}

func TestPerpendicular(t *testing.T) {

}

func TestAngle(t *testing.T) {

}

func TestByAngle(t *testing.T) {

}

type xy struct {
	x, y float64
}


func TestAddAngle(t *testing.T) {
	testRes := map[float64]xy{
		0: xy{x:100, y:110},
		0.5: xy{x:110, y:100},
		1: xy{x:100, y:90},
		1.5: xy{x:90, y:100},
	}
	startX, startY := 100.0, 100.0
	vec := GetVectorByPoint(startX, startY, 110, 100)
	vec.SetAngle(0)
	count := float64(4)
	step := 2 / count
	for i := 0.0; i < 2; i+=step {
		ve := vec
		angel := math.Pi * i
		ve.AddAngle(angel)
		var ans xy
		ans.x, ans.y = ve.GetPointFromVector(startX, startY)
		if a, ok := testRes[i]; !ok || a!= ans {
			t.Errorf("step - %f, angel - %f, recieved unexpected result, expected - %f, got - %f", i, angel, testRes[step], ans)
		}
	}
}
