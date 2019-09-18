package alive

import "agar-life/object"

//Base is basic realization on object alive
type Base struct {
	object.Base
	deed bool
}

//Die it kills the object
func (b *Base) Die(){
	b.deed = true
	b.Hidden(true)
}

//Revive it revives the object
func (b *Base) Revive(){
	b.deed = false
	b.Hidden(false)
}

//GetDead it returns status of the object
func (b Base) GetDead() bool {
	return b.deed
}

func (b Base) Grow() {
}

func (b Base) Decrease(){
}