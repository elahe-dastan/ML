package main

import (
	"ML/HW2/q1/model"
	"ML/data"
	"strconv"
	"strings"
)

func main() {
	lines := data.ReadCSVData("/home/raha/go/src/ML/HW2/q1/processed.cleveland.data")

}

func knn(data []model.Person,query model.Person, k int) {

}

func reformatLineToPerson(lines []string) []model.Person {
	persons := make([]model.Person, len(lines))

	for i := range lines {
		d := strings.Split(lines[i], ",")

		persons[i]., _ = strconv.ParseFloat(XY[0], 64)
		pts[i].Y, _ = strconv.ParseFloat(XY[1], 64)
	}

	return pts
}

