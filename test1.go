
package main

import (
    "log"
    "net/smtp"
    "strconv"
)

type EmailUser struct {
    Username, Password, EmailServer string
    Port                            int
}

const emailTemplate = `From foo.bar`

func main() {
    eu := &EmailUser{"benobrien705@gmail.com", "nqjslvzgamidktur", "smtp.gmail.com", 587}
    auth := smtp.PlainAuth("", eu.Username, eu.Password, eu.EmailServer)
    err := smtp.SendMail(eu.EmailServer+":"+strconv.Itoa(eu.Port), auth, eu.Username, []string{"lrivera@whitbyschool.org"}, []byte("hello"))
    if err != nil {
        log.Fatal(err)
    }
}
