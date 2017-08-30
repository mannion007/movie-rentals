package main

type Rental struct {
    Movie      Rentable
    DaysRented int
}

func (rental Rental) Amount() float64 {
    return rental.Movie.AmountForRentalLasting(rental.DaysRented)
}

func (rental Rental) FrequentRenterPointsEarned() int {
    return rental.Movie.FrequentRenterPointsEarnedForRentalLasting(rental.DaysRented)
}