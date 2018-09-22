package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
	"log"
	"net/http"
)

func main() {
	fmt.Println("test")
}

func SendSms(text string, phoneNumber string) {
	values := map[string]string{
		"api_key":    "7fa887f7",
		"api_secret": "a0IKffluJqZBOdVP",
		"to":         phoneNumber,
		"from":       "12892324939",
		"text":       text}
	jsonValue, _ := json.Marshal(values)
	http.Post("https://rest.nexmo.com/sms/json", "application/json", bytes.NewBuffer(jsonValue))
}

func SendEmail(subject string, toName string, toEmail string, content string) {
	from := mail.NewEmail("Baby Go Round", "support@babygoround.org")
	to := mail.NewEmail(toName, toEmail)
	plainTextContent := "and easy to do anywhere, even with Go"
	htmlContent := content
	message := mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)
	client := sendgrid.NewSendClient("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")
	response, err := client.Send(message)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(response.StatusCode)
		fmt.Println(response.Body)
		fmt.Println(response.Headers)
	}
}
