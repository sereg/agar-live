package canvas

import (
	"agar-life/object"
	"agar-life/object/alive/animal"
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
	if obj.GetHidden() {
		return
	}
	b.ctx.Call("beginPath")
	b.ctx.Call("arc", obj.GetCrd().GetX(), obj.GetCrd().GetY(), obj.GetSize(), 0, math.Pi*2, false)
	b.ctx.Set("fillStyle", obj.GetColor())
	b.ctx.Call("fill")
	b.ctx.Call("closePath")
}

func (b *Base) Refresh() {
	b.ctx.Set("fillStyle", "rgb(255, 255, 255)")
	b.ctx.Call("fillRect", 0, 0, b.wh.w, b.wh.h)
}

type Animal struct {
	Base
}

func (a *Animal) Draw(obj1 object.Object) {
	if obj1.GetHidden() {
		return
	}
	obj := obj1.(animal.Animal)
	a.Base.Draw(obj)
	a.ctx.Call("beginPath")
	a.ctx.Call("rect", obj.GetCrd().GetX()-obj.GetVision(), obj.GetCrd().GetY()-obj.GetVision(), 2*obj.GetVision(), 2*obj.GetVision())
	a.ctx.Set("strokeStyle", "#335dbb")
	a.ctx.Call("stroke")
	a.ctx.Set("setLineDash", "[5, 5]")
	a.ctx.Call("closePath")
}

func (a *Animal) Refresh() {
	a.ctx.Call("clearRect", 0, 0, a.wh.w, a.wh.h)
}
