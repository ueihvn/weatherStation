package main

import (
	"fmt"
	"log"
	"net/smtp"
	"reflect"
	"strings"
)

type Sender struct {
	UserName   string
	Password   string
	SMTPServer string
	Port       string
}

var condition = make(map[string]float32, 2)

const (
	MAX_HEAT     = 36.0
	MAX_HUMIDITY = 90.0
)

func (sender Sender) SendMail(to []string, subject, bodyMessage string) {
	auth := smtp.PlainAuth("", sender.UserName, sender.Password, sender.SMTPServer)

	msg := "From: " + sender.UserName + "\r\n" +
		"To: " + strings.Join(to, ",") + "\r\n" +
		"Subject: " + subject + "\r\n" + "\r\n" + bodyMessage + "\r\n"

	err := smtp.SendMail(sender.SMTPServer+":"+sender.Port,
		auth, sender.UserName, to, []byte(msg))
	if err != nil {
		log.Fatal(err)
	}
}

func checkConditionalForSendMail(tsLastEntry TSLastEntry, condition map[string]float32) (msg string) {
	s := reflect.ValueOf(&tsLastEntry)

	for key, val := range condition {
		valueOfField := reflect.Indirect(s).FieldByName(key).Float()

		if float32(valueOfField) >= val {
			strVal := fmt.Sprintf("%.2f", val)
			strValueOfField := fmt.Sprintf("%f", valueOfField)
			msg += strings.ToUpper(key) + " reached at: " + strValueOfField + " .LIMIT : " + strVal + "\n"
		}
	}
	return
}

func sendMailProcess(ts TSLastEntry) {

	condition["Tempature"] = MAX_HEAT
	condition["Humidity"] = MAX_HUMIDITY

	msg := checkConditionalForSendMail(ts, condition)
	fmt.Println(msg)
	if msg != "" {
		sender := Sender{
			UserName:   getEnvByKey("GMAIL_USERNAME"),
			Password:   getEnvByKey("GMAIL_PASSWORD"),
			SMTPServer: "smtp.gmail.com",
			Port:       "587",
		}
		to := []string{"18520052@gm.uit.edu.vn"}
		sender.SendMail(to, "Alert from IOT Server", msg)
	}
}
