package main

import (
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"

	"gonum.org/v1/plot"
)

func main() {
	lines := ReadCSVData("/home/raha/go/src/ML/HW1/dataset.csv")
	points := reformatLinesToScatterPoints(lines)
	show(points)
}

// Read data from csv file and return lines
func ReadCSVData(path string) []string {
	f, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	dataAsBytes, err := ioutil.ReadAll(f)
	if err != nil {
		log.Println(err)
	}

	dataAsText := string(dataAsBytes)

	lines := strings.Split(dataAsText, "\n")[1:]

	return lines
}

func reformatLinesToScatterPoints(lines []string) plotter.XYs {
	pts := make(plotter.XYs, len(lines))

	for i := range lines {
		XY := strings.Split(lines[i], ",")

		pts[i].X, _ = strconv.ParseFloat(XY[0], 64)
		pts[i].Y, _ = strconv.ParseFloat(XY[1], 64)
	}

	return pts
}

func show(pts plotter.XYs) {
	p, err := plot.New()
	if err != nil {
		log.Fatal(err)
	}

	p.Title.Text = "ML dataset"
	p.X.Label.Text = "X"
	p.Y.Label.Text = "Y"

	s, err := plotter.NewScatter(pts)
	if err != nil {
		log.Fatal(err)
	}

	p.Add(s)
	p.Legend.Add("scatter", s)

	if err := p.Save(4*vg.Inch, 4*vg.Inch, "points.png"); err != nil {
		panic(err)
	}
}

type Point struct {
	X float64
	Y float64
}

func regression(dataset []Point, degree int, learningRate float64, steps int) {
	//In the beginning I want to set a stochastic line
	// Lets put m and b equal to one
	m := float64(1)
	b := float64(1)

	for i := 0; i < steps; i++ {
		deriv_m, deriv_b := derivative(dataset, m, b)
		m -= deriv_m * learningRate
		b -= deriv_b * learningRate
	}

	
}

// derivation with respect to m
func derivative(dataset []Point, m float64, b float64) (float64, float64) {
	n := float64(len(dataset))

	sigma_m := float64(0)
	sigma_b := float64(0)
	for _, d := range dataset {
		sigma_m += -2 * d.X * (d.Y - (m * d.X + b))
		sigma_b += -2 * (d.Y - (m * d.X + b))
	}

	return sigma_m / n, sigma_b / n
}