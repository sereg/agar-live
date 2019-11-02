package main

import (
	"agar-life/canvas"
	"agar-life/world"
	"math/rand"
	"strconv"
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
	space := world.NewWorld(50, 1, jsCon.GetW(), jsCon.GetH())
	//space := world.NewWorldTest(2, 1, jsCon.GetW(), jsCon.GetH())
	fieldPlants := jsCon.NewCanvas()
	fieldAnimals := canvas.Animal{Base: jsCon.NewCanvas()}
	var cycle js.Func
	cycle = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		space.Cycle()
		plant := space.GetPlant()
		if len(plant) > 0 {//TODO rewrite method, return special marker of not update
			fieldPlants.Save()
			//println(len(plant))
			fieldPlants.Refresh()
			for _, v := range plant {
				fieldPlants.Draw(v)
			}
			fieldPlants.Restore()
		}
		animalList := space.GetAnimal()
		fieldAnimals.Save()
		fieldAnimals.Refresh()
		for _, v := range animalList {
			fieldAnimals.Draw(v)
		}
		fieldAnimals.Restore()
		//println("requestAnimationFrame"){212.83 420}{210.09 409.04}
		jsCon.GetWindow().Call("requestAnimationFrame", cycle)
		return nil
	})
	jsCon.GetWindow().Call("requestAnimationFrame", cycle)
	js.Global().Set("cycle", cycle)
	println("WASM Go Initialized field " +  strconv.Itoa(int(jsCon.GetW())) + " " + strconv.Itoa(int(jsCon.GetH())))
	select {}
}