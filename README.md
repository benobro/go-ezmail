# Go-EZmail - a simple Go program to send email (from Gmail or otherwise) 

Intro
-----

Go-EZmail is a lightweight and simple Go program which allows you to send SMTP emails using the "net/smtp" package. It allows you to use HTML syntax to customize the message body, and it also allows you to set a subject line. Go-EZmail is intented for use with Gmail, but you can configure it to send emails from the email service of your choice.

Usage
----

Go-EZmail is simple and mostly useless on its own, but is great for use in bigger projects that require a program to send mail automatically. 

To use Go-EZmail by its self:

1. Install by running `git clone https://github.com/benobro/go-ezmail/` **from your Go workspace** (make sure you already [have Go installed](https://golang.org/doc/install)) 
2. Make the neccessary changes to `go-ezmail.go` (to configure sender and reciever, message/subject, and server/port if needed) 
3. `go run go-ezmail.go` to run - if there's no output, it has been run successfully 

As far as using Go-EZmail with your project, you should be able to put `go-ezmail.go` in the directory with the rest of the Go code for your project and run `go build` to compile everything (obviously making neccessary changes beforehand to implement Go-EZmail)

Notes
----

If you're using Gmail, you need to make sure turn on [less secure apps access](https://support.google.com/accounts/answer/6010255?hl=en)

The default port is 587 (TLS) which works with Gmail - make sure to change it if as needed if you're using a different email service 

Go-EZmail is written by Ben O'Brien [(@benobro)](https://github.com/benobro). 
