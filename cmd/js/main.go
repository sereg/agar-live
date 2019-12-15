package main

import (
	"agar-life/canvas"
	"agar-life/world"
	"math/rand"
	"strconv"
	"strings"
	"syscall/js"
	"time"
)

//set GOARCH=wasm
//set GOOS=js
//go src -o ./assets/lib.wasm cmd/js/main.go
//GOARCH=wasm GOOS=js go build -o ./assets/public/lib.wasm cmd/js/main.go
//go test -cpuprofile profile.out
//go test -memprofile profile.out
//go tool pprof --web profile.out
func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	jsCon := canvas.NewJsConnect()
	w, h := jsCon.GetW(), jsCon.GetH()
	countPlants := 50
	countAnimal := 5
	space := world.NewWorld(countPlants, countAnimal, w, h)
	fieldPlants := jsCon.NewCanvas("first")
	fieldAnimals := &canvas.Animal{Base: *jsCon.NewCanvas("second")}
	story := story{}
	cycle := getCycleFn(&space, fieldPlants, fieldAnimals, &story)

	js.Global().Set("cycle", cycle)

	js.Global().Set("restart", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		story.reset()
		space = world.NewWorld(countPlants, countAnimal, w, h)
		cycle = getCycleFn(&space, fieldPlants, fieldAnimals, &story)
		js.Global().Set("cycle", cycle)
		jsCon.GetWindow().Call("requestAnimationFrame", cycle)
		return nil
	}))

	js.Global().Set("generate", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		story.reset()
		countAnimal = args[0].Int()
		countPlants = args[1].Int()
		space = world.NewWorld(countPlants, countAnimal, w, h)
		cycle = getCycleFn(&space, fieldPlants, fieldAnimals, &story)
		js.Global().Set("cycle", cycle)
		jsCon.GetWindow().Call("requestAnimationFrame", cycle)
		return nil
	}))

	js.Global().Set("export", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		test := space.ExportWorld()
		return test
	}))

	js.Global().Set("import", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		story.reset()
		data := args[0].String()
		space = world.NewWorldFromFile(strings.NewReader(data))
		cycle = getCycleFn(&space, fieldPlants, fieldAnimals, &story)
		js.Global().Set("cycle", cycle)
		jsCon.GetWindow().Call("requestAnimationFrame", cycle)
		return nil
	}))

	js.Global().Set("backward", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		if len(story.story) == 0 {
			return nil
		}
		data := story.getLast()
		story.story = story.story[:len(story.story)-1]
		space = world.NewWorldFromFile(strings.NewReader(data))
		cycle = getCycleFn(&space, fieldPlants, fieldAnimals, &story)
		js.Global().Set("cycle", cycle)
		jsCon.GetWindow().Call("requestAnimationFrame", cycle)
		return nil
	}))

	js.Global().Set("get", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		story.reset()
		x, y := args[0].Float(), args[1].Float()
		info := space.GetID(x, y)
		return info
	}))

	js.Global().Set("changePosition", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		story.reset()
		x, y := args[0].Float(), args[1].Float()
		info := space.GetEl(x, y)
		jsCon.GetWindow().Call("requestAnimationFrame", cycle)
		return info
	}))

	js.Global().Set("addFromJSON", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		story.reset()
		data := args[0].String()
		x, y := args[1].Float(), args[2].Float()
		info := space.AddFromJSON(data, x, y)
		jsCon.GetWindow().Call("requestAnimationFrame", cycle)
		return info
	}))

	js.Global().Set("add", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		story.reset()
		typ := args[0].String()
		x, y := args[1].Float(), args[2].Float()
		_, _, _ = typ, x, y
		return nil
	}))

	js.Global().Set("setSize", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		story.reset()
		data := args[0].String()
		space.SetSize(data)
		return nil
	}))

	js.Global().Set("delete", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		story.reset()
		id := args[0].Int()
		_ = id
		return nil
	}))
	jsCon.GetWindow().Call("requestAnimationFrame", cycle)
	println("WASM Go Initialized field " + strconv.Itoa(int(jsCon.GetW())) + " " + strconv.Itoa(int(jsCon.GetH())))
	select {
		case <-time.After(999999 * time.Second):
	}
}

type story struct {
	story []string
}

func (s *story) reset() {
	s.story = []string{}
}

func (s *story) getLast() string {
	return s.story[len(s.story)-1]
}

func getCycleFn(space *world.World, fieldPlants, fieldAnimals canvas.Canvas, story *story) js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		space.Cycle()
		plant := space.GetPlant()
		if len(plant) > 0 { //TODO rewrite method, return special marker of not update
			fieldPlants.Save()
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
		if space.GetCycle()%20 == 0 {
			story.story = append(story.story, space.ExportWorld())
			if len(story.story) > 20 {
				story.story = story.story[1:]
			}
		}
		return nil
	})
}
