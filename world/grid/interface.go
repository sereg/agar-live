package grid

type Grid interface {
	Set(x, y, size float64, i int)
	GetObjInRadius(x, y, radius float64, exclude int) []int
	Len() int
	Reset()
}
