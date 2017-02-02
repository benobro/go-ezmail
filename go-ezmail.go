
package ezmail

import (
    "log"
    "net/smtp"
    "strconv"
)

func SendMail(username string, password string, server string, port int, toemail string, subject string, message string) {

	msg := 
		"MIME-Version: 1.0" + " \r\n" +
		"Content-type: text/html" + "\r\n" +
		"Reply-To: + replyto + "\r\n" +
		"Subject: " + subject + "\r\n\r\n" + message + "\r\n"
    auth := smtp.PlainAuth("", username, password, server)
    err := smtp.SendMail(server+":"+strconv.Itoa(port), auth, username, []string{toemail}, []byte(msg))
    if err != nil {
        log.Fatal(err)
    }
}
