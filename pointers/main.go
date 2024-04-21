/* main.go */

package main

import "fmt"

func main() {
	var minutes int = 525600
	var pointerForInt *int = &minutes
	fmt.Printf("The address of minutes is\t%v\n", pointerForInt)

	star := "Polaris"
	starAddress := &star
	fmt.Print("The address of star is\t\t", starAddress, "\n\n")

	lyrics := "Moments so dear"
	pointerForStr := &lyrics
	fmt.Print("The lyrics are: \"", lyrics, "\"\n")

	*pointerForStr = "Journeys to plan"
	fmt.Printf("The lyrics are: \"%v\"\n", lyrics)
}
