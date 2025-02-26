package utils

import (
	"fmt"
	"os"
	"strconv"

	"gopkg.in/gomail.v2"
)

func SendUserAccessEmail(email string, machineId string, userName string, password string) error {
	smtpHost := os.Getenv("SMTP_HOST")

	if smtpHost == "" {
		return fmt.Errorf("SMTP_HOST env was missing")
	}

	smtpPort := os.Getenv("SMTP_PORT")

	if smtpPort == "" {
		return fmt.Errorf("SMTP_PORT env was missing")
	}

	hostEmail := os.Getenv("HOST_EMAIL")

	if hostEmail == "" {
		return fmt.Errorf("HOST_EMAIL env was missing")
	}

	appPassword := os.Getenv("APP_PASSWORD")

	if appPassword == "" {
		return fmt.Errorf("APP_PASSWORD env was missing")
	}

	to := email

	smtpPortInt, err := strconv.Atoi(smtpPort)

	if err != nil {
		return err
	}

	client := gomail.NewDialer(smtpHost, smtpPortInt, hostEmail, appPassword)

	htmlTemplate := GetUserAccessEmailTemplate(machineId, userName, password)

	message := gomail.NewMessage()

	message.SetHeader("From", hostEmail)
	message.SetHeader("To", to)
	message.SetHeader("Subject", "Welcome to Vithsutra Technologies")
	message.SetBody("text/html", htmlTemplate)

	if err := client.DialAndSend(message); err != nil {
		return err
	}

	return nil
}
