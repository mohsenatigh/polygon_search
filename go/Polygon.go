package main

import (
	"fmt"
	"math"
)

type Polygon struct {
	points []Point
	bBox   *Rectangle
}

//---------------------------------------------------------------------------------------
func (thisPt *Polygon) PointIsInside(point *Point) bool {

	count := 0

	if thisPt.bBox != nil && !thisPt.bBox.PointInside(point) {
		return false
	}

	size := len(thisPt.points)
	for i := 1; i < size; i++ {
		p1 := &thisPt.points[i-1]
		p2 := &thisPt.points[i]

		if p1.GetX() < point.GetX() && p2.GetX() < point.GetX() {
			continue
		}

		if p1.GetY() <= point.GetY() && p2.GetY() >= point.GetY() {
			count++
		} else if p1.GetY() >= point.GetY() && p2.GetY() <= point.GetY() {
			count++
		}
	}
	return ((count % 2) != 0)
}

//---------------------------------------------------------------------------------------
func (thisPt *Polygon) SetPoints(pointList []Point) {
	thisPt.points = pointList
}

//---------------------------------------------------------------------------------------

func (thisPt *Polygon) GetBoundingBox() *Rectangle {
	if thisPt.bBox != nil {
		return thisPt.bBox
	}

	bBox := new(Rectangle)

	start := Point{}
	end := Point{}

	start.SetX(math.MaxFloat64)
	start.SetY(math.MaxFloat64)
	end.SetX(0)
	end.SetY(0)

	for i := range thisPt.points {
		p := &thisPt.points[i]

		if p.GetX() < start.GetX() {
			start.SetX(p.GetX())
		}

		if p.GetY() < start.GetY() {
			start.SetY(p.GetY())
		}

		if p.GetX() > end.GetX() {
			end.SetX(p.GetX())
		}

		if p.GetY() > end.GetY() {
			end.SetY(p.GetY())
		}
	}

	bBox.SetEnd(&end)
	bBox.SetStart(&start)

	thisPt.bBox = bBox
	return thisPt.bBox
}

//---------------------------------------------------------------------------------------
func (thisPt *Polygon) GetWKTString() string {
	out := "POLYGON  (("
	for i, p := range thisPt.points {
		if i > 0 {
			out += ","
		}
		out += fmt.Sprintf("%f %f", p.GetX(), p.GetY())
	}
	out += "))"
	return out
}

//---------------------------------------------------------------------------------------
