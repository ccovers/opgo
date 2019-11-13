package main

import (
	"crypto/tls"
	"fmt"
	"net/smtp"
	"strings"

	"gopkg.in/gomail.v2"
)

func main() {
	//GoEmail()
	//return

	username := "alarm@hztl3.com"
	password := "Hztl1311"
	host := "smtp.exmail.qq.com:587"
	to := "jiarong.tian@hztl3.com"
	subject := "golang0"
	body := `Test send to email`
	fmt.Println("send email")

	err := sendEmail(username, password, host, to, subject, body, "html")
	if err != nil {
		fmt.Printf("send email error: %s\n", err.Error())
	} else {
		fmt.Println("send email success")
	}

}

func sendEmail(username, password, host, to, subject, body, mailtype string) error {
	hp := strings.Split(host, ":")
	auth := smtp.PlainAuth("", username, password, hp[0])

	var contentType string
	if mailtype == "html" {
		contentType = fmt.Sprintf("Content-Type: test/%s; charset=UTF-8", mailtype)
	} else {
		contentType = fmt.Sprintf("Content-Type: test/plain; charset=UTF-8")
	}

	msg := fmt.Sprintf("To: chicheng<%s>\r\nFrom: jiarong<%s>\r\nSubject: %s\r\n%s\r\n\r\n%s\r\n",
		to, username, subject, contentType, body)

	sendTo := strings.Split(to, ";")
	msg = "Mime-Version: 1.0\r\nDate: Mon, 11 Nov 2019 17:04:55 +0800\r\nFrom: alarm@hztl3.com\r\nTo: jiarong.tian@hztl3.com\r\nSubject: golang\r\nContent-Type: text/html; charset=UTF-8\r\nContent-Transfer-Encoding: quoted-printable\r\n\r\nthis is test213"
	fmt.Println(msg)
	err := smtp.SendMail(host, auth, username, sendTo, []byte(msg))
	return err
}

func GoEmail() {
	msg := gomail.NewMessage()
	msg.SetHeader("From", "alarm@hztl3.com")
	msg.SetHeader("To", "jiarong.tian@hztl3.com")
	msg.SetHeader("Subject", "golang")
	msg.SetBody("text/html", "this is test")

	dialer := gomail.NewDialer("smtp.exmail.qq.com", 587, "alarm@hztl3.com", "Hztl1311")
	dialer.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	err := dialer.DialAndSend(msg)
	if err != nil {
		fmt.Println(err)
		return
	}
}
