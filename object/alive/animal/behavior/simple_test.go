package behavior

import (
	"agar-life/object/alive/animal/species"
	"agar-life/object/generate"
	gnt "agar-life/object/generate"
	"testing"
)

var (
	w, h = 100.0, 100.0
)

func TestSetDirection(t *testing.T) {
	sb := NewSimple(w, h)
	animal2 := species.NewBase()
	generate.Generate(animal2, gnt.WorldWH(w, h), gnt.Name("a"), gnt.Size(6))
	animal2.SetXY(0, 0)
	sb.Action(animal2, nil, nil, 0)
	_ = 4
}