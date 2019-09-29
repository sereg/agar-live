package world

import (
	"fmt"
	"testing"
)

func TestCycle(t *testing.T) {
	world := NewWorld(0, 1, 100, 100)
	for i := 0; i < 100; i++ {
		world.Cycle()
		animalList := world.GetAnimal()
		for _, v := range animalList {
			fmt.Println(v)
		}
	}
}
