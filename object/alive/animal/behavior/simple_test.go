package behavior

import (
	"agar-life/object/alive/animal"
	"agar-life/object/generate"
	"testing"
)

var (
	w, h = 100.0, 100.0
)

func TestSetDirection(t *testing.T) {
	sb := NewSimple(w, h)
	animal := animal.NewBase()
	generate.Generate(animal, w, h, "1")
	animal.Crd(50, 50)
	sb.SetDirection(animal, nil, nil)
	_ = 4
}