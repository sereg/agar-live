package generate

import (
	"agar-life/math"
	"agar-life/object/alive"
	"strconv"
)

func Generate(el alive.Alive, opts ...Option)	 {
	opt := DefaultOptions()
	for _, o := range opts {
		o(&opt)
	}

	el.Color(getRandomColor())
	el.Size(opt.size)
	el.Crd(
		float64(math.Random(int(0), int(opt.w))),
		float64(math.Random(int(0), int(opt.h))),
	)
	el.Hidden(false)
	el.Name(opt.name)
}

type Options struct {
	w, h, size float64
	name string
}

func DefaultOptions() Options{
	return Options{
		size: 3,
	}
}

//Option it is type for config of declare option
type Option func(*Options)

// Name sets name option
func Name(name string) Option {
	return func(o *Options) {
		o.name = name
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
