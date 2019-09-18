package plant

import "agar-life/object/alive"

type Plant interface {
	alive.Alive
	GetDanger() bool
	GetEdible() bool
}
