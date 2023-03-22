package main

import (
	"log"
	"math"
)

// Change Function Declaration lets us understand a function better
// by its name, reduces coupling, and increases encapsulaton. In
// general, there are three cases and the are:
//
//   Case A) Renaming a function using simple mechanics
//   Case B) Renaming a function using migration mechanics
//   Case C) Adding a parameter using migration mechanics
//   Case D) Changing a parameter to one of its properties using migration mechanics
//
// Simple Mechanics:
//
//   1. Verify parameters that will be removed aren't referenced in
//      the body of the function.
//   2. Copy the original method and give the copy a unique name.
//   3. Change the declaration of the new method.
//   4. Find all references to the old method and change them to call
//      the new method.
//   5. Remove the old method.
//
// Migration Mechanics:
//
//   1. Refactor the body of the function to make things easy.
//   2. Extract the body of the function into a new function.
//   3. Use simple mechanics to add parameters to the new function
//      if needed.
//   4. Inline the old function.
//   5. Rename the new function if a temporary name was used.

func printCircumference(radius float64) {
	c := circum(radius)
	log.Println("Circumference:", c)
}

func circum(radius float64) float64 {
	return 2 * math.Pi * radius
}
