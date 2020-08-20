package data

import (
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

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

func Show(pts plotter.XYs) {
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

func ReformatLinesToScatterPoints(lines []string) plotter.XYs {
	pts := make(plotter.XYs, len(lines))

	for i := range lines {
		XY := strings.Split(lines[i], ",")

		pts[i].X, _ = strconv.ParseFloat(XY[0], 64)
		pts[i].Y, _ = strconv.ParseFloat(XY[1], 64)
	}

	return pts
}