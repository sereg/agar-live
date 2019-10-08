package alive

import (
	"agar-life/object"
)

//Base is basic realization on object alive
type Base struct {
	object.Base
	deed bool
	name string
	id int
}

//ID it gets id
func (b Base) ID() int{
	return b.id
}

//SetID it sets id
func (b *Base) SetID(id int){
	b.id = id
}

//Die it kills the object
func (b *Base) Die(){
	b.deed = true
	b.SetHidden(true)
}

//Revive it revives the object
func (b *Base) Revive(){
	b.deed = false
	b.SetHidden(false)
}

//GetDead it returns status of the object
func (b Base) GetDead() bool {
	return b.deed
}

func (b Base) Grow() {
}

func (b Base) Decrease(){
}

func (b Base) Group() string{
	return b.name
}
func (b *Base) SetGroup(name string) {
	b.name = name
}

func (b Base) GlueTime() uint64 {
	return 0
}

func (b Base) SetGlueTime(cycle uint64) {
}