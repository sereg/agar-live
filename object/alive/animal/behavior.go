package animal

import (
	"agar-life/object"
	"agar-life/object/alive"
)

type Behavior interface{
	Direction(self Animal, animals []alive.Alive, plants []alive.Alive, cycle uint64) (object.Crd, bool)
}
