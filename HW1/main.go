package main

import (
	"fmt"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
	"io/ioutil"
	"log"
	"math"
	"os"
	"strconv"
	"strings"

	"gonum.org/v1/plot"
)

func main() {
	lines := ReadCSVData("/home/raha/go/src/ML/HW1/dataset.csv")
	//points := reformatLinesToScatterPoints(lines)
	//show(points)
	pts := reformatLinesToPoints(lines)
	fmt.Println(nonLinearRegression(pts, 5, 0.5, 3000, 0.01))
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

func reformatLinesToPoints(lines []string) []Point {
	pts := make([]Point, len(lines))

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

func linearRegression(dataset []Point, degree int, learningRate float64, steps int) {
	//In the beginning I want to set a stochastic line
	// Lets put m and b equal to one
	m := float64(1)
	b := float64(1)

	for i := 0; i < steps; i++ {
		deriv_m, deriv_b := derivative(dataset, m, b)
		m -= deriv_m * learningRate
		b -= deriv_b * learningRate
	}

	// Plot the line
	lineData := plotter.XYs{plotter.XY{X: 1, Y: m + b}, plotter.XY{X: 2, Y: 2 * m + b}}
	p, err := plot.New()
	if err != nil {
		panic(err)
	}
	p.Title.Text = "Points Example"
	p.X.Label.Text = "X"
	p.Y.Label.Text = "Y"

	l, err := plotter.NewLine(lineData)
	if err != nil {
		panic(err)
	}

	p.Add(l)
	p.Legend.Add("line", l)

	if err := p.Save(4*vg.Inch, 4*vg.Inch, "line.png"); err != nil {
		panic(err)
	}
}

// returns MSE error
func nonLinearRegression(dataset []Point, degree int, learningRate float64, steps int, lambda float64) float64 {
	coefficients := make([]float64, degree + 1)
	for i := range coefficients {
		coefficients[i] = 1
	}

	for i := 0; i < steps; i++ {
		deriv := nonLinearDerivative(dataset, coefficients, lambda)
		for j := range coefficients {
			coefficients[j] -= deriv[j] * learningRate
		}
	}

	// Plot the function
	p, err := plot.New()
	if err != nil {
		panic(err)
	}
	p.Title.Text = "functions"
	p.X.Label.Text = "X"
	p.Y.Label.Text = "Y"

	l := plotter.NewFunction(func(x float64) float64 {
		fixedPart := float64(0)
		for i, c := range coefficients {
			fixedPart += c * math.Pow(x, float64(i))
		}

		return fixedPart
	})

	p.Add(l)
	p.Legend.Add("function", l)

	if err := p.Save(4*vg.Inch, 4*vg.Inch, "nonlinear.png"); err != nil {
		panic(err)
	}

	// return MSE
	squredError := float64(0)
	for _, d := range dataset {
		estimatedY := float64(0)
		for i, c := range coefficients {
			estimatedY += c * math.Pow(d.X, float64(i))
		}
		squredError += math.Pow(d.Y - estimatedY, float64(2))
	}

	return squredError / float64(len(dataset))
}

func nonLinearDerivative(dataset []Point, coefficients []float64, lambda float64) []float64 {
	n := float64(len(dataset))
	sigma := make([]float64, len(coefficients))

	for i := range sigma {
		sigma[i] = float64(0)
	}

	// hard
	for _, d := range dataset {
		fixedPart := float64(0)
		for i, c := range coefficients {
			fixedPart += c * math.Pow(d.X, float64(i))
		}
		for i, s := range sigma {
			s += -2 * math.Pow(d.X, float64(i)) * (d.Y - fixedPart)
		}
	}

	res := make([]float64, len(coefficients))

	for i, s := range sigma{
		if i == 0 {
			res[i] = s/n
		}else {
			res[i] = (s + 2 * lambda * coefficients[i])/n
		}
	}

	return res
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