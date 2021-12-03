package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/mitchellh/mapstructure"
)

type Point []float64

type Points []Point
type MultiPoints []Points
type MultiMultiPoints []MultiPoints

type Polygon struct {
	Type        string   `json:"type"`
	Coordinates []Points `json:"coordinates"`
}

type MultiPolygon struct {
	Type        string        `json:"type"`
	Coordinates []MultiPoints `json:"coordinates"`
}

type Properties struct {
	Plz       string  `json:"plz"`
	Note      string  `json:"note"`
	Qkm       float64 `json:"qkm"`
	Einwohner uint64  `json:"einwohner"`
}

func main() {
	in := flag.String("i", "/root/projects/polygon_search/data/plz-5stellig.geojson", "in file")
	out := flag.String("o", "/root/projects/polygon_search/data/data.csv", "out file")

	flag.Parse()

	outF, err := os.Create(*out)
	if err != nil {
		log.Fatalln(err)
	}
	defer outF.Close()

	//read file
	buffer, err := os.ReadFile(*in)
	if err != nil {
		log.Fatalln(err)
	}

	//load collection
	dataMap := make(map[string]interface{})

	//collection := FeatureCollection{}
	if err := json.Unmarshal(buffer, &dataMap); err != nil {
		log.Fatalln(err)
	}

	features := dataMap["features"].([]interface{})
	for _, obj := range features {
		objM := obj.(map[string]interface{})
		geometry := objM["geometry"].(map[string]interface{})
		properties := objM["properties"].(map[string]interface{})
		if geometry["type"] == "Polygon" {
			geo := Polygon{}
			if err := mapstructure.Decode(geometry, &geo); err != nil {
				log.Fatalln(err)
			}

			result := fmt.Sprintf("%v", properties["plz"])
			for _, c := range geo.Coordinates[0] {
				result += fmt.Sprintf("|%v,%v", c[0], c[1])
			}
			outF.WriteString(result + "\n")

		} else if geometry["type"] == "MultiPolygon" {
			geo := MultiPolygon{}
			if err := mapstructure.Decode(geometry, &geo); err != nil {
				log.Fatalln(err)
			}

			for _, c := range geo.Coordinates {
				result := fmt.Sprintf("%v", properties["plz"])
				for _, p := range c[0] {
					result += fmt.Sprintf("|%v,%v", p[0], p[1])
				}
				outF.WriteString(result + "\n")
			}
		}
	}

}
