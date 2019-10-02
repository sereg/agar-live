package world

import (
	"fmt"
	"math"
	"testing"
)

func TestCycle(t *testing.T) {
	world := NewWorldTest(2, 1, 1000, 1000)
	for i := 0; i < 10000; i++ {
		world.Cycle()
		animalList := world.GetAnimal()
		for _, v := range animalList {
			_ = v
			//fmt.Println(v)
		}
	}
}

func Tes1tLog(t *testing.T) {
	for i := 0; i < 100; i += 2 {
		fmt.Printf("%d - %f\r\n", i, reduce(float64(i)))
	}
	for i := 0; i < 1000; i += 10 {
		fmt.Printf("%d - %f\r\n", i, reduce(float64(i)))
	}
}
func reduce(i float64) float64 {
	return (8 - math.Log(i*0.25)) / 10
}
