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

func (g array) Len() int {
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

func (g *array) Reset() {
	for i, v1 := range g.data {
		for i1, v := range v1 {
			if len(v) > 0 {
				g.data[i][i1] = make([]int, 0, 10)
			}
		}
	}
}

func (g *array) Set(x, y float64, i int) {
	xInt := toInt(x / g.cellSize)
	yInt := toInt(y / g.cellSize)
	if len(g.data) > xInt{
		d := g.data[xInt]
		if len(d) > yInt {
			d[yInt] = append(d[yInt] , i)
			return
		}
	}
	panic(fmt.Sprintf("inner array has not enough count of element, x - %d, y - %d", xInt, yInt))
}

func (g array) GetObjInRadius(x, y, radius float64, exclude int) []int {
	ltx, lty := toInt((x-radius)/g.cellSize), toInt((y-radius)/g.cellSize)
	rdx, rdy := toInt((x+radius)/g.cellSize), toInt((y+radius)/g.cellSize)
	xO := toInt(x/g.cellSize)
	yO := toInt(y/g.cellSize)
	var obj []int
	for cx := ltx; cx <= rdx; cx++ {
		for cy := lty; cy <= rdy; cy++ {
			if len(g.data) > cx{
				d := g.data[cx]
				if len(d) > cy {
					if exclude != -1 && xO == cx && yO == cy {
						obj = append(obj, excludeByID(d[cy], exclude)...)
					} else {
						obj = append(obj, d[cy]...)
					}
				}
			}
		}
	}
	return obj
}

func excludeByID(a []int, id int) []int{
	for k, v := range a {
		if v == k {
			return removeFromInt(a, k)
		}
	}
	return a
}

func removeFromInt(a []int, i int) []int {
	a[i] = a[len(a)-1] // Copy last element to index i.
	a[len(a)-1] = 0  // Erase last element (write zero value).
	a = a[:len(a)-1]
	return a
}

func toInt(f float64) int{
	if f < 0 {
		return 0
	}
	return int(f)
}