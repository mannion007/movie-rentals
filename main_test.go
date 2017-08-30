package main

import "testing"

func TestPlainTextStatementForSingleRegularMovieForMultipleDays(t *testing.T) {
    c := Customer{Name: "Bert Sampson"}
    c.AddRental(Rental{Movie: RegularMovie{"Some Regular Movie"}, DaysRented: 1})
    
    var expected string = "Rental Record for Bert Sampson\n\tSome Regular Movie\t2.00\nAmount owed is 2.00\nYou earned 1 frequent renter points"
    var actual string = c.ProduceStatement(PlainTextStatementProducer{})
    
    if actual != expected {
        t.Fatalf("Expected %s, got %s", expected, actual)
    }
}

func TestPlainTextStatementForMultipleRegularMoviesForOneDay(t *testing.T) {
    c := Customer{Name: "Bert Sampson"}
    c.AddRental(Rental{Movie: RegularMovie{"Some Regular Movie"}, DaysRented: 5})
    
    var expected string = "Rental Record for Bert Sampson\n\tSome Regular Movie\t6.50\nAmount owed is 6.50\nYou earned 1 frequent renter points"
    var actual string = c.ProduceStatement(PlainTextStatementProducer{})
    
    if actual != expected {
        t.Fatalf("Expected %s, got %s", expected, actual)
    }
}

func TestPlainTextStatementForMultipleRegularMoviesForMultipleDays(t *testing.T) {
    c := Customer{Name: "Bert Sampson"}
    c.AddRental(Rental{Movie: RegularMovie{"Some Regular Movie"}, DaysRented: 5})
    c.AddRental(Rental{Movie: RegularMovie{"Another Regular Movie"}, DaysRented: 6})
    
    var expected string = "Rental Record for Bert Sampson\n\tSome Regular Movie\t6.50\n\tAnother Regular Movie\t8.00\nAmount owed is 14.50\nYou earned 2 frequent renter points"
    var actual string = c.ProduceStatement(PlainTextStatementProducer{})
    
    if actual != expected {
        t.Fatalf("Expected %s, got %s", expected, actual)
    }
}

func TestPlainTextStatementForSingleNewReleaseMovieForOneDay(t *testing.T) {
    c := Customer{Name: "Bert Sampson"}
    c.AddRental(Rental{Movie: NewReleaseMovie{"Some New Release Movie"}, DaysRented: 1})
    
    var expected string = "Rental Record for Bert Sampson\n\tSome New Release Movie\t3.00\nAmount owed is 3.00\nYou earned 1 frequent renter points"
    var actual string = c.ProduceStatement(PlainTextStatementProducer{})
    
    if actual != expected {
        t.Fatalf("Expected %s, got %s", expected, actual)
    }
}

func TestPlainTextStatementForSingleNewReleaseMovieForMultipleDays(t *testing.T) {
    c := Customer{Name: "Bert Sampson"}
    c.AddRental(Rental{Movie: NewReleaseMovie{"Some New Release Movie"}, DaysRented: 5})
    
    var expected string = "Rental Record for Bert Sampson\n\tSome New Release Movie\t15.00\nAmount owed is 15.00\nYou earned 2 frequent renter points"
    var actual string = c.ProduceStatement(PlainTextStatementProducer{})
    
    if actual != expected {
        t.Fatalf("Expected %s, got %s", expected, actual)
    }
}

func TestPlainTextStatementForMultipleNewReleaseMovieForOneDay(t *testing.T) {
    c := Customer{Name: "Bert Sampson"}
    c.AddRental(Rental{Movie: NewReleaseMovie{"Some New Release Movie"}, DaysRented: 1})
    c.AddRental(Rental{Movie: NewReleaseMovie{"Another New Release Movie"}, DaysRented: 1})
    
    var expected string = "Rental Record for Bert Sampson\n\tSome New Release Movie\t3.00\n\tAnother New Release Movie\t3.00\nAmount owed is 6.00\nYou earned 2 frequent renter points"
    var actual string = c.ProduceStatement(PlainTextStatementProducer{})
    
    if actual != expected {
        t.Fatalf("Expected %s, got %s", expected, actual)
    }
}

func TestPlainTextStatementForMultipleNewReleaseMovieForMultipleDays(t *testing.T) {
    c := Customer{Name: "Bert Sampson"}
    c.AddRental(Rental{Movie: NewReleaseMovie{"Some New Release Movie"}, DaysRented: 5})
    c.AddRental(Rental{Movie: NewReleaseMovie{"Another New Release Movie"}, DaysRented: 5})
    
    var expected string = "Rental Record for Bert Sampson\n\tSome New Release Movie\t15.00\n\tAnother New Release Movie\t15.00\nAmount owed is 30.00\nYou earned 4 frequent renter points"
    var actual string = c.ProduceStatement(PlainTextStatementProducer{})
    
    if actual != expected {
        t.Fatalf("Expected %s, got %s", expected, actual)
    }
}

