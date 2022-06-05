package gmail

import (
	"bytes"
	"fmt"
	"html/template"

	"net/smtp"
)

var auth smtp.Auth

type Ax interface {
	SEndlogin()
}
type GAmll struct {
	Email    string
	Password string
}

func (g *GAmll) Login(Email, Password string) {

	auth = smtp.PlainAuth("", "ax47chaos@gmail.com", "mki8mki8", "smtp.gmail.com")
	//"ax47chaos@gmail.com", "mki8mki8"
}
func (g GAmll) SEndlogin(Username, otp string) {
	from := "axc47chaos@gmail.com"
	password := "mki8mki8"
	toList := []string{"axc47y@gmail.com"}
	host := "smtp.gmail.com"
	port := "587"
	msg := "Hello geeks!!!"
	body := []byte(msg)
	auth := smtp.PlainAuth("", from, password, host)
	err := smtp.SendMail(host+":"+port, auth, from, toList, body)
	if err != nil {
		fmt.Println(err)

	}

}

type Request struct {
	to      []string
	subject string
	body    string
}

func NewRequest(to []string, subject, body string) *Request {
	return &Request{
		to:      to,
		subject: subject,
		body:    body,
	}
}
func (r *Request) SendEmail() (bool, error) {
	mime := "MIME-version: 1.0;\nContent-Type: text/plain; charset=\"UTF-8\";\n\n"
	subject := "Subject: " + r.subject + "!\n"
	msg := []byte(subject + mime + "\n" + r.body)
	addr := "smtp.gmail.com:587"

	if err := smtp.SendMail(addr, auth, "dhanush@geektrust.in", r.to, msg); err != nil {
		return false, err
	}
	return true, nil
}

func (r *Request) ParseTemplate(templateFileName string, data interface{}) error {
	t, err := template.ParseFiles(templateFileName)
	if err != nil {
		return err
	}
	buf := new(bytes.Buffer)
	if err = t.Execute(buf, data); err != nil {
		return err
	}
	r.body = buf.String()
	return nil
}
