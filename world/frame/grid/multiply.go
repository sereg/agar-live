package grid

type multiply struct {
	cellSize     float64
	data map[int][]int
}

func NewMultiply(size float64, cap int) Grid {
	return &multiply{
		cellSize:     size,
		data: make(map[int][]int, cap),
	}
}

func (g *multiply) Len() int {
	return len(g.data)
}

func (g *multiply) Reset() {

}

func (g *multiply) Set(x, y, size float64, i int) {
	xInt := int(x / g.cellSize)
	yInt := int(y / g.cellSize)
	key := multiplyKey(xInt, yInt)
	if _, ok := g.data[key]; ok {
		g.data[key] = append(g.data[key], i)
	} else {
		g.data[key] = []int{i}
	}
}

func multiplyKey(xInt, yInt int) int {
	return (xInt + 1) * 10000 + yInt
}

func (g *multiply) GetObjInRadius(x, y, radius float64, size float64, exclude int) ([]int, int) {
	ltx, lty := int((x-radius)/g.cellSize), int((y-radius)/g.cellSize)
	rdx, rdy := int((x+radius)/g.cellSize), int((y+radius)/g.cellSize)
	var obj []int
	for cx := ltx; cx <= rdx; cx++ {
		for cy := lty; cy <= rdy; cy++ {
			key := multiplyKey(cx, cy)
			if ob, ok := g.data[key]; ok {
				obj = append(obj, ob...)
			}
		}
	}
	return obj, 0
}