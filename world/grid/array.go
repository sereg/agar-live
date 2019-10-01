package grid

import "fmt"

type array struct {
	cellSize     float64
	data [][][]int
}

func NewArray(size, w, h float64) Grid {
	rows := int(w / size) + 1
	columns := int(h / size) + 1
	data := make([][][]int, rows)
	for i := range data {
		data[i] = make([][]int, columns)
	}
	return &array{
		cellSize:     size,
		data: data,
	}
}

func (g *array) Len() int {
	count := 0
	for _, v1 := range g.data {
		for _, v := range v1 {
			if len(v) > 0 {
				count++
			}
		}
	}
	return count
}

func (g *array) Set(x, y float64, i int) {
	xInt := int(x / g.cellSize)
	yInt := int(y / g.cellSize)
	if len(g.data) > xInt{
		d := g.data[xInt]
		if len(d) > yInt {
			d[yInt] = append(d[yInt] , i)
			return
		}
	}
	panic(fmt.Sprintf("inner array has not enough count of element, x - %d, y - %d", xInt, yInt))
}

func (g *array) GetObjInRadius(x, y, vision float64) []int {
	ltx, lty := toInt((x-vision)/g.cellSize), toInt((y-vision)/g.cellSize)
	rdx, rdy := toInt((x+vision)/g.cellSize), toInt((y+vision)/g.cellSize)
	var obj []int
	for cx := ltx; cx <= rdx; cx++ {
		for cy := lty; cy <= rdy; cy++ {
			if len(g.data) > cx{
				d := g.data[cx]
				if len(d) > cy {
					obj = append(obj, d[cy]...)
				}
			}
		}
	}
	return obj
}

func toInt(f float64) int{
	if f < 0 {
		return 0
	}
	return int(f)
}