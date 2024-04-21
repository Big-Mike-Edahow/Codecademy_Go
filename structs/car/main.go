// main.go

package main

import "fmt"

type Car struct {
	make  string
	model string
	year  int
}

func main() {
	var car = Car{
		make:  "Ford",
		model: "Transit Connect",
		year:  2015,
	}

	fmt.Println("Make:\t", car.make)
	fmt.Println("Model:\t", car.model)
	fmt.Println("Year:\t", car.year)
}
