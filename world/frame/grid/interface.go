package grid

type Grid interface {
	Set(x, y, size float64, i int)
	GetObjInRadius(x, y, radius, size float64, exclude int) ([]int, int)
	Len() int
	Reset()
}
