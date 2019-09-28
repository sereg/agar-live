package world

import (
	"fmt"
	"testing"
)

func TestCycle(t *testing.T) {
	world := NewWorld(0, 1, 1000, 1000)
	world.Cycle()
	animalList := world.GetAnimal()
	for _, v := range animalList {
		fmt.Println(v)
	}
}
