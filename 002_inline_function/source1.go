package main

// Inline Function copies the body of a function into the callers so
// that the function can be removed. It's used to remove needless
// indirection, undo a previous extraction, or to refactor a bunch
// of related functions in a new way. It is the inverse of Extract
// Function.
//
// Functions shouldn't be inlined when there is recursion, multiple
// return points, or when the function is a method being inlined
// into another object and you don't have accessors.
//
// Mechanics:
//
//   1. Don't inline if this function is a polymorphic method.
//   2. Find all of the callers of the function.
//   3. Replace each call with the body of the function.
//   4. Remove the function definition.

type Driver struct {
	NumberOfLateDeliveries int
}

func rating(aDriver Driver) int {
	if moreThanFiveLateDeliveries(aDriver) {
		return 2
	} else {
		return 1
	}
}

func moreThanFiveLateDeliveries(dvr Driver) bool {
	return dvr.NumberOfLateDeliveries > 5
}
