package main

type RegularMovie struct {
    title string
}

func (regular RegularMovie) AmountForRentalLasting(days int) float64 {
    var amount float64
    if days > 2 {
        amount += (float64)(days-2) * 1.5
    }
    amount += 2.0
    return amount
}

func (regular RegularMovie) FrequentRenterPointsEarnedForRentalLasting(days int) int {
    return 1
}

func (regular RegularMovie) Title() string {
    return regular.title
}