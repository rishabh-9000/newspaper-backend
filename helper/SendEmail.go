package helper

import (
	"log"
	"net/smtp"
)

// SendEmail : Sends Email
// body : Body of Email
// to : Receiver's Email
// subject : Subject of Email
func SendEmail(to string, subject string, body string) (string, error) {
	from := "my.newspaper.app@gmail.com"
	pass := "Shinratensei@99"

	msg := "From: " + from + "\n" +
		"To: " + to + "\n" +
		"Subject: " + subject + "\n\n" +
		body

	e := smtp.SendMail("smtp.gmail.com:587",
		smtp.PlainAuth("", from, pass, "smtp.gmail.com"),
		from, []string{to}, []byte(msg))

	if e != nil {
		log.Printf("smtp error: %s", e)
		return "", e
	}

	result := "Successfully Sent Email to : " + to

	return result, nil
}
