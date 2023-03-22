package main

type Customer struct {
}

type Book struct {
	reservations []Customer
}

func (b *Book) AddReservation(customer Customer) {
	b.reservations = append(b.reservations, customer)
}

func ReserveBook(book Book, customer Customer) {
	book.AddReservation(customer)
}