func TestPlainTextStatementForSingleChildrenRentalForOneDay(t *testing.T) {
    c := Customer{Name: "Bert Sampson"}
    c.AddRental(Rental{Movie: ChildrensMovie{"Some Children's Movie"}, DaysRented: 1})
    
    var expected string = "Rental Record for Bert Sampson\n\tSome Children's Movie\t1.50\nAmount owed is 1.50\nYou earned 1 frequent renter points"
    var actual string = c.ProduceStatement(PlainTextStatementProducer{})
    
    if actual != expected {
        t.Fatalf("Expected %s, got %s", expected, actual)
    }
}

func testPlainTextStatementForSingleChildrenRentalForMultipleDays(t *testing.T) {
    c := Customer{Name: "Bert Sampson"}
    c.AddRental(Rental{Movie: ChildrensMovie{"Some Children's Movie"}, DaysRented: 5})
    
    var expected string = "Rental Record for Bert Sampson\n\tSome Children's Movie\t4.50\nAmount owed is 4.50\nYou earned 1 frequent renter points"
    var actual string = c.ProduceStatement(PlainTextStatementProducer{})
    
    if actual != expected {
        t.Fatalf("Expected %s, got %s", expected, actual)
    }
}

func TestPlainTextStatementForMultipleChildrenRentalForOneDay(t *testing.T) {
    c := Customer{Name: "Bert Sampson"}
    c.AddRental(Rental{Movie: ChildrensMovie{"Some Children's Movie"}, DaysRented: 1})
    c.AddRental(Rental{Movie: ChildrensMovie{"Another Children's Movie"}, DaysRented: 1})
    
    var expected string = "Rental Record for Bert Sampson\n\tSome Children's Movie\t1.50\n\tAnother Children's Movie\t1.50\nAmount owed is 3.00\nYou earned 2 frequent renter points"
    var actual string = c.ProduceStatement(PlainTextStatementProducer{})
    
    if actual != expected {
        t.Fatalf("Expected %s, got %s", expected, actual)
    }
}

func TestPlainTextStatementForMultipleChildrenRentalForMultipleDays(t *testing.T) {
    c := Customer{Name: "Bert Sampson"}
    c.AddRental(Rental{Movie: ChildrensMovie{"Some Children's Movie"}, DaysRented: 5})
    c.AddRental(Rental{Movie: ChildrensMovie{"Another Children's Movie"}, DaysRented: 5})
    
    var expected string = "Rental Record for Bert Sampson\n\tSome Children's Movie\t4.50\n\tAnother Children's Movie\t4.50\nAmount owed is 9.00\nYou earned 2 frequent renter points"
    var actual string = c.ProduceStatement(PlainTextStatementProducer{})
    
    if actual != expected {
        t.Fatalf("Expected %s, got %s", expected, actual)
    }
}

func TestPlainTextStatementForACombinationOfRentalsForOneDay(t *testing.T) {
    c := Customer{Name: "Bert Sampson"}
    c.AddRental(Rental{Movie: RegularMovie{"Some Regular Movie"}, DaysRented: 1})
    c.AddRental(Rental{Movie: NewReleaseMovie{"Some New Release Movie"}, DaysRented: 1})
    c.AddRental(Rental{Movie: ChildrensMovie{"Some Children's Movie"}, DaysRented: 1})
    
    var expected string = "Rental Record for Bert Sampson\n\tSome Regular Movie\t2.00\n\tSome New Release Movie\t3.00\n\tSome Children's Movie\t1.50\nAmount owed is 6.50\nYou earned 3 frequent renter points"
    var actual string = c.ProduceStatement(PlainTextStatementProducer{})
    
    if actual != expected {
        t.Fatalf("Expected %s, got %s", expected, actual)
    }
}

func TestPlainTextStatementForACombinationOfRentalsForMultipleDays(t *testing.T) {
    c := Customer{Name: "Bert Sampson"}
    c.AddRental(Rental{Movie: RegularMovie{"Some Regular Movie"}, DaysRented: 10})
    c.AddRental(Rental{Movie: NewReleaseMovie{"Some New Release Movie"}, DaysRented: 10})
    c.AddRental(Rental{Movie: ChildrensMovie{"Some Children's Movie"}, DaysRented: 10})
    
    var expected string = "Rental Record for Bert Sampson\n\tSome Regular Movie\t14.00\n\tSome New Release Movie\t30.00\n\tSome Children's Movie\t12.00\nAmount owed is 56.00\nYou earned 4 frequent renter points"
    var actual string = c.ProduceStatement(PlainTextStatementProducer{})
    
    if actual != expected {
        t.Fatalf("Expected %s, got %s", expected, actual)
    }
}