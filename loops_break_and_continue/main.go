/* main.go */

package main

import (
	"fmt"
)

func main() {
	for count := 0; count < 20; count++ {
		if count == 8 {
			continue	// 8 will not be printed.
		}
		if count == 15 {
			break		// Loop will exit. 15 will not be printed.
		}
		fmt.Println(count)
	}
}
