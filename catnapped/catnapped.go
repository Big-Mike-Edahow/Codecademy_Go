/* catnapped.go */

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func getRandomElement(slice []string) string {
	length := len(slice)
	elementIndex := rand.Intn(length)
	return slice[elementIndex]
}

func main() {
	rand.Seed(time.Now().UnixNano())

	guests := [5]string{"Anna", "Jose", "Louis", "Tina", "Raymond"}
	objects := [4]string{"Toy Chest", "Vase", "Box", "Cat Tree"}
	fmt.Println("Guests:", guests)
	fmt.Println("Objects:", objects)

	culprit := getRandomElement(guests[:])
	secretObject := getRandomElement(objects[:])
	fmt.Print("\n", culprit, " hid the cat in the ", secretObject, ".\n")
}
