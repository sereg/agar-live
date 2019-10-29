package grid

type array struct {
	cellSize float64
	data     [][][]int
}

func NewArray(size, w, h float64) Grid {
	rows := int(w/size) + 1
	columns := int(h/size) + 1
	data := make([][][]int, rows)
	for i := range data {
		data[i] = make([][]int, columns)
	}
	return &array{
		cellSize: size,
		data:     data,
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

func (g *array) Set(x, y, size float64, i int) {
	ltx, lty := toInt((x-size)/g.cellSize), toInt((y-size)/g.cellSize)
	rdx, rdy := toInt((x+size)/g.cellSize), toInt((y+size)/g.cellSize)
	cx, cy := 0, 0
	for cx = ltx; cx <= rdx; cx++ {
		for cy = lty; cy <= rdy; cy++ {
			if len(g.data) > cx {
				d := g.data[cx]
				if len(d) > cy {
					d[cy] = append(d[cy], i)
					continue
				}
			}
		}
	}
}

func (g array) GetObjInRadius(x, y, radius float64, exclude int) []int {
	ltx, lty := toInt((x-radius)/g.cellSize), toInt((y-radius)/g.cellSize)
	rdx, rdy := toInt((x+radius)/g.cellSize), toInt((y+radius)/g.cellSize)
	obj := make([]int, 0, 20)
	for cx := ltx; cx <= rdx; cx++ {
		for cy := lty; cy <= rdy; cy++ {
			if len(g.data) <= cx {
				continue
			}
			d := g.data[cx]
			if len(d) <= cy {
				continue
			}
			obj = append(obj, d[cy]...)
		}
	}
	//cx, cy := toInt((x)/g.cellSize), toInt((y)/g.cellSize)
	//sp := spiral(cx, cy)
	//count := 0.0
	//radius = radius * 2
	//limitCount := radius * radius
	//for {
	//	//if cx > rdx || cy > rdy || cx < ltx || cy < lty{
	//	if limitCount < count{
	//		break
	//	}
	//	count++
	//	if len(g.data) <= cx || cx < 0 {
	//		cx, cy = sp()
	//		continue
	//	}
	//	d := g.data[cx]
	//	if len(d) <= cy || cy < 0 {
	//		cx, cy = sp()
	//		continue
	//	}
	//	obj = append(obj, d[cy]...)
	//	cx, cy = sp()
	//}
	return excludeByID(obj, exclude)
}

var directions = map[string]string{
	"left":  "down",
	"down":  "right",
	"right": "top",
	"top":   "left",
}

func spiral(x, y int) func() (int, int) {
	step := 1
	iterStep := 0
	countStep := 0
	direct := "left"
	return func() (int, int) {
		x, y = getPoint(direct, x, y)
		countStep += 1
		if step == countStep { // step is finished when then countStep is equal step
			countStep = 0
			iterStep += 1
			direct = directions[direct] // after each step change direction
			if iterStep == 2 {                         // after each second step increas step by one
				step += 1
				iterStep = 0
			}
		}
		return x, y
	}
}

func getPoint(direct string, x, y int) (x1, y1 int) {
	x1, y1 = x, y
	switch direct {
	case "left":
		x1 -= 1
	case "down":
		y1 += 1
	case "right":
		x1 += 1
	case "top":
		y1 -= 1
	}
	return
}

func excludeByID(a []int, id int) []int {
	for k := 0; k < len(a); k++ {
		v := a[k]
		if v == id {
			a = removeFromInt(a, k)
			k--
		}
	}
	return a
}

func removeFromInt(a []int, i int) []int {
	a[i] = a[len(a)-1] // Copy last element to index i.
	a[len(a)-1] = 0    // Erase last element (write zero value).
	a = a[:len(a)-1]
	return a
}

func toInt(f float64) int {
	if f < 0 {
		return 0
	}
	return int(f)
}
