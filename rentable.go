package main

type Rentable interface {
     Title() string
     AmountForRentalLasting(days int) float64
     FrequentRenterPointsEarnedForRentalLasting(days int) int
}