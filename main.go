package main

import "fmt"

func main() {
	var conferenceName = "Go Conference"
	const conferenceTickets = 50
	var remainingTickets = 50

	fmt.Printf("Welcome to our %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend", conferenceName)

	var userName string
	var userTickets int
	// ask user for their name

	userName = "Tom"
	userTickets = 2

	fmt.Printf("User %v booked %v tickets.\n", userName, userTickets)

}
