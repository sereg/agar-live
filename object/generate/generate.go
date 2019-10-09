package generate

import (
	"agar-life/math"
	"agar-life/object/alive"
	_const "agar-life/world/const"
	"strconv"
)

func Generate(el alive.Alive, opts ...Option) {
	opt := DefaultOptions()
	for _, o := range opts {
		o(&opt)
	}
	if el.Danger() {
		el.SetColor(_const.PoisonColor)
		el.SetSize(_const.PoisonSize)
	} else {
		if opt.color == nil {
			el.SetColor(getRandomColor())
		} else {
			el.SetColor(*opt.color)
		}
		el.SetSize(opt.size)
	}
	el.SetCrd(opt.crdFn(opt.w, opt.h))
	el.Revive()
	if opt.name != "" {
		el.SetGroup(opt.name)
	}
}

type crdFunc func(x, y float64) (float64, float64)

type Options struct {
	w, h, size float64
	name       string
	crdFn      crdFunc
	color *string
}

func DefaultOptions() Options {
	return Options{
		size:  3,
		crdFn: RandomCrd,
	}
}

//Option it is type for config of declare option
type Option func(*Options)

// SetGroup sets name option
func Name(name string) Option {
	return func(o *Options) {
		o.name = name
	}
}

// SetColor sets color option
func Color(color string) Option {
	return func(o *Options) {
		o.color = &color
	}
}

// Crd sets function for setting crdFn
func Crd(crdFn crdFunc) Option {
	return func(o *Options) {
		o.crdFn = crdFn
	}
}

func RandomCrd(x, y float64) (float64, float64) {
	return float64(math.Random(0, int(x))), float64(math.Random(0, int(y)))
}

func FixCrd(x, y float64) crdFunc {
	return func(float64, float64) (float64, float64) {
		return x, y
	}
}

// WorldWH sets size of world option
func WorldWH(w, h float64) Option {
	return func(o *Options) {
		o.w, o.h = w, h
	}
}

// Type sets name option
func Size(size float64) Option {
	return func(o *Options) {
		o.size = size
	}
}

func getRandomColor() string {
	r := strconv.FormatInt(int64(math.Random(50, 250)), 16)
	g := strconv.FormatInt(int64(math.Random(50, 250)), 16)
	b := strconv.FormatInt(int64(math.Random(50, 250)), 16)
	return "#" + r + g + b
}
