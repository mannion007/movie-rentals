package main

import "fmt"

type HtmlStatementProducer struct {}

func (htmlStatementProducer HtmlStatementProducer) ProduceStatementFor(customer Customer) string {
    var content = fmt.Sprintf("<h1>Rental Record for %s</h1>", customer.Name)
    content += "<ul>"
    for _, rental := range customer.Rentals {
        content += fmt.Sprintf("<li>%s - %.2f</li>", rental.Movie.Title(), rental.Amount())
    }
    content += "</ul>"
    content += fmt.Sprintf("<p><strong>Amount owed is %.2f\n</strong></p>", customer.AmountOwed())
    content += fmt.Sprintf("<p><strong>You earned %d frequent renter points</strong></p>", customer.FrequentRenterPointsEarned())
    return content
}