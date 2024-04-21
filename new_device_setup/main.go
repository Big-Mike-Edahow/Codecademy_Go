/* main.go */

package main

import "fmt"

func main() {
  var name string
  fmt.Println("Please enter your name:")
  fmt.Scan(&name)
  fmt.Print("Hello, ", name, "!\n")

  var age int
  fmt.Println("Please enter your age:")
  fmt.Scan(&age)
  fmt.Printf("%v is %d year old.\n", name, age)

  var newMessage = fmt.Sprintf("Hi %v, you are %d years old.", name, age)
  fmt.Println(newMessage)
}

