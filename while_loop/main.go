/* main.go */

package main

import (
	"fmt"
)

func ask() int {
	var input int
	fmt.Print("I am thinking of a number between 1-10: ")

	fmt.Scan(&input)
	return input
}

func main() {
	var guess int

	for guess != 7 {
		guess = ask()
		if guess == 7 {
			fmt.Println("You are correct! You may go through to the treasure!")
		} else if guess < 7 {
			fmt.Println("Too low...")
		} else {
			fmt.Println("Too high...")
		}
	}

}
