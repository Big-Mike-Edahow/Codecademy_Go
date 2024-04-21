/* main.go */

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	isHeistOn := true
	eludedGuards := rand.Intn(100)
	if eludedGuards >= 50 {
		fmt.Println("Looks like you've managed to make it past the guards. Good job, but remember, this is the first step.")
	} else {
		isHeistOn = false
		fmt.Println("\nPlan a better disguise next time?")
	}

	openedVault := rand.Intn(100)
	if isHeistOn && openedVault >= 70 {
		fmt.Print("\nGrab and GO!\n")
	} else if isHeistOn {
		isHeistOn = false
		fmt.Println("\nThe vault can't be opened.")
	}

	leftSafely := rand.Intn(5)
	if isHeistOn {
		switch leftSafely {
		case 0:
			isHeistOn = false
			fmt.Println("\nThe jig is up. The heist failed.")
		case 1:
			isHeistOn = false
			fmt.Println("\nWhen did they start raising dogs in vaults??")
		case 2:
			isHeistOn = false
			fmt.Println("\nLooks like this fingerprint scanner won’t accept any fingerprint…")
		case 3:
			isHeistOn = false
			fmt.Println("\nDid I even pack the burlap bags?")
		default:
			fmt.Println("\nStart the getaway car!")
		}
	}

	if isHeistOn {
		amtStolen := 10000 + rand.Intn(1000000)
		fmt.Printf("\nWe got away with $%d!\n", amtStolen)
	}

	fmt.Println("\nisHeistOn is currently:", isHeistOn)
}

