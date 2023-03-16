package main

import (
	"fmt"
	"strings"
)

const conferenceTickets int = 50

var conferenceName string = "Go conference"
var remainingTickets uint = 50
var bookings = []string{}

func main() {

	greetUser()

	for remainingTickets > 0 && len(bookings) <= 50 {

		userName, lastName, email, userTickets := getUserInput()

		validUser, validEmail, validNumber := valideUserInputs(userName, lastName, email, userTickets)

		if validUser && validEmail && validNumber {

			bookingLogic(remainingTickets, userTickets, userName, lastName, email)

			firstNames := getFirstNames()
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

func greetUser() {
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend.")
}

func getFirstNames() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		var names = strings.Fields(booking)
		firstNames = append(firstNames, names[0])
	}
	return firstNames

}
func valideUserInputs(userName string, lastName string, email string, userTickets uint) (bool, bool, bool) {
	validUser := len(userName) >= 2 && len(lastName) >= 2
	validEmail := strings.Contains(email, "@")
	validNumber := userTickets > 0 && userTickets <= remainingTickets
	return validUser, validEmail, validNumber
}
func getUserInput() (string, string, string, uint) {

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

	return userName, lastName, email, userTickets

}

func bookingLogic(remainingTickets uint, userTickets uint, userName string, lastName string, email string) {
	remainingTickets = remainingTickets - userTickets
	bookings = append(bookings, userName+" "+lastName)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive your confirmation number at %v\n", userName, lastName, userTickets, email)
	fmt.Printf("%v tickets reamining for %v\n", remainingTickets, conferenceName)
}
