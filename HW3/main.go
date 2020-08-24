package main

import (
	"ML/HW3/model"
	"ML/data"
	"fmt"
	"math"
	"strconv"
	"strings"
)

func main() {
	lines := data.ReadCSVData("/home/raha/go/src/ML/HW3/wine.data")
	wines := reformatLineToWine(lines)

	class1 := make([]model.Wine, 0)
	class2 := make([]model.Wine, 0)
	class3 := make([]model.Wine, 0)

	for _, wine := range wines {
		if wine.Class == 1 {
			class1 = append(class1, wine)
		}
		if wine.Class == 2 {
			class2 = append(class2, wine)
		}
		if wine.Class == 3 {
			class3 = append(class3, wine)
		}
	}

	// for Test
	testWine := model.Wine{
		Info:            []float64{14.23,1.71,2.43,15.6,127,2.8,3.06,.28,2.29,5.64,1.04,3.92,1065},
	}

	diagrams1 := model.CalculateDiagram(class1)
	diagrams2 := model.CalculateDiagram(class2)
	//diagrams3 := model.CalculateDiagram(class3)

	probability1 := math.Log(float64(len(class1) / len(wines)))

	for i, diagram := range diagrams1 {
		probability1 += math.Log(likelihood(diagram.Mean, diagram.Variance, testWine.Info[i]))
	}

	probability2 := math.Log(float64(len(class2) / len(wines)))

	for i, diagram := range diagrams2 {
		probability2 += math.Log(likelihood(diagram.Mean, diagram.Variance, testWine.Info[i]))
	}

	if probability1 > probability2 {
		fmt.Println(probability1)
	}else {
		fmt.Println(probability2)
	}
}

func reformatLineToWine(lines []string) []model.Wine {
	wines := make([]model.Wine, len(lines))

	for i := range lines {
		d := strings.Split(lines[i], ",")

		wines[i].Class, _ = strconv.Atoi(d[0])
		wines[i].Alcohol, _ = strconv.ParseFloat(d[1], 64)
		wines[i].Malic, _ = strconv.ParseFloat(d[2], 64)
		wines[i].Acid, _ = strconv.ParseFloat(d[3], 64)
		wines[i].Ash, _ = strconv.ParseFloat(d[4], 64)
		wines[i].AlcalinityOfAsh, _ = strconv.ParseFloat(d[5], 64)
		wines[i].Magnesium, _ = strconv.ParseFloat(d[6], 64)
		wines[i].TotalPhenols, _ = strconv.ParseFloat(d[7], 64)
		wines[i].Flavanoids, _ = strconv.ParseFloat(d[8], 64)
		wines[i].Nonflavanoid, _ = strconv.ParseFloat(d[9], 64)
		wines[i].Phenols, _ = strconv.ParseFloat(d[10], 64)
		wines[i].Proanthocyanins, _ = strconv.ParseFloat(d[11], 64)
		wines[i].ColorIntensity, _ = strconv.ParseFloat(d[12], 64)
		wines[i].Hue, _ = strconv.ParseFloat(d[13], 64)
		wines[i].DilutedWines, _ = strconv.ParseFloat(d[14], 64)
		wines[i].Proline, _ = strconv.ParseFloat(d[15], 64)

		wines[i].Info = []float64{wines[i].Alcohol, wines[i].Malic, wines[i].Acid, wines[i].Ash, wines[i].AlcalinityOfAsh,
			wines[i].Magnesium, wines[i].TotalPhenols, wines[i].Flavanoids, wines[i].Nonflavanoid, wines[i].Phenols,
			wines[i].Proanthocyanins, wines[i].ColorIntensity, wines[i].Hue, wines[i].DilutedWines, wines[i].Proline}
	}

	return wines
}

func likelihood(mean float64, variance float64, x float64) float64 {
	return (1 / math.Sqrt(2 * math.Pi * variance)) * math.Pow(math.E, -1 * math.Pow(x - mean, 2) / (2 * variance))
}
