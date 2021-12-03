package main

import (
	"fmt"
	"math"
)

type MatrixPoint [2]uint64

//---------------------------------------------------------------------------------------
const ROW = 1000

//---------------------------------------------------------------------------------------
type GeoMatrixData struct {
	polygon Polygon
	data    float64
}

//---------------------------------------------------------------------------------------
const (
	GeoMatrixAccuracyHigh   = 0
	GeoMatrixAccuracyMedium = 1
	GeoMatrixAccuracyLow    = 2
)

//---------------------------------------------------------------------------------------
type GeoMatrix struct {
	bBox  Rectangle
	pMap  map[MatrixPoint][]*GeoMatrixData
	items []GeoMatrixData
}

//---------------------------------------------------------------------------------------
func (thisPt *GeoMatrix) GetMultiPolygonWKT(point MatrixPoint, extPoint *Point) string {
	polyList := thisPt.pMap[point]
	if polyList == nil {
		return ""
	}

	out := "GEOMETRYCOLLECTION("
	for i, p := range polyList {
		if i > 0 {
			out += ","
		}
		out += p.polygon.GetWKTString()
	}

	if extPoint != nil {
		out += "," + extPoint.GetWKTString()
	}

	out += ")"
	return out
}

//---------------------------------------------------------------------------------------
func (thisPt *GeoMatrix) GetPointMosaic(point MatrixPoint) string {
	polyList := thisPt.pMap[point]
	if polyList == nil {
		return ""
	}

	tmpMap := make(map[MatrixPoint]bool)

	for _, p := range polyList {
		rect := p.polygon.GetBoundingBox()
		start := thisPt.GetMatrixPoint(&rect.start)
		end := thisPt.GetMatrixPoint(&rect.end)
		for x := start[0]; x <= end[0]; x++ {
			for y := start[1]; y <= end[1]; y++ {
				tmpMap[MatrixPoint{x, y}] = true
			}
		}
	}

	cnt := 0
	out := "GEOMETRYCOLLECTION("
	for k, _ := range tmpMap {
		if cnt > 0 {
			out += ","
		}
		cnt++
		rect := thisPt.GetCellLatLong(k)
		out += rect.GetWKTString()
	}
	out += ")"
	return out
}

//---------------------------------------------------------------------------------------
func (thisPt *GeoMatrix) GetMatrixPoint(in *Point) MatrixPoint {

	ext := thisPt.bBox.GetLength()

	slotX := (ext.GetX() / ROW)
	slotY := (ext.GetY() / ROW)

	x := ((in.GetX() - thisPt.bBox.GetStart().GetX()) / slotX)
	y := ((in.GetY() - thisPt.bBox.GetStart().GetY()) / slotY)

	x = math.Min(x, ROW-1)
	y = math.Min(y, ROW-1)

	return MatrixPoint{uint64(x), uint64(y)}
}

//---------------------------------------------------------------------------------------
func (thisPt *GeoMatrix) CoverArea(start MatrixPoint, end MatrixPoint, itemInfo *GeoMatrixData) {
	for x := start[0]; x <= end[0]; x++ {
		for y := start[1]; y <= end[1]; y++ {
			point := MatrixPoint{x, y}
			thisPt.pMap[point] = append(thisPt.pMap[point], itemInfo)
		}
	}
}

//---------------------------------------------------------------------------------------
func (thisPt *GeoMatrix) Add(poly Polygon, data float64) {
	rect := poly.GetBoundingBox()

	if len(thisPt.items) == 0 {
		thisPt.bBox = *rect
	} else {
		thisPt.bBox.Add(rect)
	}

	thisPt.items = append(thisPt.items, GeoMatrixData{poly, data})
}

//---------------------------------------------------------------------------------------
func (thisPt *GeoMatrix) Build() {
	for i := range thisPt.items {
		rect := thisPt.items[i].polygon.GetBoundingBox()
		start := thisPt.GetMatrixPoint(rect.GetStart())
		end := thisPt.GetMatrixPoint(rect.GetEnd())
		thisPt.CoverArea(start, end, &thisPt.items[i])
	}
}

//---------------------------------------------------------------------------------------

