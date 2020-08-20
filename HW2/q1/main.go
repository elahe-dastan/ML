package main

import (
	"ML/HW2/q1/model"
	"ML/data"
	"math"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

func main() {
	lines := data.ReadCSVData("/home/raha/go/src/ML/HW2/q1/processed.cleveland.data")
	persons := reformatLineToPerson(lines)

	rand.Seed(time.Now().Unix())
	rand.Shuffle(len(persons), func(i, j int) {
		persons[i], persons[j] = persons[j], persons[i]
	})

	fraction := len(persons) * 2 / 3
	trainigData := persons[:fraction]
	testData := persons[fraction:]
	
}

func knn(data []model.Person, query model.Person, k int) int {
	persons := make([]model.Person, len(data))
	distances := make([]float64, len(data))

	for i, d := range data {
		distance := EuclideanDistance(query.Info, d.Info)
		persons[i] = d
		distances[i] = distance
	}

	InsertionSort(distances, persons)

	firstK := persons[:k]

	return mode(firstK)
}

func EuclideanDistance(firstPoint []float64, secondPoint []float64) float64 {
	sumSquaredDistance := float64(0)

	for i := range firstPoint {
		sumSquaredDistance += math.Pow(firstPoint[i]-secondPoint[i], 2)
	}

	return math.Sqrt(sumSquaredDistance)
}

func reformatLineToPerson(lines []string) []model.Person {
	persons := make([]model.Person, len(lines))

	for i := range lines {
		d := strings.Split(lines[i], ",")

		persons[i].Age, _ = strconv.ParseFloat(d[0], 64)
		persons[i].Sex, _ = strconv.ParseFloat(d[1], 64)
		persons[i].CP, _ = strconv.ParseFloat(d[2], 64)
		persons[i].Trestbps, _ = strconv.ParseFloat(d[3], 64)
		persons[i].Chol, _ = strconv.ParseFloat(d[4], 64)
		persons[i].Fbs, _ = strconv.ParseFloat(d[5], 64)
		persons[i].Restecg, _ = strconv.ParseFloat(d[6], 64)
		persons[i].Thalach, _ = strconv.ParseFloat(d[7], 64)
		persons[i].Exang, _ = strconv.ParseFloat(d[8], 64)
		persons[i].Oldpeak, _ = strconv.ParseFloat(d[9], 64)
		persons[i].Slope, _ = strconv.ParseFloat(d[10], 64)
		persons[i].Ca, _ = strconv.ParseFloat(d[11], 64)
		persons[i].Thal, _ = strconv.ParseFloat(d[12], 64)
		persons[i].Num, _ = strconv.Atoi(d[13])

		persons[i].Info = []float64{persons[i].Age, persons[i].Sex, persons[i].CP, persons[i].Trestbps, persons[i].Chol,
			persons[i].Fbs, persons[i].Restecg, persons[i].Thalach, persons[i].Exang, persons[i].Oldpeak, persons[i].Slope,
			persons[i].Ca, persons[i].Thal}
	}

	return persons
}

func InsertionSort(items []float64, persons []model.Person) {

	L := len(items)

	for i := 1; i < L; i++ {

		j := i
		for j > 0 && items[j] < items[j-1] {
			items[j], items[j-1] = items[j-1], items[j]
			persons[j], persons[j-1] = persons[j-1], persons[j]
			j -= 1
		}

	}

}

func mode (persons []model.Person) int {
	labels := make([]int, 5)

	for _, d := range persons {
		labels[d.Num]++
	}

	max := 0

	for _, v := range labels {
		if v > max {
			max = v
		}
	}

	return max
}