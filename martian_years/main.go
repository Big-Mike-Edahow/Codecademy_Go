/* main.go */

package main
import "fmt"

func computeMarsYears(earthYears int) int {
	earthDays := earthYears * 365
	marsYears := earthDays / 687
	return marsYears
}

func main() {
	var myAge int
	fmt.Println("Greeting Earthling.\nPlease enter your age:")
	fmt.Scan(&myAge)

	myMartianAge := computeMarsYears(myAge)
	fmt.Printf("Your age in Martian years is %d.\n", myMartianAge)
}
