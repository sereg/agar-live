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

func Te1stLog(t *testing.T) {
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

//s= v0*t * 0.5*a*t^2
func TestReduceSpeed(t *testing.T){
	dist := 200.0
	tt := 60
	v0 := 40.0
	a := getAcceleration(v0, float64(tt), dist)
	for i:= 0; i < tt; i++{
		v0 -=a
	}
	fmt.Println(a)
	fmt.Println(v0)
}

//a=(s-v0*t)/(0.5*t*t)
func getAcceleration(v0, t, s float64) float64 {
	return math.Abs((s-v0*t) / (t*t))
}
