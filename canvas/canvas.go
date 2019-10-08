package canvas

import (
	"agar-life/object"
	"agar-life/object/alive/animal"
	"agar-life/world/const"
	"math"
	"syscall/js"
)

type WH struct{ w, h float64 }

type JS struct {
	value             int
	window, doc, body js.Value
	wh                WH
}

func NewJsConnect() JS {
	jc := JS{}
	jc.window = js.Global()
	jc.doc = jc.window.Get("document")
	jc.body = jc.doc.Get("body")
	jc.wh.h = jc.window.Get("innerHeight").Float()
	jc.wh.w = jc.window.Get("innerWidth").Float()
	return jc
}

func (j *JS) GetWindow() js.Value {
	return j.window
}

func (j *JS) GetW() float64 {
	return j.wh.w
}

func (j *JS) GetH() float64 {
	return j.wh.h
}

func (j *JS) NewCanvas() Base {
	canvas := j.doc.Call("createElement", "canvas")
	canvas.Set("className", "canvas first")
	canvas.Set("height", j.wh.h)
	canvas.Set("width", j.wh.w)
	j.body.Call("appendChild", canvas)
	ctx := canvas.Call("getContext", "2d")
	return Base{canvas: canvas, ctx:ctx, wh:j.wh}
}

type Base struct {
	canvas, ctx js.Value
	wh          WH
}

func (b *Base) Draw(obj object.Object) {
	if obj.Hidden() {
		println("hidden")
		return
	}
	b.ctx.Call("beginPath")
	b.ctx.Call("arc", obj.GetCrd().GetX(), obj.GetCrd().GetY(), obj.Size(), 0, math.Pi*2, false)
	b.ctx.Set("fillStyle", obj.Color())
	b.ctx.Call("fill")
	b.ctx.Call("closePath")
}

func (b *Base) Refresh() {
	b.ctx.Set("fillStyle", "rgb(255, 255, 255)")
	b.ctx.Call("fillRect", 0, 0, b.wh.w, b.wh.h)
	b.Grid(_const.GridSize)
}

func (b *Base) Grid(step float64) {
	b.ctx.Set("strokeStyle", "#cecaca")
	for i := step; i < b.wh.w; i += step {
		b.ctx.Call("beginPath")
		b.ctx.Call("moveTo", i, 0)
		b.ctx.Call("lineTo", i, b.wh.h)
		b.ctx.Call("stroke")
	}
	for i := step; i < b.wh.h; i += step {
		b.ctx.Call("beginPath")
		b.ctx.Call("moveTo", 0, i)
		b.ctx.Call("lineTo", b.wh.w, i)
		b.ctx.Call("stroke")
	}
}

type Animal struct {
	Base
}

func (a *Animal) Draw(obj1 object.Object) {
	if obj1.Hidden() {
		return
	}
	obj := obj1.(animal.Animal)
	a.Base.Draw(obj)
	a.ctx.Call("beginPath")
	a.ctx.Call("rect", obj.GetCrd().GetX()-obj.Vision(), obj.GetCrd().GetY()-obj.Vision(), 2*obj.Vision(), 2*obj.Vision())
	a.ctx.Set("strokeStyle", "#335dbb")
	a.ctx.Call("stroke")
	a.ctx.Set("setLineDash", "[5, 5]")
	a.ctx.Call("closePath")
}

func (a *Animal) Refresh() {
	a.ctx.Call("clearRect", 0, 0, a.wh.w, a.wh.h)
}
