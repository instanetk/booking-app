package main

import "fmt"

func main() {
	conferenceName := "Go Conference"
	const conferenceTickets = 50
	var remainingTickets = 50

	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have %v tickets and %v are still available\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

	fmt.Printf("conferenceName is type %T, conferenceTickets is type %T and remainingTickets is type %T.\n", conferenceName, conferenceTickets, remainingTickets)

	var userName string
	var userTickets int
	// ask user for their name

	userName = "David"
	userTickets = 2

	fmt.Printf("User %v booked %v tickets.\n", userName, userTickets)

}
