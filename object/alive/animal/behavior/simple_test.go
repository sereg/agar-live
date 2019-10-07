package behavior

import (
	"agar-life/object/alive/animal"
	"agar-life/object/generate"
	gnt "agar-life/object/generate"
	"testing"
)

var (
	w, h = 100.0, 100.0
)

func TestSetDirection(t *testing.T) {
	sb := NewSimple(w, h)
	animal2 := animal.NewBase()
	generate.Generate(animal2, gnt.WorldWH(w, h), gnt.Name("a"), gnt.Size(6))
	animal2.SetCrd(0, 0)
	sb.Direction(animal2, nil, nil)
	_ = 4
}