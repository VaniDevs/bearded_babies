package notification

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
	"log"
	"net/http"
	"os"
)

func Send(phone string) {
	message := "Please, schedule your appointment at http://www.babygoround.ca/x32dsf3d"
	SendSMS(phone, message)
}

func SendSMS(phone string, message string) {
	values := map[string]string{
		"api_key":    os.Getenv("NOTIFICATION_API_KEY"),
		"api_secret": os.Getenv("NOTIFICATION_API_SECRET"),
		"to":         phone,
		"from":       os.Getenv("NOTIFICATION_SMS_FROM"),
		"text":       message}
	jsonValue, _ := json.Marshal(values)
	http.Post("https://rest.nexmo.com/sms/json", "application/json", bytes.NewBuffer(jsonValue))
}

func SendEmail(email string, name string, subject string, body string) {
	from := mail.NewEmail("Baby Go Round", "support@babygoround.org")
	to := mail.NewEmail(name, email)
	plainTextContent := "and easy to do anywhere, even with Go"
	htmlContent := body
	message := mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)
	client := sendgrid.NewSendClient(os.Getenv("SANDGRID_API_KEY"))
	response, err := client.Send(message)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(response.StatusCode)
		fmt.Println(response.Body)
		fmt.Println(response.Headers)
	}
}
