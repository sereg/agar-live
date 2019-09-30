package world

type grid struct {
	cellSize     float64
	//data         map[xy][]int
	//dataString   map[string][]int
	dataMultiply map[int][]int
}

func newGrid(size float64, cap int) grid {
	return grid{
		cellSize:     size,
		//data:         make(map[xy][]int, 6400),
		//dataString:   make(map[string][]int, 6400),
		dataMultiply: make(map[int][]int, cap),
	}
}

func (g *grid) set(x, y float64, i int) {
	xInt := int(x / g.cellSize)
	yInt := int(y / g.cellSize)
	key := multiplyKey(xInt, yInt)
	if _, ok := g.dataMultiply[key]; ok {
		g.dataMultiply[key] = append(g.dataMultiply[key], i)
	} else {
		g.dataMultiply[key] = []int{i}
	}
}

func multiplyKey(xInt, yInt int) int {
	return (xInt + 1) * 10000 + yInt
}

func (g *grid) getObjInVision(x, y, vision float64) []int {
	ltx, lty := int((x-vision)/g.cellSize), int((y-vision)/g.cellSize)
	rdx, rdy := int((x+vision)/g.cellSize), int((y+vision)/g.cellSize)
	//fmt.Printf("point ltx - %d, lty - %d, rdx - %d, rdy -%d\r\n", ltx, lty, rdx, rdy)
	var obj []int
	for cx := ltx; cx <= rdx; cx++ {
		for cy := lty; cy <= rdy; cy++ {
			key := multiplyKey(cx, cy)
			if ob, ok := g.dataMultiply[key]; ok {
				obj = append(obj, ob...)
			}
		}
	}
	//fmt.Printf("count - %d\r\n", len(obj))
	return obj
}

//type xy struct {
//	x, y int
//}

//func (g *grid) setStruct(x, y float64, i int) {
//	xInt := int(x / g.cellSize)
//	yInt := int(y / g.cellSize)
//	if _, ok := g.data[xy{x: xInt, y: yInt}]; ok {
//		g.data[xy{x: xInt, y: yInt}] = append(g.data[xy{x: xInt, y: yInt}], i)
//	} else {
//		g.data[xy{x: xInt, y: yInt}] = []int{i}
//	}
//}
//
//func (g *grid) setString(x, y float64, i int) {
//	xInt := int(x / g.cellSize)
//	yInt := int(y / g.cellSize)
//	key := stringKey(xInt, yInt)
//	if _, ok := g.dataString[key]; ok {
//		g.dataString[key] = append(g.dataString[key], i)
//	} else {
//		g.dataString[key] = []int{i}
//	}
//}
//
//func (g *grid) getObjInVisionStruct(x, y, vision float64) []int {
//	ltx, lty := int((x-vision)/g.cellSize), int((y-vision)/g.cellSize)
//	rdx, rdy := int((x+vision)/g.cellSize), int((y+vision)/g.cellSize)
//	//fmt.Printf("point ltx - %d, lty - %d, rdx - %d, rdy -%d\r\n", ltx, lty, rdx, rdy)
//	var obj []int
//	for cx := ltx; cx <= rdx; cx++ {
//		for cy := lty; cy <= rdy; cy++ {
//			if ob, ok := g.data[xy{cx, cy}]; ok {
//				obj = append(obj, ob...)
//			}
//		}
//	}
//	//fmt.Printf("count - %d\r\n", len(obj))
//	return obj
//}
//
//func (g *grid) getObjInVisionString(x, y, vision float64) []int {
//	ltx, lty := int((x-vision)/g.cellSize), int((y-vision)/g.cellSize)
//	rdx, rdy := int((x+vision)/g.cellSize), int((y+vision)/g.cellSize)
//	//fmt.Printf("point ltx - %d, lty - %d, rdx - %d, rdy -%d\r\n", ltx, lty, rdx, rdy)
//	var obj []int
//	for cx := ltx; cx <= rdx; cx++ {
//		for cy := lty; cy <= rdy; cy++ {
//			key := stringKey(cx, cy)
//			if ob, ok := g.dataString[key]; ok {
//				obj = append(obj, ob...)
//			}
//		}
//	}
//	//fmt.Printf("count - %d\r\n", len(obj))
//	return obj
//}
//
//
//func stringKey(xInt, yInt int) string {
//	return strconv.Itoa(xInt) +"x:"+ strconv.Itoa(yInt)+"y"
//}
