/* main.go */

package main

import (
	"fmt"
)

/* Global Variable */
var number int

func count() {
	var input int
	fmt.Print("Number of guests: ")
	fmt.Scan(&input)
	number = number + input
	fmt.Println("Total guests:", number)
}

func main() {
	for {
		count()
	}
}

