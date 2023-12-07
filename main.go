package main

import (
	"fmt" 
	"strings"
)

func main() {
	event := "Developer's Club"
	const totalTickets = 50
	var ticketsRemaining uint = 50
	bookings := []string{}
	fmt.Printf("Welcome to %v,the total number of tickets are %v and there are only %v left , Grab your copy now !!!", event, totalTickets, ticketsRemaining)

	for{
		var fullName string
		var email string
		var userTickets uint


		fmt.Println("Enter your name")
		fmt.Scanln(&fullName)

		fmt.Println("Enter your email")
		fmt.Scanln(&email)

		fmt.Println("Enter the number of tickets you want to book")
		fmt.Scanln(&userTickets)

		isValidName := len(fullName) >= 2 
		isValidEmail := len(email) >= 5 && strings.Contains(email, "@")
		isValidTicket := userTickets > 0 && userTickets <= ticketsRemaining 

		var isValidParams bool = isValidName && isValidEmail  && isValidTicket

		if isValidParams{
			ticketsRemaining = ticketsRemaining- userTickets
			bookings = append(bookings, userTickets)

			fmt.Printf("Thank you %v for booking %v tickets for %v', name, userTickets, event")
			fmt.Printf("There are %v tickets remaining", ticketsRemaining)

			fullNames := []string{}

			for _,booking := range bookings{
				var names = strings.Fields(booking)
				fullNames = append(fullNames,names[0])
			}

			fmt.Printf("The people who have booked the tickets are %v", fullNames)

			if ticketsRemaining == 0{
				fmt.Println("Sorry, the tickets are sold out")
				break
			}
		}else{
			if !isValidName{
				fmt.Println("Sorry, the name is invalid")
			}
			if !isValidEmail{
				fmt.Println("Sorry, the name is invalid")
			}
			if !isValidTicket{
				fmt.Println("Sorry, the name is invalid")
			}
		}


	}
}