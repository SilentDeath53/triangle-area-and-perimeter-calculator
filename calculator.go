package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b, c float64

	fmt.Println("Enter the lengths of the sides of the triangle:")
	fmt.Print("Side a: ")
	fmt.Scanln(&a)
	fmt.Print("Side b: ")
	fmt.Scanln(&b)
	fmt.Print("Side c: ")
	fmt.Scanln(&c)

	perimeter := a + b + c

	semiperimeter := perimeter / 2

	area := math.Sqrt(semiperimeter * (semiperimeter - a) * (semiperimeter - b) * (semiperimeter - c))

	fmt.Printf("Perimeter: %.2f\n", perimeter)
	fmt.Printf("Area: %.2f\n", area)
}
