package behavior

import (
	"agar-life/object/alive"
	"agar-life/object/alive/animal"
)

type Behavior interface{
	SetDirection(self animal.Animal, animals []alive.Alive, plants []alive.Alive)
}
