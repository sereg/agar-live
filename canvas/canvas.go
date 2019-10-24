package canvas

import (
	math2 "agar-life/math"
	"agar-life/math/crd"
	"agar-life/math/vector"
	"agar-life/object"
	"agar-life/object/alive"
	"agar-life/object/alive/animal"
	"agar-life/world/const"
	"fmt"
	"math"
	"strconv"
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
	img := j.window.Call("eval", "new Image()")
	img.Set("src", "/img/thorn.png")
	wait := make(chan struct{})
	addImg := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		println("imag loaded")
		wait <- struct{}{}
		return nil
	})
	img.Call("addEventListener", "load", addImg)
	<-wait
	return Base{canvas: canvas, ctx: ctx, img: img, wh: j.wh}
}

type Base struct {
	canvas, ctx, img js.Value
	wh               WH
}

func (b *Base) Save() {
	b.ctx.Call("save")
}

func (b *Base) Restore() {
	b.ctx.Call("restore")
}

//js.Global().Set("add", js.NewCallback(add))
func (b *Base) Draw(obj1 object.Object) {
	if obj1.Hidden() {
		println("hidden")
		return
	}
	obj := obj1.(alive.Alive)
	if obj.Danger() {
		size := obj.Size() * 2.15
		b.ctx.Call("drawImage", b.img, obj.X(), obj.Y(), size, size)
	} else {
		b.ctx.Call("beginPath")
		b.ctx.Call("arc", obj.X(), obj.Y(), obj.Size(), 0, math.Pi*2, false)
		b.ctx.Set("fillStyle", obj.Color())
		b.ctx.Call("fill")
		b.ctx.Call("closePath")
	}
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

//func (a *Animal) Draw(obj1 object.Object) {
//	if obj1.Hidden() {
//		return
//	}
//	obj := obj1.(animal.Animal)
//	a.Base.Draw(obj)
//	a.ctx.Call("beginPath")
//	a.ctx.Call("rect", obj.X()-obj.Vision(), obj.Y()-obj.Vision(), 2*obj.Vision(), 2*obj.Vision())
//	a.ctx.Set("strokeStyle", "#335dbb")
//	a.ctx.Call("stroke")
//	a.ctx.Set("setLineDash", "[5, 5]")
//	a.ctx.Call("closePath")
//
//	a.ctx.Call("beginPath")
//	a.ctx.Call("moveTo", obj.X(), obj.Y())
//	a.ctx.Call("lineTo", obj.Direction().X(), obj.Direction().Y())
//	a.ctx.Call("stroke")
//
//	a.ctx.Set("fillStyle", "#000")
//	a.ctx.Set("font", "bold 12px Arial")
//	a.ctx.Call("fillText", strconv.Itoa(obj.Count())+"/"+strconv.Itoa(int(obj.Size())), obj.X()-obj.Size(), obj.Y())
//}

func (a *Animal) Draw(obj1 object.Object) {
	obj := obj1.(animal.Animal)
	a.ctx.Call("beginPath")
	a.ctx.Call("arc", obj.X(), obj.Y(), 2, 0, math.Pi*2, false)
	a.ctx.Call("stroke")
	count := int((obj.Vision() * math.Pi * 2) / obj.Size())
	addAngel := 2.0 * math.Pi / float64(count)
	addAngelV := 2.0 * math.Pi / float64(count) * 2
	angel := math.Pi
	angelV := math.Pi + addAngelV
	sift := 4.0
	fmt.Println(count)
	expectedDir := -1.0
	for i := 0.0; i < 5; i++ {
		//a.Refresh()
		a.ctx.Set("strokeStyle", getRandomColor())
		//xs1 := obj.X() + obj.Size()*math.Cos(angel)
		//ys1 := obj.Y() + obj.Size()*math.Sin(angel)
		//a.ctx.Call("beginPath")
		//a.ctx.Call("arc", xs1, ys1, 2, 0, math.Pi*2, false)
		//a.ctx.Call("stroke")

		//angel += math.Pi
		//xs2 := obj.X() + obj.Size()*math.Cos(angel)
		//ys2 := obj.Y() + obj.Size()*math.Sin(angel)
		//a.ctx.Call("beginPath")
		//a.ctx.Call("arc", xs2, ys2, 2, 0, math.Pi*2, false)
		//a.ctx.Call("stroke")
		//angel -= math.Pi

		angelV = angel + addAngel*sift + math.Pi/25
		//xf1 := obj.X() + obj.Vision()*math.Cos(angelV)
		//yf1 := obj.Y() + obj.Vision()*math.Sin(angelV)
		//a.ctx.Call("beginPath")
		//a.ctx.Call("arc", xf1, yf1, 2, 0, math.Pi*2, false)
		//a.ctx.Call("stroke")
		angelV += addAngelV
		//xf2 := obj.X() + obj.Vision()*math.Cos(angelV)
		//yf2 := obj.Y() + obj.Vision()*math.Sin(angelV)
		//a.ctx.Call("beginPath")
		//a.ctx.Call("arc", xf2, yf2, 2, 0, math.Pi*2, false)
		//a.ctx.Call("stroke")
		//
		//a.ctx.Call("beginPath")
		//a.ctx.Call("moveTo", xs1, ys1)
		//a.ctx.Call("lineTo", xf1, yf1)
		//a.ctx.Call("stroke")
		//
		//a.ctx.Call("beginPath")
		//a.ctx.Call("moveTo", xs2, ys2)
		//a.ctx.Call("lineTo", xf2, yf2)
		//a.ctx.Call("stroke")

		angelV -=addAngel
		xd := obj.X() + obj.Vision()*math.Cos(angelV)
		yd := obj.Y() + obj.Vision()*math.Sin(angelV)
		a.ctx.Call("beginPath")
		a.ctx.Call("moveTo", obj.X(), obj.Y())
		a.ctx.Call("lineTo", xd, yd)
		a.ctx.Call("stroke")
		if expectedDir == -1 {
			vec := vector.GetVectorByPoint(crd.NewCrd(obj.X(), obj.Y()), crd.NewCrd(xd, yd))
			expectedDir = vec.GetAngle()
			fmt.Println(expectedDir)
		} else {
			a.ctx.Set("strokeStyle", "#000")
			vec := vector.GetVectorByPoint(crd.NewCrd(obj.X(), obj.Y()), crd.NewCrd(obj.X()+obj.Vision(), obj.Y()))
			vec.SetAngle(expectedDir-addAngel*i)
			c := vec.GetPointFromVector(crd.NewCrd(obj.X(), obj.Y()))
			a.ctx.Call("beginPath")
			a.ctx.Call("moveTo", obj.X(), obj.Y())
			a.ctx.Call("lineTo", c.X(), c.Y())
			a.ctx.Call("stroke")
		}
		//break
		angel += addAngel
	}

}

func getRandomColor() string {
	r := strconv.FormatInt(int64(math2.Random(80, 255)), 16)
	g := strconv.FormatInt(int64(math2.Random(20, 180)), 16)
	b := strconv.FormatInt(int64(math2.Random(30, 180)), 16)
	return "#" + r + g + b
}

func (a *Animal) Refresh() {
	a.ctx.Call("clearRect", 0, 0, a.wh.w, a.wh.h)
}
