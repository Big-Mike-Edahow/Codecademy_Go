/* main.go */

package main

import (
	"fmt"
	"math"
)

func main() {
	var num float64

	fmt.Println("Please enter a floating point number:")
	fmt.Scan(&num)

	fmt.Printf("Logarithm of %.2f is %f.\n", num, math.Log(num))
	fmt.Printf("Squareroot of %.2f is %f.\n", num, math.Sqrt(num))
	fmt.Printf("Tangent of %.2f is %f.\n", num, math.Tan(num))
}

