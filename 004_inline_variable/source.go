package main

// Inline Variable helps us to remove useless indirection when the
// name of a variable doesn't convey more information than the
// righthand side of the expression does, or if the variable gets
// in the way of another refactoring. It's the inverse of Extract
// Variable.

type Order struct {
	BasePrice float64
}

func allowFreeShipping(anOrder Order) bool {
	basePrice := anOrder.BasePrice
	return basePrice > 1000
}
