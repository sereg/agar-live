package world

import (
	_const "agar-life/world/const"
	"fmt"
	"math"
	"testing"
)

func TestCycle(t *testing.T) {
	//world := NewWorldTest(2, 1, 1000, 1000)
	world := NewWorld(1000, 10, 1000, 1000)
	for i := 0; i < 10000; i++ {
		world.Cycle()
		animalList := world.GetAnimal()
		for _, v := range animalList {
			_ = v
			//fmt.Println(v)
		}
	}
}

func TestLog(t *testing.T) {
	for i := 1; i < 100; i += 2 {
		fmt.Printf("%d - %f\r\n", i, reduce(float64(i)))
	}
	for i := 1; i < 1000; i += 10 {
		fmt.Printf("%d - %f\r\n", i, reduce(float64(i)))
	}
}
func reduce(i float64) float64 {
	return (_const.StartSpeed - math.Log(i * _const.SpeedRatio)) / 10
}
