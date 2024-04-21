// main.go

package main

import "fmt"

type Cake struct {
	variety string
	weight  string
}

func main() {

	var cakes = []Cake{{"Chocolate", "1 lb"}, {"Carrot", "1/2 lb"}, {"Ice Cream", "2 lbs"}}

	fmt.Println(cakes[0].weight)
	cakes[1].weight = "2.5 lbs"
	fmt.Println(cakes)
}
