// main.go

package main

import "fmt"

type Student struct {
	firstName string
	lastName  string
	age       int
	grade     int
}

func main() {
	var peter = Student{firstName: "Peter", lastName: "Bookman", age: 16, grade: 11}
	fmt.Println(peter)
	scott := Student{"Scott", "Peterson", 17, 12}
	fmt.Println(scott)
}
