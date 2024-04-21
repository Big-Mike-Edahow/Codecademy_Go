/* main.go */

package main

import "fmt"

func fuelGauge(fuel int) {
	fmt.Printf("You have %d gallons of fuel left!\n\n", fuel)
}

func calculateFuel(planet string) int {
	var fuel int

	switch planet {
	case "Venus":
		fuel = 300000
	case "Mercury":
		fuel = 500000
	case "Mars":
		fuel = 700000
	case "Jupiter":
		fuel = 900000
	case "Saturn":
		fuel = 1100000
	case "Uranus":
		fuel = 1300000
	case "Neptune":
		fuel = 1500000
	}

	return fuel
}

func greetPlanet(planet string) {
	fmt.Print("You have traveled across the cosmos.\n")
	fmt.Printf("Welcome to %v!\n", planet)
}

func cantFly() {
	fmt.Println("We do not have the available fuel to fly there.")
}

func flyToPlanet(planet string, fuel int) int {
	var fuelCost int
	fuelRemaining := fuel

	fuelCost = calculateFuel(planet)
	if fuelRemaining >= fuelCost {
		greetPlanet(planet)
		fuelRemaining -= fuelCost
	} else {
		cantFly()
	}

	return fuelRemaining
}

func main() {
	fmt.Print("Welcome to the Interstellar Travel Agency.\n\n")
	var fuel int = 1000000
	fuelGauge(fuel)

	var planetChoice string
	fmt.Println("Which planet would you like to travel to?")
	fmt.Scan(&planetChoice)

	fuelNeeded := calculateFuel(planetChoice)
	fmt.Printf("\nYou will need %d gallons of fuel to travel to %v.\n\n", fuelNeeded, planetChoice)

	fuel = flyToPlanet(planetChoice, fuel)
	fmt.Print("\nYou have ", fuel, " gallons of fuel remaining.\n")
}

