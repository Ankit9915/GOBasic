package main

import (
	"fmt"
	"strings"
)

func main() {

	var conferenceName = "Go Conference"
	const conferenceTickets = 50
	var remainingTickets = 50
	var bookings []string
	greetUser(conferenceName, conferenceTickets, remainingTickets)
	for {

		firstName, lastName, email, userTickets := userInput()
		isValidName, isValidEmail, isValidTicketNumber := validateUserInput(firstName, lastName, email, userTickets, remainingTickets)

		if isValidName && isValidEmail && isValidTicketNumber {
			bookingTicket(remainingTickets, userTickets, bookings, firstName, lastName, email, conferenceName)

			name := printFirstNames(bookings)
			fmt.Printf("Users name are: %v\n", name)

			if remainingTickets == 0 {
				//end logic
				fmt.Println("SOLD OUT!")
				break
			}
		} else {
			if !isValidName {
				fmt.Println("first or last name is too short")
			}
			if !isValidEmail {
				fmt.Println("Enter valid email id")
			}
			if !isValidTicketNumber {
				fmt.Println("Enter valid tickets")
			}

		}

	}

}
func greetUser(confName string, conferenceTickets int, remainingTickets int) {
	fmt.Printf("Welcome to %v Booking Application\n", confName)
	fmt.Printf("we have total %v Ticktes,HURRY UP!,and %v are still remaining tickets\n", conferenceTickets, remainingTickets)
	fmt.Println("Get Your Tickets Here!")
}
func printFirstNames(bookings []string) []string {
	firstNames := []string{}
	for _, booking := range bookings {
		var names = strings.Fields(booking)
		firstNames = append(firstNames, names[0])
	}
	return firstNames
}
func validateUserInput(firstName string, lastName string, email string, userTickets int, remainingTickets int) (bool, bool, bool) {

	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets

	return isValidName, isValidEmail, isValidTicketNumber
}
func userInput() (string, string, string, int) {

	var firstName string
	var lastName string
	var email string
	var userTickets int
	//ask user for input
	fmt.Println("Enter your first name: ")
	fmt.Scan(&firstName)
	fmt.Println("Enter your last name: ")
	fmt.Scan(&lastName)
	fmt.Println("Enter your email: ")
	fmt.Scan(&email)
	fmt.Println("Enter Number of Tickets: ")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets

}
func bookingTicket(remainingTickets int, userTickets int, bookings []string, firstName string, lastName string, email string, conferenceName string) {
	remainingTickets = remainingTickets - userTickets
	bookings = append(bookings, firstName+" "+lastName)

	fmt.Printf("Thank You! %v %v  for booking %v Tickets ,You will recieve a confirmation email on %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
}
