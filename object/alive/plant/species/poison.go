package species

import (
	"agar-life/object/alive/plant"
)

type poison struct {
	Base
}

func NewBeast() plant.Plant {
	p := poison{}
	p.SetDanger(true)
	p.SetEdible(false)
	return &p
}
