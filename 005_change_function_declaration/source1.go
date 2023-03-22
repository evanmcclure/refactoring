package main

import (
	"log"
	"math"
)

// Change Function Declaration lets us understand a function better
// by its name, reduces coupling, and increases encapsulaton. In
// general, there are three cases and the are:
//
// Case 1) Renaming a function using simple mechanics
// Case 2) Renaming a function using migration mechanics
// Case 2) Adding a parameter using migration mechanics
// Case 3) Changing a parameter to one of its properties using migration mechanics

func printCircumference(radius float64) {
	c := circum(radius)
	log.Println("Circumference:", c)
}

func circum(radius float64) float64 {
	return 2 * math.Pi * radius
}
