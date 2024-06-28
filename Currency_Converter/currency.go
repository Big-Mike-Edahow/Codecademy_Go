/* currency.go */

package main

import (
	"fmt"
)

func main() {
	currencies := map[string]float32{"EUR": 0.95, "JPY": 130}

	var dollarAmount float32
	var currency string

	fmt.Println("Please enter a dollar amount:")
	fmt.Scan(&dollarAmount)

	if dollarAmount == 0 {
		fmt.Println("Invalid dollar amount.")
	} else {
		fmt.Println("Please enter the currency: (EUR or JPY)")
		fmt.Scan(&currency)

		rate, valid := currencies[currency]
		if valid == true {
			fmt.Println(dollarAmount, "USD converts to", rate*dollarAmount, currency)
		} else {
			fmt.Println("Cannot convert USD to", currency)
		}
	}
}
