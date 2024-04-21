/* main.go */

package main

import "fmt"

func main() {
  fmt.Print("Practice using the fmt package!\n\n")

  fmt.Println("What would you like for lunch?")
  var food string
  fmt.Scan(&food)
  fmt.Printf("Sure, we can have %v for lunch.\n", food)

  template := "I wish I had a %v."
  pet := "puppy"
  var wish = fmt.Sprintf(template, pet)
  fmt.Println(wish)

  step1 := "Breathe in..."
  step2 := "Breathe out..."
  
  meditation := fmt.Sprintln(step1, step2)
  fmt.Print(meditation)

  floatExample := 1.75
  fmt.Printf("Working with a %T\n", floatExample) 
  
  yearsOfExp := 3
  reqYearsExp := 15
  fmt.Printf("I have %d years of Go experience,\nand this job is asking for %d years.\n", yearsOfExp, reqYearsExp) 
  
  stockPrice := 3.50
  fmt.Printf("Each share of Gopher feed is $%.2f!\n", stockPrice) 

  animal1 := "cat"
  animal2 := "dog"
  fmt.Printf("Are you a %v or a %v person?\n", animal1, animal2)
}

