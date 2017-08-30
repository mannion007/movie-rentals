package main

import "fmt"

func main() {
    c := Customer{Name: "James Mannion"}
    c.AddRental(Rental{Movie: ChildrensMovie{"Finding Nemo"}, DaysRented: 10})
    c.AddRental(Rental{Movie: RegularMovie{"Terminator 2"}, DaysRented: 5})
    c.AddRental(Rental{Movie: ChildrensMovie{"Monsters inc"}, DaysRented: 7})
    c.AddRental(Rental{Movie: NewReleaseMovie{"Atomic Blonde"}, DaysRented: 1})
    fmt.Println(c.ProduceStatement(PlainTextStatementProducer{}))
    fmt.Println(c.ProduceStatement(HtmlStatementProducer{}))
}