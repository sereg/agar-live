package world

import (
	"fmt"
	"math"
	"testing"
)

func TestCycle(t *testing.T) {
	world := NewWorld(1, 1, 100, 100)
	for i := 0; i < 100; i++ {
		world.Cycle()
		animalList := world.GetAnimal()
		for _, v := range animalList {
			fmt.Println(v)
		}
	}
}

func TestLog(t *testing.T) {
	for i:= 0; i < 100; i +=2{
		fmt.Printf("%d - %f\r\n", i, reduce(float64(i)))
	}
	for i:= 0; i < 1000; i +=10{
		fmt.Printf("%d - %f\r\n", i, reduce(float64(i)))
	}
}
func reduce(i float64) float64{
	return (20 - math.Log(i * 0.25)) / 10
}
