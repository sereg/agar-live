package world

import (
	_const "agar-life/world/const"
	"fmt"
	"math"
	"testing"
)

func TestCycle(t *testing.T) {
	cycle()
}

func cycle() {
	world := NewWorldTest(2, 1, 1000, 1000)
	//world := NewWorld(1000, 10, 1000, 1000)
	for i := 0; i < 10000; i++ {
		world.Cycle()
		animalList := world.GetAnimal()
		for _, v := range animalList {
			_ = v
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

//s= v0*t + 0.5*a*t^2
func Tes1tReduceSpeed(t *testing.T){
	dist := 0.0
	tt := 60.0
	v := 5.0
	//a := getAcceleration(v, tt, dist)
	a := 0.1
	for i:= 0.0; i < tt; i++{
		dist += v
		v -=a
		fmt.Println(i)
		fmt.Println(v)
		fmt.Println(dist)
		if v <= 0 {
			break
		}
	}
	//fmt.Println(a)
	//fmt.Println(getDist(5.0, tt, a))
}

//a=(s-v0*t)/(0.5*t*t)
func getAcceleration(v0, t, s float64) float64 {
	return math.Abs((s-v0*t) / (0.5*t*t))
}

func getDist(v0, t, a float64) float64{
	return v0*t + 0.5*a*t*t
}
