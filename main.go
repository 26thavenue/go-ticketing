package main

import "fmt"

func main() {
	event := "Developer's Club"
	const totalTickets = 50
	ticketsRemaining := 50
	fmt.Printf("Welcome to %v,the total number of tickets are %v and there are only %v left , Grab your copy now !!!", event, totalTickets, ticketsRemaining)
}