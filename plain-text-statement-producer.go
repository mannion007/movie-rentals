package main

import "fmt"

type PlainTextStatementProducer struct {}

func (plainTextStatementProducer PlainTextStatementProducer) ProduceStatementFor(customer Customer) string {
    var content = fmt.Sprintf("Rental Record for %s\n", customer.Name)
    for _, rental := range customer.Rentals {
        content += fmt.Sprintf("\t%s\t%.2f\n", rental.Movie.Title(), rental.Amount())
    }
    content += fmt.Sprintf("Amount owed is %.2f\n", customer.AmountOwed())
    content += fmt.Sprintf("You earned %d frequent renter points", customer.FrequentRenterPointsEarned())
    return content
}