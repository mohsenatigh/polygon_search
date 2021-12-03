package main

import (
	"flag"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func readFile(name string, geoMatrix *GeoMatrix) {

	buf, err := os.ReadFile(name)
	if err != nil {
		log.Fatalln(err)
	}

	lines := strings.Split(string(buf), "\n")

	readPoint := func(val string) Point {
		parts := strings.Split(val, ",")

		if len(parts) != 2 {
			log.Fatalf("invalid point value %s\n", val)
		}

		var err error
		var x float64
		var y float64

		if x, err = strconv.ParseFloat(parts[0], 64); err != nil {
			log.Fatalf("invalid point value %s\n", val)
		}

		if y, err = strconv.ParseFloat(parts[1], 64); err != nil {
			log.Fatalf("invalid point value %s\n", val)
		}

		return Point{x, y}
	}

	for _, l := range lines {
		points := []Point{}
		parts := strings.Split(l, "|")

		if len(parts) < 2 {
			continue
		}

		for i := 1; i < len(parts); i++ {
			if len(parts[i]) > 0 {
				points = append(points, readPoint(parts[i]))
			}
		}

		pol := Polygon{}
		pol.SetPoints(points)

		//convert id to number
		id, err := strconv.ParseFloat(parts[0], 64)
		if err != nil {
			log.Fatalln(err)
		}

		//
		geoMatrix.Add(pol, id)
	}

}

func main() {

	matrix := CreateMatrix()

	acc := GeoMatrixAccuracyMedium

	dataFile := flag.String("f", "../data.csv", "data file")
	lat := flag.Float64("a", 7.847786843776704, "lat")
	lng := flag.Float64("o", 47.99549694289439, "long")
	H := flag.Bool("H", false, "search with high accuracy")
	L := flag.Bool("L", false, "search with low accuracy")
	c := flag.Int("c", 10000000, "loop count")
	d := flag.Bool("d", false, "dump stat")
	useRandom := flag.Bool("R", false, "use random points")

	flag.Parse()

	if *H {
		acc = GeoMatrixAccuracyHigh
	} else if *L {
		acc = GeoMatrixAccuracyLow
	}

	readFile(*dataFile, matrix)
	matrix.Build()

	//test
	point := Point{*lat, *lng}

	//extract some statistics
	if *d {
		matrix.DumpStat()
		fmt.Printf(" Search area is  \n %s \n\n", matrix.GetMultiPolygonWKT(matrix.GetMatrixPoint(&point), &point))
		fmt.Printf(" Result area is  \n %s \n\n", matrix.GetResultWKT(&point, acc))
	}

	out := matrix.Query(&point, acc, -1)
	for i := range out {
		fmt.Println(out[i])
	}

	points := []Point{}
	for i := 0; i < *c; i++ {
		x := rand.Float64() / 10000
		y := rand.Float64() / 10000
		nPoint := Point{point.x + x, point.y + y}
		points = append(points, nPoint)
	}

	start := time.Now()
	for i := 0; i < *c; i++ {
		if *useRandom {
			matrix.Query(&points[i], acc, 1)
		} else {
			matrix.Query(&point, acc, 1)
		}
	}
	log.Printf("%v search takes %d microsec \n", *c, time.Since(start)/1000)

}
