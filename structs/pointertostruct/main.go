// main.go

package main

import "fmt"

type Triangle struct {
	height float32
	base   float32
}

func (tri *Triangle) updateBase(newBase float32) {
	tri.base = newBase
}

func main() {

	triangle := Triangle{10, 4}

	// Call the function here
	triangle.updateBase(13)

	fmt.Println(triangle)

}
