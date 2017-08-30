package main

import "fmt"

type Customer struct {
    Name    string
    Rentals []Rental
}

func (customer *Customer) AddRental(rental Rental) {
    customer.Rentals = append(customer.Rentals, rental)
}

func (customer *Customer) AmountOwed() float64 {
    var amount float64
    for _, rental := range customer.Rentals {
        amount += rental.Amount()
    }
    return amount
}

func (customer *Customer) FrequentRenterPointsEarned() int {
    var points int
    for _, rental := range customer.Rentals {
        points += rental.FrequentRenterPointsEarned()
    }
    return points
}

func (customer *Customer) ProduceStatement() Statement {
    var content = fmt.Sprintf("Rental Record for %s\n", customer.Name)
    for _, rental := range customer.Rentals {
        content += fmt.Sprintf("\t%s\t%.2f\n", rental.Movie.Title(), rental.Amount())
    }
    content += fmt.Sprintf("Amount owed is %.2f\n", customer.AmountOwed())
    content += fmt.Sprintf("You earned %d frequent renter points", customer.FrequentRenterPointsEarned())
    return Statement{Output: content}
}