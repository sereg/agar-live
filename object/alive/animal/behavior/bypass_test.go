package behavior

import (
	"agar-life/object"
	"agar-life/object/alive"
	"agar-life/object/alive/animal/species"
	sp "agar-life/object/alive/plant/species"
	gnt "agar-life/object/generate"
	"strconv"
	"testing"
)

var ()

func TestBypass(t *testing.T) {
	el := species.NewBase()
	gnt.Generate(el, gnt.WorldWH(w, h), gnt.Name("a"+strconv.Itoa(1)), gnt.Size(41), gnt.Crd(gnt.FixCrd(100, 100)))
	direction := object.NewCrd(200, 100)
	poisons := make([]alive.Alive, 1)
	el1 := sp.NewPoison()
	gnt.Generate(el1, gnt.WorldWH(w, h), gnt.Name("a"+strconv.Itoa(1)), gnt.Size(41), gnt.Crd(gnt.FixCrd(150, 100)))
	poisons[0] = el1
	bypass(el, direction, poisons)
}
