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
	show()
}

// I want to read the data from dataset and plot it
func show() {
	f, err := os.Open("/home/raha/go/src/ML/HW1/dataset.csv")
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	dataAsBytes, err := ioutil.ReadAll(f)
	if err != nil {
		log.Println(err)
	}

	dataAsText := string(dataAsBytes)

	points := strings.Split(dataAsText, "\n")[1:]

	p, err := plot.New()
	if err != nil {
		log.Fatal(err)
	}

	p.Title.Text = "ML dataset"
	p.X.Label.Text = "X"
	p.Y.Label.Text = "Y"

	pts := reformatPoints(points)

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

func reformatPoints(points []string) plotter.XYs {
	pts := make(plotter.XYs, len(points))
	for i := range points {
		XY := strings.Split(points[i], ",")

		pts[i].X, _ = strconv.ParseFloat(XY[0], 64)
		pts[i].Y, _ = strconv.ParseFloat(XY[1], 64)
	}

	return pts
}