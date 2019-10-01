package grid

type Grid interface {
	Set(x, y float64, i int)
	GetObjInRadius(x, y, vision float64) []int
	Len() int
}
