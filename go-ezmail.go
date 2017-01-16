
package main

import (
    "log"
    "net/smtp"
    "strconv"
)

type EmailUser struct {
    Username, Password, 
    EmailServer string
    Port int
}

type MailContent struct {
	body string
	subject string
}

func main() {
	
	mc := &MailContent{"body goes here", "subject goes here"}

	msg := 
		"MIME-Version: 1.0" + " \r\n" +
		"Content-type: text/html" + "\r\n" +
		"Subject: " + mc.subject + "\r\n\r\n" +
		mc.body + "\r\n"
    eu := &EmailUser{"send@email.com", "PASSWORD", "smtp.gmail.com", 587}
    auth := smtp.PlainAuth("", eu.Username, eu.Password, eu.EmailServer)
    err := smtp.SendMail(eu.EmailServer+":"+strconv.Itoa(eu.Port), auth, eu.Username, []string{"receive@email.com"}, []byte(msg))
    if err != nil {
        log.Fatal(err)
    }
}
