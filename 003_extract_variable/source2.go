package main

import "math"

type Record struct {
	Quantity  float64
	ItemPrice float64
}

type Order struct {
	data Record
}

func createOrder(aRecord Record) Order {
	return Order{
		data: aRecord,
	}
}

func (o Order) quantity() float64 {
	return o.data.Quantity
}

func (o Order) itemPrice() float64 {
	return o.data.ItemPrice
}

func (o Order) price() float64 {
	// price is base price - quantity discount + shipping
	return o.quantity()*o.itemPrice() - math.Max(0, o.quantity()-500)*o.itemPrice()*0.5 + math.Min(o.quantity()*o.itemPrice()*0.1, 100)
}
