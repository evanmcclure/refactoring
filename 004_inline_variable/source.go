package main

// Inline Variable helps us to remove useless indirection when the
// name of a variable doesn't convey more information than the
// righthand side of the expression does, or if the variable gets
// in the way of another refactoring. It's the inverse of Extract
// Variable.
//
// Mechanics:
//
//   1. Don't inline if there will be side-effects in the right-
//      hand side of the statement.
//   2. Make the variable immutable.
//   3. Find each reference to the variable and replace it with the
//      expression.
//   4. Remove the variable.

type Order struct {
	BasePrice float64
}

func allowFreeShipping(anOrder Order) bool {
	basePrice := anOrder.BasePrice
	return basePrice > 1000
}
