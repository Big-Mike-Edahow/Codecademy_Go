/* main.go */

package main

import (
	"fmt"
	"time"
)

// Add "string" as the return type of this function
func isItLateInNewYork() string {
	var lateMessage string
	t := time.Now()
	tz, _ := time.LoadLocation("America/New_York")
	nyHour := t.In(tz).Hour()
	
	if nyHour < 5 {
		lateMessage = fmt.Sprintf("Goodness it is late, it is %v:00.", nyHour)
	} else if nyHour < 16 {
		lateMessage = fmt.Sprintf("It's not late at all, it's %v:00.", nyHour)
	} else if nyHour < 19 {
		lateMessage = fmt.Sprintf("I guess it's getting kind of late, it's %v:00.", nyHour)
	} else {
		lateMessage = fmt.Sprintf("It's late, it's %v:00.", nyHour)
	}

	// Return the string lateMessage
	return lateMessage
}

func main() {
	var nyLate string = isItLateInNewYork()
	fmt.Println(nyLate)
}

