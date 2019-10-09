package species

import (
	"agar-life/object/alive/plant"
)

type poison struct {
	Base
}

func NewPoison() plant.Plant {
	p := poison{}
	p.SetDanger(true)
	p.SetEdible(false)
	return &p
}
