package behavior

import (
	"agar-life/object/alive/animal"
	"testing"
)

func TestSetDirection(t *testing.T) {
	sb := NewSimple(100, 100)
	animal := animal.Newbase()
	animal.Crd(50, 50)
	sb.SetDirection(animal, nil, nil)
	_ = 4
}