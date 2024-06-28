/* order.go */

package main

import (
	"fmt"
)

func askOrder() string {
	var input string
	fmt.Print("What would you like to eat: ")
	// Get the input from the user
	fmt.Scan(&input)
	return input
}

func contains(menu []string, order string) bool {
	for _, item := range menu {
		if order == item {
			return true
		}
	}
	return false
}

func main() {
	fastfoodMenu := []string{"Burger", "Nuggets", "Fries", "Cola"}
	fmt.Printf("The fast food menu has these items:\n%v\n", fastfoodMenu)

  var total int
  var order string
	for order != "quit" {
		order = askOrder()
		if contains(fastfoodMenu, order) {
			total += 2
		} else if order == "quit" {
			break
		} else {
			fmt.Println("This item is not on the menu.")
		}
	}
	fmt.Printf("\nThe total for the order is: %v\n\n", total)
	for amount := total; amount > 0; amount -= 2 {
		fmt.Println("The amount owed is:", amount)
		fmt.Println("Handing over a $2 bill.")
	}
	fmt.Println("Thank you for your order.")
}
