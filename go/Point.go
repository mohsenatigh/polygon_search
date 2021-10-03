package main

import "fmt"

type Point struct {
	x float64
	y float64
}

//---------------------------------------------------------------------------------------
func (thisPt *Point) Equal(other *Point) bool {
	return (thisPt.x == other.x && thisPt.y == other.y)
}

//---------------------------------------------------------------------------------------
func (thisPt *Point) SetX(x float64) {
	thisPt.x = x
}

//---------------------------------------------------------------------------------------
func (thisPt *Point) SetY(y float64) {
	thisPt.y = y
}

//---------------------------------------------------------------------------------------
func (thisPt *Point) GetX() float64 {
	return thisPt.x
}

//---------------------------------------------------------------------------------------
func (thisPt *Point) GetY() float64 {
	return thisPt.y
}

//---------------------------------------------------------------------------------------
func (thisPt *Point) SetXL(x uint64) {
	thisPt.x = float64(x)
}

//---------------------------------------------------------------------------------------
func (thisPt *Point) SetYL(y uint64) {
	thisPt.y = float64(y)
}

//---------------------------------------------------------------------------------------
func (thisPt *Point) Gethash() uint64 {
	xl := uint64(thisPt.x)
	return (xl<<32 | uint64(thisPt.y))
}

//---------------------------------------------------------------------------------------
func (thisPt *Point) GetWKTString() string {
	return fmt.Sprintf("POINT(%f %f)", thisPt.x, thisPt.y)
}
