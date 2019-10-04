package main

import (
	"agar-life/canvas"
	"agar-life/world"
	"math/rand"
	"syscall/js"
	"time"
)
//set GOARCH=wasm
//set GOOS=js
//go build -o ./assets/lib.wasm cmd/js/main.go
//GOARCH=wasm GOOS=js go build -o ./assets/lib.wasm cmd/js/main.go
//go test -cpuprofile profile.out
//go test -memprofile profile.out
//go tool pprof --web profile.out
func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	jsCon := canvas.NewJsConnect()
	space := world.NewWorld(100, 10, jsCon.GetW(), jsCon.GetH())
	//space := world.NewWorldTest(2, 1, jsCon.GetW(), jsCon.GetH())
	fieldPlants := jsCon.NewCanvas()
	fieldAnimals := canvas.Animal{Base: jsCon.NewCanvas()}
	var cycle js.Func
	cycle = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		space.Cycle()
		plant := space.GetPlant()
		if len(plant) > 0 {//TODO rewrite method, return special marker of not update
			fieldPlants.Refresh()
			for _, v := range plant {
				fieldPlants.Draw(v)
			}
		}
		animalList := space.GetAnimal()
		fieldAnimals.Refresh()
		for _, v := range animalList {
			fieldAnimals.Draw(v)
		}
		//println("requestAnimationFrame")
		jsCon.GetWindow().Call("requestAnimationFrame", cycle)
		return nil
	})
	jsCon.GetWindow().Call("requestAnimationFrame", cycle)
	println("WASM Go Initialized")
	select {}
}