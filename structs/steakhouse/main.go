// main.go

package main

import "fmt"

type Restaurant struct {
	name             string
	typeOfRestaurant string
	yearEstablished  int
}

func main() {

	// Add your code here.
	restaurant := Restaurant{"Codecademy Steakhouse", "Japanese", 2011}
	fmt.Println(restaurant)

	fmt.Printf("Name: %v\nType: %v\nEstablished: %v\n", restaurant.name, restaurant.typeOfRestaurant,
		restaurant.yearEstablished)

	restaurant.name = "Skillsoft Steakhouse"
	restaurant.yearEstablished = 2022
	
	fmt.Printf("Name: %v\nType: %v\nEstablished: %v\n", restaurant.name, restaurant.typeOfRestaurant,
		restaurant.yearEstablished)
}
