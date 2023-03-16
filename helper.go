package main

import "strings"

func valideUserInputs(userName string, lastName string, email string, userTickets uint) (bool, bool, bool) {
	validUser := len(userName) >= 2 && len(lastName) >= 2
	validEmail := strings.Contains(email, "@")
	validNumber := userTickets > 0 && userTickets <= remainingTickets
	return validUser, validEmail, validNumber
}
