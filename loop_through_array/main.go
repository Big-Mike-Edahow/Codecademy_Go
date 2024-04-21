/* main.go */

package main

import (
	"fmt"
)

func main() {
	menu := []string{"Hamburgers", "Cheeseburgers", "Fries"}

	fmt.Println("The menu:")
	for number, item := range menu {
		fmt.Println(number, ":", item)
	}

	orders := map[string]string{
		"John":   "Cheeseburgers",
		"Janet":  "Hamburgers",
		"Jordan": "Fries",
	}
	// A late order
	orders["James"] = "Chicken Sandwich"

	fmt.Println("\nThe friend's orders:")
	for name, order := range orders {
		fmt.Println(name, "wants some", order)
	}
}
