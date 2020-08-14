package main

import (
	"io/ioutil"
	"log"
	"os"
	"strings"

	"gonum.org/v1/plot"
)

func main() {

}

// I want to read the data from dataset and plot it
func show() {
	f, err := os.Open("dataset.csv")
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

}