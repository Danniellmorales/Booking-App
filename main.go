package main

import (
	"fmt"
	"strings"
)

func main() {
	var conferenceName string = "Go conference"
	const conferenceTickets int = 50
	var remainingTickets uint = 50
	bookings := []string{}

	greetUser(conferenceName, conferenceTickets, remainingTickets)

	for remainingTickets > 0 && len(bookings) <= 50 {

		var userName string
		var lastName string
		var userTickets uint
		var email string

		//Ask their user their personal information
		fmt.Println("Enter your first name:")
		fmt.Scan(&userName)

		fmt.Println("Enter your last name:")
		fmt.Scan(&lastName)

		fmt.Println("Can you please provide your email addres?")
		fmt.Scan(&email)

		fmt.Println("How many tickets do you want:")
		fmt.Scan(&userTickets)

		validUser, validEmail, validNumber := valideUserInputs(userName, lastName, email, userTickets, remainingTickets)

		if validUser && validEmail && validNumber {

			remainingTickets = remainingTickets - userTickets
			bookings = append(bookings, userName+" "+lastName)

			fmt.Printf("Thank you %v %v for booking %v tickets. You will receive your confirmation number at %v\n", userName, lastName, userTickets, email)
			fmt.Printf("%v tickets reamining for %v\n", remainingTickets, conferenceName)

			firstNames := getFirstNames(bookings)
			fmt.Printf("The first names of booking are %v\n", firstNames)

			if remainingTickets == 0 {
				fmt.Println("Our conference is booked out. Come back next year.")
				break
			}
		} else {
			if !validUser {
				println("Your name or last name is too short")
			}
			if !validEmail {
				println("Your email addres doesn't contain @ sing")
			}
			if !validNumber {
				println("The number of yout ticket is invalide")
			}

		}

	}

}

func greetUser(conferenceName string, conferenceTickets int, remainingTickets uint) {
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend.")
}

func getFirstNames(bookings []string) []string {
	firstNames := []string{}
	for _, booking := range bookings {
		var names = strings.Fields(booking)
		firstNames = append(firstNames, names[0])
	}
	return firstNames

}
func valideUserInputs(userName string, lastName string, email string, userTickets uint, remainingTickets uint) (bool, bool, bool) {
	validUser := len(userName) >= 2 && len(lastName) >= 2
	validEmail := strings.Contains(email, "@")
	validNumber := userTickets > 0 && userTickets <= remainingTickets
	return validUser, validEmail, validNumber
}
