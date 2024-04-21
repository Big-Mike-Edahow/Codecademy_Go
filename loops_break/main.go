/* main.go */

package main

import "fmt"

func main() {
	animals := []string{"Cat", "Dog", "Fish", "Turtle"}

	for index := 0; index < len(animals); index++ {
		if animals[index] == "Dog" {
			fmt.Print("At index ",  index, " is ", animals[index], ".\n")
			fmt.Println("Found the perfect animal!")
			break // Stop searching the array
		}
	}
}
