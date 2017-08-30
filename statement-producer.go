package main

type StatementProducer interface {
    ProduceStatementFor(customer Customer) string
}