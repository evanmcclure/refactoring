package main

import "math"

// Extract Variable makes complex code more readable and gives us
// variables that can be read in a debugger by being set to a copy
// of the expression being broken apart. It's the inverse of
// Inline Variable.

type Order struct {
	Quantity  float64
	ItemPrice float64
}

func price(order Order) float64 {
	// price is base price - quantity discount + shipping
	return order.Quantity*order.ItemPrice - math.Max(0, order.Quantity-500)*order.ItemPrice*0.5 + math.Min(order.Quantity*order.ItemPrice*0.1, 100)
}
