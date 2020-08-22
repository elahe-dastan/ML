package model

import "math"

type Diagram struct {
	Mean     float64
	Variance float64
}

func CalculateDiagram(wines []Wine) []Diagram {
	diagrams := make([]Diagram, 16)

	for i := 0; i < len(diagrams); i++ {
		sum := float64(0)
		for _, wine := range wines {
			sum += wine.Info[i]
		}
		diagrams[i].Mean = sum / float64(len(wines))

		sd := float64(0)
		for _, wine := range wines {
			sd += math.Pow(wine.Info[i] - diagrams[i].Mean, 2)
		}
		diagrams[i].Variance = sd / float64(len(wines))
	}

	return diagrams
}