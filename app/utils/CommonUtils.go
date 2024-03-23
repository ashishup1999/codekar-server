package utils

import (
	"fmt"
	"net/smtp"
	"os"
	"strings"
)

func GetProjectPreview(html, css, javascript string) string {
	var previewString = fmt.Sprintf(`
	<!DOCTYPE html>
	<html lang="en">
	<head>
	  <meta charset="UTF-8">
	  <meta name="viewport" content="width=device-width, initial-scale=1.0">
	  <title>Your Title Here</title>
	  <style>
		*{
			overflow: hidden;
		}
		html {
		  width: 100%%;
		  height: 100%%;
		}
		body {
			width: 100%%;
			height: 100%%;
			margin:0;
			padding:0;
			box-sizing: border-box;
			pointer-events: none;
		}
		%s
	  </style>
	</head>
	<body>
	  %s
	  <script>
	  %s
	  </script>
	</body>
	</html>`, css, html, javascript)
	return previewString
}

func SendOtpOnMail(email string, otp string) error {
	//Send otp on email
	sender := os.Getenv("MAIL")
	senderPassword := os.Getenv("MAIL_PASSWORD")
	host := os.Getenv("SMTP_HOST")
	port := os.Getenv("SMTP_PORT")

	subject := "Password Change Verification Code"
	body := "Your verification code for changing OTP is: " + otp

	// Set up authentication information.
	auth := smtp.PlainAuth("", sender, senderPassword, host)

	// Compose email headers
	headers := make(map[string]string)
	headers["From"] = sender
	headers["To"] = email
	headers["Subject"] = subject

	// Compose email body
	var message strings.Builder
	for key, value := range headers {
		message.WriteString(fmt.Sprintf("%s: %s\r\n", key, value))
	}
	message.WriteString("\r\n" + body)

	// Send email
	err := smtp.SendMail(host+":"+port, auth, sender, []string{email}, []byte(message.String()))
	if err != nil {
		return err
	}
	return nil
}
