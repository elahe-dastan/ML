package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	// reading data
	f, err := os.Open("/home/raha/go/src/ML/HW3/q2/train-images.idx3-ubyte")
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	dataAsBytes, err := ioutil.ReadAll(f)
	if err != nil {
		log.Println(err)
	}

	fmt.Println(dataAsBytes[3])

	//data := dataAsBytes[4:8]
	//var a int32
	//err = binary.Read(bytes.NewReader(data), binary.BigEndian, &a)
	//fmt.Println(err)
	//fmt.Println(a)
	//
	//data1 := dataAsBytes[8:12]
	//var a1 int32
	//err = binary.Read(bytes.NewReader(data1), binary.BigEndian, &a1)
	//fmt.Println(err)
	//fmt.Println(a1)
	//
	//data2 := dataAsBytes[12:16]
	//var a2 int32
	//err = binary.Read(bytes.NewReader(data2), binary.BigEndian, &a2)
	//fmt.Println(err)
	//fmt.Println(a2)
	//
	//fmt.Println(a1 * a2 * a + 4 + 12)
	//
	//fmt.Println(len(dataAsBytes))
	//fmt.Println(int32(dataAsBytes[4:8]))
	//dataAsText := string(dataAsBytes)
	// change it to one vs all
	// logistic regression
	// probability
}
