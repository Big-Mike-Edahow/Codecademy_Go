/* main.go */

package main

import "fmt"

func main() {
	gradeBook := map[string]string{
		"Steve": "A",
		"Jared": "B+",
		"Mike":  "C",
		"Gabe":  "A",
	}

	for key, value := range gradeBook {
		fmt.Println("Name:", key, "\tGrade:", value)
	}
}
