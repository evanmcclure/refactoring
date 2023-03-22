package main

// Extract Function puts a fragment of code into its own function,
// gives the new function an appropriate name, and changes the
// calling code to call the new function instead, done in three
// different ways:
//
// Case 1) No variables out of scope
// Case 2) Using local variables
// Case 3) Reassigning a local variable

import (
	"log"
	"time"
)

type Invoice struct {
	Customer string
	DueDate  time.Time
	Orders   []Order
}

type Order struct {
	Amount float64
}

func printOwing(invoice Invoice) {
	outstanding := 0.0

	log.Println("***********************")
	log.Println("**** Customer Owes ****")
	log.Println("***********************")

	// calculate outstanding
	for _, o := range invoice.Orders {
		outstanding += o.Amount
	}

	// record due date
	today := time.Now()
	invoice.DueDate = today.Add(30 * 24 * time.Hour)

	// print details
	log.Printf("name: %s\n", invoice.Customer)
	log.Printf("amount: %f\n", outstanding)
	log.Printf("due: %v\n", invoice.DueDate.Local())
}
