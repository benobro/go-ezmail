
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


const emailTemplate = `From someone`
func main() {
	msg := "MIME-Version: 1.0" + " \r\n" +
		"Content-type: text/html" + "\r\n" +
		"Subject: subject goes here" + "\r\n\r\n" +
		"message goes here" + "\r\n"
    eu := &EmailUser{"benobrien705@gmail.com", "nqjslvzgamidktur", "smtp.gmail.com", 587}
    auth := smtp.PlainAuth("", eu.Username, eu.Password, eu.EmailServer)
    err := smtp.SendMail(eu.EmailServer+":"+strconv.Itoa(eu.Port), auth, eu.Username, []string{"ben.obrien000@gmail.com"}, []byte(msg))
    if err != nil {
        log.Fatal(err)
    }
}
