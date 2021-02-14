package main

import (
	"fmt"
	"log"

	driver "github.com/MrHakimov/buy-event/pkg"
)

func main() {
	customer, notificationType := driver.Config()
	msg := driver.ShowPurchase(*driver.SamplePurchase)
	fmt.Println(msg)

	result, err := customer.Notify(msg, notificationType)
	if err != nil {
		log.Println(err)
	}

	log.Println(result)
}
