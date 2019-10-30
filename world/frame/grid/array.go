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
	far := make([]int, 0, 5)
	normal := make([]int, 0, 5)
	closest := make([]int, 0, 5)
	elx, ely := toInt((x)/g.cellSize), toInt((y)/g.cellSize)
	for cx := ltx; cx <= rdx; cx++ {
		for cy := lty; cy <= rdy; cy++ {
			if len(g.data) <= cx {
				continue
			}
			d := g.data[cx]
			if len(d) <= cy {
				continue
			}
			if cx == elx && cy == ely {
				closest = append(closest, d[cy]...)
			} else if cx == elx || cy == ely {
				normal = append(normal, d[cy]...)
			} else {
				far = append(far, d[cy]...)
			}
		}
	}
	obj := append(closest, normal...)
	obj = append(obj, far...)
	return excludeByID(obj, exclude)
}

func excludeByID(a []int, id int) []int {
	prev := -1
	for k := 0; k < len(a); k++ {
		v := a[k]
		if v == id || v == prev {
			a = removeFromInt(a, k)
			k--
		}
		prev = v
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
