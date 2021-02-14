package pkg

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/smtp"
	"net/url"
	"os"
	"strings"
)

var (
	// Twilio API configuration
	mainPhone  = os.Getenv("MAIN_PHONE")
	accountSID = os.Getenv("TWILIO_ACCOUNT_SID")
	authToken  = os.Getenv("TWILIO_AUTH_TOKEN")
	baseURL    = "https://api.twilio.com/2010-04-01/Accounts/" + accountSID + "/Messages.json"

	// Google SMTP configuration
	mainEmail    = os.Getenv("MAIN_EMAIL")
	mainPassword = os.Getenv("MAIN_EMAIL_PASSWORD")
	host         = "smtp.gmail.com"
	port         = ":587"
)

const (
	phoneNotification      = "1"
	emailNotification      = "2"
	phoneEmailNotification = "3"
)

var nameToEvent = map[string]func(customer *Customer, msg string) (string, error){
	phoneNotification:      phone,
	emailNotification:      email,
	phoneEmailNotification: phoneEmail,
}

func phone(customer *Customer, msg string) (string, error) {
	// Request body configuration
	reqData := url.Values{}
	reqData.Set("To", customer.phone)
	reqData.Set("From", mainPhone)
	reqData.Set("Body", msg)

	reader := *strings.NewReader(reqData.Encode())
	req, _ := http.NewRequest("POST", baseURL, &reader)
	req.SetBasicAuth(accountSID, authToken)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	httpClient := &http.Client{}
	rsp, err := httpClient.Do(req)
	if err != nil {
		log.Println("Error while sending notification to phone: ", err)
		return "Unable to send notification to phone number " + customer.phone, err
	}

	if rsp.StatusCode < 200 || rsp.StatusCode >= 300 {
		log.Println("Phone method error: status: ", rsp.Status)
		return "Unable to send notification to phone number " + customer.phone, err
	}

	var rspData map[string]interface{}
	decoder := json.NewDecoder(rsp.Body)
	if err := decoder.Decode(&rspData); err != nil {
		fmt.Println(rspData["sid"])
		log.Println("Server response error: ", rspData["sid"])
		return "Unable to send notification to phone number " + customer.phone, err
	}

	return "Successfully sent notification to phone number " + customer.phone, nil
}

func email(customer *Customer, msg string) (string, error) {
	auth := smtp.PlainAuth("", mainEmail, mainPassword, host)

	to := []string{customer.email}
	body := "To: " + customer.email + "\n" +
		"Subject: Alif Shop Purchase Order\n" +
		"\n" +
		msg + "\n"

	if err := smtp.SendMail(host+port, auth, mainEmail, to, []byte(body)); err != nil {
		log.Println("Couldn't send an email: ", err)
		return "Unable to send notification to email number " + customer.email, err
	}

	return "Successfully sent order to an email " + customer.email, nil
}

func phoneEmail(customer *Customer, msg string) (string, error) {
	phoneResult, err := phone(customer, msg)
	if err != nil {
		log.Printf("%v", err)
	}

	emailResult, err := email(customer, msg)
	if err != nil {
		log.Printf("%v", err)
	}

	return fmt.Sprintf("%v\n%v", phoneResult, emailResult), nil
}

func (b *Customer) Notify(msg, notificationType string) (string, error) {
	method := nameToEvent[notificationType]
	return method(b, msg)
}
