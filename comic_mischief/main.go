/* main.go */

package main

import "fmt"

func main() {
  var publisher string
  var writer string
  var artist string
  var title string
  var year int
  var pageNumber int
  var grade float32

  title = "Mr. GoToSleep"
  writer = "Tracey Hatchet"
  artist = "Jewel Tampson"
  publisher = "DizzyBooks Publishing Inc"
  year = 1997
  pageNumber = 14
  grade = 6.5

  fmt.Print("*** Comic Mischief ***\n\n")

  fmt.Println("Title:\t\t", title)
  fmt.Println("Publisher:\t", publisher)
  fmt.Println("Written by:\t", writer)
  fmt.Println("Drawn by:\t", artist)
  fmt.Println("Pages:\t\t", pageNumber)
  fmt.Println("Year:\t\t", year)
  fmt.Println("Grade:\t\t", grade)

  title = "Epic Vol. 1"
  writer = "Ryan N. Shawn"
  artist = "Phoebe Paperclips"
  publisher = "DizzyBooks Publishing Inc"
  year = 2013
  pageNumber = 160
  grade = 9

  fmt.Println("\nTitle:\t\t", title)
  fmt.Println("Publisher:\t", publisher)
  fmt.Println("Written by:\t", writer)
  fmt.Println("Drawn by:\t", artist)
  fmt.Println("Pages:\t\t", pageNumber)
  fmt.Println("Year:\t\t", year)
  fmt.Println("Grade:\t\t", grade)
}

