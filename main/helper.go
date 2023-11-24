package main

import (
	"fmt"
	"strings"
)

func ValidateUserInput(firstName string, lastName string, email string, userTickets uint8, remainingTickets uint8) (bool, bool, bool) {
	var isValidNameLent bool = len(firstName) >= 2 && len(lastName) >= 2
	var isValidEmail bool = strings.Contains(email, "@")
	var isValidTicketNum bool = userTickets <= remainingTickets && userTickets > 0
	return isValidNameLent, isValidEmail, isValidTicketNum
}

func GreetUsers(conferenceName string, totalTickets uint8, remainingTickets uint8) {
	fmt.Printf("Welcome to our %v\n", conferenceName)
	fmt.Printf("We have a total of %v total tickets and %v are still remaining \n", totalTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")
}
