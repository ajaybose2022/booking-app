package main

import (
	"fmt"
)

// Package Level Variables
var conferenceName string = "Go Conference"
var remainingTickets uint8 = 50
var bookings = make([]UserData, 0)

type UserData struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets uint8
}

const totalTickets uint8 = 50

func main() {

	GreetUsers(conferenceName, totalTickets, remainingTickets)

	for {
		// Getting the User Input
		firstName, lastName, email, userTickets := getUserInput()
		// Validating the User Input
		isValidNameLent, isValidEmail, isValidTicketNum := ValidateUserInput(firstName, lastName, email, userTickets, remainingTickets)

		if isValidTicketNum {
			remainingTickets = remainingTickets - userTickets

			var userData = UserData{
				firstName:       firstName,
				lastName:        lastName,
				email:           email,
				numberOfTickets: userTickets,
			}
			bookings = append(bookings, userData)
			fmt.Printf("List of bookings is %v\n", bookings)

			fmt.Printf(" \nThanks %v,%v for booking %v tickets. You will receive confirmation email at: %v \n\n", lastName, firstName, userTickets, email)
			fmt.Printf("We have a total of %v total tickets and %v are still remaining \n", totalTickets, remainingTickets)

			firstNames := getFirstNames()
			fmt.Printf("The first Names of our bookings : %v \n\n", firstNames)

			if remainingTickets == 0 {
				fmt.Println("Our Conference is booked out. Come back next year.")
				break
			}
		} else {
			if isValidNameLent {
				fmt.Printf("First Name or Last Name is too short")
			}
			if isValidEmail {
				fmt.Println("Email doesn't contain a '@' sign ")
			}
			if isValidTicketNum {
				fmt.Println("Number of tickets is invalid")
			}
			continue
		}
	}
}

func getFirstNames() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		firstNames = append(firstNames, booking.firstName)
	}
	return firstNames
}

func getUserInput() (string, string, string, uint8) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint8

	// ask the user for their first name
	fmt.Printf("Enter your first name: ")
	fmt.Scan(&firstName)

	// ask the user for their last name
	fmt.Printf("Enter your last name: ")
	fmt.Scan(&lastName)

	// ask the user for their email
	fmt.Printf("Enter your email: ")
	fmt.Scan(&email)

	// ask the user for the number of tickets needed
	fmt.Printf("Enter Number of tickets you needed : ")
	fmt.Scan(&userTickets)

	fmt.Printf("\n")
	return firstName, lastName, email, userTickets
}
