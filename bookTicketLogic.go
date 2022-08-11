package main

import "fmt"

func main() {
	var conferenceName = "Go conference"
	const conferenceTickets int = 50
	var remainingTickets uint = 50
	var bookings [50]string

	fmt.Println("Welcome to", conferenceName, "booking application")
	fmt.Println("We have total of", conferenceTickets, "tickets and", remainingTickets)
	fmt.Println("get your tickets here to attend")

	var firstName string
	var lastName string
	var email string
	var userTickets uint

	fmt.Println("Enter your first name: ")
	fmt.Scan(&firstName)
	fmt.Println("Enter your first name: ")
	fmt.Scan(&lastName)
	fmt.Println("Enter your email: ")
	fmt.Scan(&email)
	fmt.Println("Enter number of ticket : ")
	fmt.Scan(&userTickets)

	remainingTickets = remainingTickets - userTickets

	bookings[0] = firstName + " " + lastName

	fmt.Printf("The Whole array: %v\n", bookings)
	fmt.Printf("The Whole array: %v\n", bookings[0])

	fmt.Printf("Array type: %T\n", bookings)
	fmt.Printf("Array length: %v\n", len(bookings))

	fmt.Printf("Thank you %v %v for booking %v tickets. You will  receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)

}
