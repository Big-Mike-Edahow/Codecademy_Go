/* main.go */

package main

import "fmt"

func addHundred(numPtr *int) {
	// Dereference the address.
	*numPtr += 100
}

func brainwash(saying *string) {
	// Dereference saying below:
	*saying = "Beep Boop."
}

func main() {
	x := 1
	fmt.Printf("X is: %d\n", x)
	addHundred(&x)
	fmt.Print("X is now: ", x, "\n\n")

	greeting := "Hello there!"
	fmt.Println("Greeting is:", greeting)
	brainwash(&greeting)
	fmt.Println("Greeting is now:", greeting)
}

