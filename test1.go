
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


const emailTemplate = `From curry`
func main() {
    eu := &EmailUser{"my@email.com", "PASSWORDr", "smtp.gmail.com", 587}
    auth := smtp.PlainAuth("", eu.Username, eu.Password, eu.EmailServer)
    err := smtp.SendMail(eu.EmailServer+":"+strconv.Itoa(eu.Port), auth, eu.Username, []string{"email@email.com"}, []byte("hello"))
    if err != nil {
        log.Fatal(err)
    }
}
