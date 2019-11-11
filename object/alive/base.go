package alive

import (
	"agar-life/object"
)

//Base is basic realization on object alive
type Base struct {
	object.Base
	Deed   bool
	Name   string
	ID     int
	Danger bool
	Edible bool
}
//GetDanger it returns is Danger of the object
func (b Base) GetDanger() bool{
	return b.Danger
}
//GetEdible it returns is edibility
func (b Base) GetEdible() bool{
	return b.Edible
}
//SetDanger set Danger
func (b *Base)SetDanger(st bool) {
	b.Danger = st
}
//SetEdible set Edible
func (b *Base) SetEdible(st bool) {
	b.Edible = st
}

//GetID it gets ID
func (b Base) GetID() int{
	return b.ID
}

//SetID it sets ID
func (b *Base) SetID(id int){
	b.ID = id
}

//Die it kills the object
func (b *Base) Die(){
	b.Deed = true
	b.SetHidden(true)
}

//Revive it revives the object
func (b *Base) Revive(){
	b.Deed = false
	b.SetHidden(false)
}

//GetDead it returns status of the object
func (b Base) GetDead() bool {
	return b.Deed
}

func (b Base) Grow() {
}

func (b Base) Decrease(){
}

func (b Base) GetGroup() string{
	return b.Name
}
func (b *Base) SetGroup(name string) {
	b.Name = name
}

func (b Base) GetGlueTime() uint64 {
	return 0
}

func (b Base) SetGlueTime(cycle uint64) {
}