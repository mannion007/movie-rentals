package main

type NewReleaseMovie struct {
    title string
}

func (newRelease NewReleaseMovie) AmountForRentalLasting(days int) float64 {
    return float64(days * 3)
}

func (newRelease NewReleaseMovie) FrequentRenterPointsEarnedForRentalLasting(days int) int {
    if days > 1 {
        return 2
    }
    return 1
}

func (newRelease NewReleaseMovie) Title() string {
    return newRelease.title
}
