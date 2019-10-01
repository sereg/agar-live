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

func (g *multiply) Set(x, y float64, i int) {
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

func (g *multiply) GetObjInRadius(x, y, vision float64) []int {
	ltx, lty := int((x-vision)/g.cellSize), int((y-vision)/g.cellSize)
	rdx, rdy := int((x+vision)/g.cellSize), int((y+vision)/g.cellSize)
	var obj []int
	for cx := ltx; cx <= rdx; cx++ {
		for cy := lty; cy <= rdy; cy++ {
			key := multiplyKey(cx, cy)
			if ob, ok := g.data[key]; ok {
				obj = append(obj, ob...)
			}
		}
	}
	return obj
}