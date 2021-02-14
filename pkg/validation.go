package pkg

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"regexp"
	"strings"
)

var (
	emailRegex = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	phoneRegex = regexp.MustCompile(`^(?:(?:\(?(?:00|\+)([1-4]\d\d|[1-9]\d?)\)?)?[\-\.\ \\\/]?)?((?:\(?\d{1,}\)?[\-\.\ \\\/]?){0,})(?:[\-\.\ \\\/]?(?:#|ext\.?|extension|x)[\-\.\ \\\/]?(\d+))?$`)
)

// Config gets user input
func Config() (*Customer, string) {
	reader := bufio.NewReader(os.Stdin)

	phone := getInput("Please, enter your phone number: ", reader, IsPhoneValid)
	email := getInput("and your email: ", reader, IsEmailValid)

	customer := &Customer{
		ID:    1,
		phone: phone,
		email: email,
	}

	eventType := getInput(
		"\nWhere do you want to receive your purchase order?\n"+
			"Please, enter\n"+
			"- 1 - for phone via SMS\n"+
			"- 2 - for email\n"+
			"- 3 - for both\n",
		reader,
		isNotificationTypeValid)

	return customer, eventType
}

func IsPhoneValid(p string) bool {
	return phoneRegex.MatchString(p)
}

func IsEmailValid(e string) bool {
	if !emailRegex.MatchString(e) {
		return false
	}
	if mx, err := net.LookupMX(strings.Split(e, "@")[1]); err != nil || len(mx) == 0 {
		return false
	}

	return true
}

func isNotificationTypeValid(text string) bool {
	if _, ok := nameToEvent[text]; ok {
		return true
	}

	return false
}

func getInput(intro string, reader *bufio.Reader, validator func(text string) bool) string {
	for {
		fmt.Print(intro)

		input, err := reader.ReadString('\n')
		if err != nil {
			log.Printf("%v", err)
			fmt.Print("\n\nInvalid input, please, try again.\n")
			continue
		}

		input = strings.TrimSpace(input)
		if len(input) == 0 {
			fmt.Print("Input shouldn't be empty!")
			continue
		}

		if !validator(input) {
			log.Printf("Invalid email or phone, please, try again: %v", input)
			continue
		}

		return input
	}
}
