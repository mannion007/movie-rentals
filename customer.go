package main

type Customer struct {
    Name    string
    Rentals []Rental
}

func (customer *Customer) AddRental(rental Rental) {
    customer.Rentals = append(customer.Rentals, rental)
}

func (customer Customer) AmountOwed() float64 {
    var amount float64
    for _, rental := range customer.Rentals {
        amount += rental.Amount()
    }
    return amount
}

func (customer Customer) FrequentRenterPointsEarned() int {
    var points int
    for _, rental := range customer.Rentals {
        points += rental.FrequentRenterPointsEarned()
    }
    return points
}

func (customer Customer) ProduceStatement(producer StatementProducer) string {    
    return producer.ProduceStatementFor(customer)
}