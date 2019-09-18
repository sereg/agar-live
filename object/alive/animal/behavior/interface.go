package behavior

import (
	"agar-life/object/alive/animal"
	"agar-life/object/alive/plant"
)

type Behavior interface{
	SetDirection(animals []animal.Animal, plants []plant.Plant)
}
