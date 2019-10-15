package animal

import (
	"agar-life/math/crd"
	"agar-life/object/alive"
)

type Behavior interface{
	Action(self Animal, animals []alive.Alive, plants []alive.Alive, cycle uint64) (crd.Crd, bool)
}
