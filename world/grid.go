package world

import "fmt"

type xy struct {
	x, y int
}

type grid struct {
	cellSize float64
	data     map[xy][]int
}

func newGrid(size float64) grid {
	return grid{
		cellSize: size,
		data:     make(map[xy][]int),
	}
}

func (g *grid) set(x, y float64, i int) {
	xInt := int(x / g.cellSize)
	yInt := int(y / g.cellSize)
	if _, ok := g.data[xy{x: xInt, y: yInt}]; ok {
		g.data[xy{x: xInt, y: yInt}] = append(g.data[xy{x: xInt, y: yInt}], i)
	} else {
		g.data[xy{x: xInt, y: yInt}] = []int{i}
	}
}

func (g *grid) getObjInVision(x, y, vision float64) []int {
	ltx, lty := int((x-vision)/g.cellSize), int((y-vision)/g.cellSize)
	rdx, rdy := int((x+vision)/g.cellSize), int((y+vision)/g.cellSize)
	fmt.Printf("point ltx - %d, lty - %d, rdx - %d, rdy -%d\r\n", ltx, lty, rdx, rdy)
	var obj []int
	for cx := ltx; cx < rdx; cx++ {
		for cy := lty; cy < rdy; cy++ {
			if ob, ok := g.data[xy{cx, cy}]; ok {
				obj = append(obj, ob...)
			}
		}
	}
	fmt.Printf("count - %d\r\n", len(obj))
	return obj
}