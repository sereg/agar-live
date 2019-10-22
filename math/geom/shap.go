package geom

func Intersects( cx,  cy,  radius,  left,  top,  right,  bottom float64) bool{
	closestX := 0.0
	if cx < left {
		closestX = left
	} else if cx > right {
		closestX = right
	}else {
		closestX = cx
	}
	closestY := 0.0
	if cy < top {
		closestY = top
	} else if cy > bottom {
		closestY = bottom
	}else {
		closestY = cy
	}
	dx := closestX - cx
	dy := closestY - cy
	return ( dx * dx + dy * dy ) <= radius * radius
}
