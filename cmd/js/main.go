package main

import (
	"agar-life/object"
	"agar-life/object/alive/animal"
	"agar-life/wolrd"
	"fmt"
	"math"
	"math/rand"
	"syscall/js"
	"time"
)
//set GOARCH=wasm
//set GOOS=js
//go build -o ./assets/lib.wasm cmd/js/main.go
//GOARCH=wasm GOOS=js go build -o ./assets/lib.wasm cmd/js/main.go
//go test -cpuprofile profile.out
//go tool pprof --web profile.out
func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	jsCon := newJsConnect()
	world := wolrd.NewWorld(0, 1, jsCon.wh.w, jsCon.wh.h)
	fieldPlants := jsCon.GetCanvas()
	fieldAnimals := jsCon.GetCanvas()
	var cycle js.Func
	cycle = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		plant := world.GetPlant()
		if len(plant) > 0 {
			fieldPlants.refresh()
			for _, v := range plant {
				fieldPlants.draw(v)
			}
		}
		animalList := world.GetAnimal()
		fieldAnimals.refresh()
		for _, v := range animalList {
			fieldAnimals.draw(v)
		}
		println("requestAnimationFrame")
		jsCon.window.Call("requestAnimationFrame", cycle)
		return nil
	})
	jsCon.window.Call("requestAnimationFrame", cycle)
	println("WASM Go Initialized")
	select {}
}

type wh struct{ w, h float64 }

type jsConnect struct {
	value             int
	window, doc, body js.Value
	wh                wh
}

func newJsConnect() jsConnect {
	jc := jsConnect{}
	jc.window = js.Global()
	jc.doc = jc.window.Get("document")
	jc.body = jc.doc.Get("body")
	jc.wh.h = jc.window.Get("innerHeight").Float()
	jc.wh.w = jc.window.Get("innerWidth").Float()
	return jc
}

func (j *jsConnect) GetCanvas() baseCanvas {
	canvas := j.doc.Call("createElement", "canvas")
	canvas.Set("className", "canvas first")
	canvas.Set("height", j.wh.h)
	canvas.Set("width", j.wh.w)
	j.body.Call("appendChild", canvas)
	ctx := canvas.Call("getContext", "2d")
	return baseCanvas{canvas:canvas, ctx:ctx, wh:j.wh}
}

type baseCanvas struct {
	canvas, ctx js.Value
	wh          wh
}

func (b *baseCanvas) draw(obj object.Object) {
	if obj.GetHidden() {
		return
	}
	fmt.Println(obj.GetCrd().GetX(), obj.GetCrd().GetY())
	b.ctx.Call("beginPath")
	b.ctx.Call("arc", obj.GetCrd().GetX(), obj.GetCrd().GetY(), obj.GetSize(), 0, math.Pi*2, false)
	b.ctx.Set("fillStyle", obj.GetColor())
	b.ctx.Call("fill")
	b.ctx.Call("closePath")
}

func (b *baseCanvas) refresh() {
	b.ctx.Set("fillStyle", "rgb(255, 255, 255)")
	b.ctx.Call("fillRect", 0, 0, b.wh.w, b.wh.h)
}

type animalCanvas struct {
	baseCanvas
}

func (a *animalCanvas) draw(obj1 object.Object) {
	if obj1.GetHidden() {
		return
	}
	obj := obj1.(animal.Animal)
	a.baseCanvas.draw(obj)
	a.ctx.Call("beginPath")
	a.ctx.Call("rect", obj.GetCrd().GetX()-obj.GetVision(), obj.GetCrd().GetY()-obj.GetVision(), 2*obj.GetVision(), 2*obj.GetVision())
	a.ctx.Set("strokeStyle", "#335dbb")
	a.ctx.Call("stroke")
	a.ctx.Set("setLineDash", "[5, 5]")
	a.ctx.Call("closePath")
}

func (a *animalCanvas) refresh() {
	a.ctx.Call("clearRect", 0, 0, a.wh.w, a.wh.h)
}
