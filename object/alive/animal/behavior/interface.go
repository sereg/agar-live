package behavior

import "agar-life/object/alive"

type Behavior interface{
	SetDirection(animals []alive.Alive, plants []alive.Alive)
}
