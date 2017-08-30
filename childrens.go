package main

type ChildrensMovie struct {
    title string
}

func (childrens ChildrensMovie) AmountForRentalLasting(days int) float64 {
    amount := 1.5
    if days > 3 {
        amount += (float64)(days-3) * 1.5
    }
    return amount
}

func (childrens ChildrensMovie) FrequentRenterPointsEarnedForRentalLasting(days int) int {
    return 1
}

func (childrens ChildrensMovie) Title() string {
    return childrens.title
}