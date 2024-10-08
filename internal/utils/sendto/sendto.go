package sendto

import (
	"bytes"
	"fmt"
	"html/template"
	"net/smtp"

	"github.com/longln/go-ecommerce-backend-api/global"
	"go.uber.org/zap"
)


type EmailAddress struct {
	Addres string `json:"address"`
	Name string `json:"name"`
}

type Mail struct {
	From EmailAddress
	To []string
	Subject string
	Body string
}

func BuildMessage(mail Mail) string {
	msg := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	msg += fmt.Sprintf("From: %s\n", mail.From.Addres)
	msg += fmt.Sprintf("To: %s\n", mail.To)
	msg += fmt.Sprintf("Subject: %s\n", mail.Subject)
	msg += fmt.Sprintf("\n%s\n", mail.Body)
	return msg
}


func SendTextEmailOTP(to []string, from string, otp string) error {
	contentEmail := Mail{
		From: EmailAddress{Addres: from, Name: "test"},
		To: to,
		Subject: "OTP Verification",
		Body: fmt.Sprintf("Your OTP is %s, please enter it in the app", otp),
	}

	messageEmail := BuildMessage(contentEmail)

	// send smtp
	auth := smtp.PlainAuth("", "SMTP_USER", "SMTP_PASSWORD", "SMTP_HOST")
	err := smtp.SendMail("SMTP_HOST:SMTP_PORT", auth, "SMT_USER", to, []byte(messageEmail))
	if err != nil {
		global.Logger.Error("Failed to send email", zap.Error(err))
		return err
	}

	return nil
}


func SendTemplateEmailOTP(
	to []string, from string, nameTemplate string, dataTemplate map[string]interface{}) error {

	htmlBody, err := getMailTemplate(nameTemplate, dataTemplate)
	if err != nil {
		return err
	}

	return send(to, from, htmlBody)
}



func getMailTemplate(nameTempalte string, dataTemplate map[string]interface{}) (string, error) {
	htmlTemplate := new(bytes.Buffer)
	t := template.Must(template.New(nameTempalte).ParseFiles("/templates-email/" + nameTempalte))
	err := t.Execute(htmlTemplate, dataTemplate)
	if err != nil {
		return "", err
	}

	return htmlTemplate.String(), nil
}


func send(to []string, from string, htmlTemplate string) error {
	contentEmail := Mail{
		From: EmailAddress{Addres: from, Name: "test"},
		To: to,
		Subject: "OTP Verification",
		Body: htmlTemplate,
	}

	messageEmail := BuildMessage(contentEmail)

	// send smtp
	auth := smtp.PlainAuth("", "SMTP_USER", "SMTP_PASSWORD", "SMTP_HOST")
	err := smtp.SendMail("SMTP_HOST:SMTP_PORT", auth, "SMT_USER", to, []byte(messageEmail))
	if err != nil {
		global.Logger.Error("Failed to send email", zap.Error(err))
		return err
	}

	return nil
}  