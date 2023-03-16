package helper

import "strings"

func ValideUserInputs(userName string, lastName string, email string, userTickets uint, remainingTickets uint) (bool, bool, bool) {
	validUser := len(userName) >= 2 && len(lastName) >= 2
	validEmail := strings.Contains(email, "@")
	validNumber := userTickets > 0 && userTickets <= remainingTickets
	return validUser, validEmail, validNumber
}
