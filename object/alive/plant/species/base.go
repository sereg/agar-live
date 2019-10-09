package species

import (
	"agar-life/object/alive"
	"agar-life/object/alive/plant"
)

type Base struct {
	alive.Base
}

func NewBase() plant.Plant{
	return &Base{}
}
