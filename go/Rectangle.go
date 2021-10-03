package main

import (
	"fmt"
	"math"
)

type Rectangle struct {
	start Point
	end   Point
}

//---------------------------------------------------------------------------------------
func (thisPt *Rectangle) SetStart(p *Point) {
	thisPt.start = *p
}

//---------------------------------------------------------------------------------------
func (thisPt *Rectangle) SetEnd(p *Point) {
	thisPt.end = *p
}

//---------------------------------------------------------------------------------------
func (thisPt *Rectangle) PointInside(point *Point) bool {
	checkV := func(start float64, end float64, val float64) bool {
		if val >= start && val <= end {
			return true
		}
		if val <= start && val >= end {
			return true
		}
		return false
	}
	return (checkV(thisPt.start.GetX(), thisPt.end.GetX(), point.GetX()) &&
		checkV(thisPt.start.GetY(), thisPt.end.GetY(), point.GetY()))
}

//---------------------------------------------------------------------------------------
func (thisPt *Rectangle) GetStart() *Point {
	return &thisPt.start
}

//---------------------------------------------------------------------------------------
func (thisPt *Rectangle) GetEnd() *Point {
	return &thisPt.end
}

//---------------------------------------------------------------------------------------
func (thisPt *Rectangle) GetLength() Point {
	p := Point{}
	p.SetX(thisPt.end.GetX() - thisPt.start.GetX())
	p.SetY(thisPt.end.GetY() - thisPt.start.GetY())
	return p
}

//---------------------------------------------------------------------------------------
func (thisPt *Rectangle) Add(other *Rectangle) {
	nStart := Point{}
	nEnd := Point{}
	nStart.SetX(math.Min(other.start.GetX(), thisPt.start.GetX()))
	nStart.SetY(math.Min(other.start.GetY(), thisPt.start.GetY()))
	nEnd.SetX(math.Max(other.end.GetX(), thisPt.end.GetX()))
	nEnd.SetY(math.Max(other.end.GetY(), thisPt.end.GetY()))
	thisPt.start = nStart
	thisPt.end = nEnd
}

//---------------------------------------------------------------------------------------
func (thisPt *Rectangle) GetWKTString() string {
	out := "POLYGON  (("
	out += fmt.Sprintf("%f %f,%f %f,%f %f,%f %f,%f %f",
		thisPt.start.GetX(), thisPt.start.GetY(),
		thisPt.end.GetX(), thisPt.start.GetY(),
		thisPt.end.GetX(), thisPt.end.GetY(),
		thisPt.start.GetX(), thisPt.end.GetY(),
		thisPt.start.GetX(), thisPt.start.GetY())

	out += "))"
	return out
}