func (thisPt *GeoMatrix) Query(point *Point, acc int, maxCount int) []float64 {
	mPoint := thisPt.GetMatrixPoint(point)
	items := thisPt.pMap[mPoint]
	out := []float64{}

	if items == nil {
		return nil
	}

	if len(items) == 1 {
		out = append(out, items[0].data)
		return out
	}

	for i := range items {

		if len(out) == maxCount {
			return out
		}

		switch acc {
		case GeoMatrixAccuracyHigh:
			{
				if items[i].polygon.PointIsInside(point) {
					out = append(out, items[i].data)
					return out
				}
			}
		case GeoMatrixAccuracyMedium:
			{
				if items[i].polygon.GetBoundingBox().PointInside(point) {
					out = append(out, items[i].data)
				}
			}
		default:
			{
				out = append(out, items[i].data)
			}

		}
	}
	return out
}

//---------------------------------------------------------------------------------------
func (thisPt *GeoMatrix) GetCellLatLong(cell MatrixPoint) Rectangle {
	ext := thisPt.bBox.GetLength()

	slotX := (ext.GetX() / ROW)
	slotY := (ext.GetY() / ROW)

	start := Point{thisPt.bBox.start.x + (float64(cell[0]) * slotX), thisPt.bBox.start.y + (float64(cell[1]) * slotY)}
	end := Point{start.GetX() + slotX, start.GetY() + slotY}

	return Rectangle{start, end}
}

//---------------------------------------------------------------------------------------
func (thisPt *GeoMatrix) DumpStat() {
	maxPoint := MatrixPoint{}
	sampleO1Point := MatrixPoint{}
	maxLen := 0
	oOnePoints := 0

	//extract some info
	for k, p := range thisPt.pMap {
		if len(p) > maxLen {
			maxPoint = k
			maxLen = len(p)
		}

		if len(p) == 1 {
			oOnePoints++
			sampleO1Point = k
		}
	}

	maxRect := thisPt.GetCellLatLong(maxPoint)
	o1Rect := thisPt.GetCellLatLong(sampleO1Point)
	oOnePoints += (ROW * ROW) - len(thisPt.pMap)

	data :=
		`
Total Cell 				: %d 
Max M 					: %d POL
Max Cell Info 			: %s
Max Cell Pol 			: %s
Max Cell Mosaic 		: %s
O(1) 					: %d
Sample O(1) Rect		: %s
Sample O(1) Pol 		: %s
Sample O(1) Mosaic 		: %s
`

	fmt.Printf(data,
		len(thisPt.pMap),
		maxLen,
		maxRect.GetWKTString(),
		thisPt.GetMultiPolygonWKT(maxPoint, nil),
		thisPt.GetPointMosaic(maxPoint),
		oOnePoints,
		o1Rect.GetWKTString(),
		thisPt.GetMultiPolygonWKT(sampleO1Point, nil),
		thisPt.GetPointMosaic(sampleO1Point),
	)
}

//---------------------------------------------------------------------------------------
func (thisPt *GeoMatrix) GetResultWKT(pt *Point, acc int) string {
	mPoint := thisPt.GetMatrixPoint(pt)
	items := thisPt.pMap[mPoint]
	matchList := []Polygon{}

	if items == nil {
		return ""
	}

	for i := range items {
		item := items[i]
		if acc == GeoMatrixAccuracyHigh {
			if item.polygon.PointIsInside(pt) {
				matchList = append(matchList, item.polygon)
				break
			}
		} else if acc == GeoMatrixAccuracyMedium {
			if item.polygon.GetBoundingBox().PointInside(pt) {
				matchList = append(matchList, item.polygon)
			}
		} else {
			matchList = append(matchList, item.polygon)
		}
	}

	out := "GEOMETRYCOLLECTION("
	for i, p := range matchList {
		if i > 0 {
			out += ","
		}
		out += p.GetWKTString()
	}
	out += "," + pt.GetWKTString()
	out += ")"

	return out
}

//---------------------------------------------------------------------------------------

func CreateMatrix() *GeoMatrix {
	matrix := new(GeoMatrix)
	matrix.pMap = make(map[MatrixPoint][]*GeoMatrixData)
	return matrix
}
