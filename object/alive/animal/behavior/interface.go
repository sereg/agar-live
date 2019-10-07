package behavior

import (
	"agar-life/object"
	"agar-life/object/alive"
	"agar-life/object/alive/animal"
)

type Behavior interface{
	Direction(self animal.Animal, animals []alive.Alive, plants []alive.Alive, cycle uint64) object.Crd
}
