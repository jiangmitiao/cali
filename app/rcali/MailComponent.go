package rcali

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"net/smtp"
	"strings"
	"time"
)

/**
user := "xxxxx"
	password := ""
	host := "smtp.sina.com:25"
	components.SendToMail(user, password, host, to, subject, body, "plain")
*/

func SendToMail(user, password, host, to, subject, body, mailtype string) error {
	hp := strings.Split(host, ":")
	auth := smtp.PlainAuth("", user, password, hp[0])
	var content_type string
	if mailtype == "html" {
		content_type = "Content-Type: text/" + mailtype + "; charset=UTF-8"
	} else {
		content_type = "Content-Type: text/plain" + "; charset=UTF-8"
	}

	msg := []byte("To: " + to + "\r\nFrom: " + user + "\r\nSubject: " + subject + "\r\n" + content_type + "\r\n\r\n" + body)
	send_to := strings.Split(to, ";")
	//fmt.Println(host)
	//fmt.Println(auth)
	//fmt.Println(user)
	//fmt.Println(send_to)

	err := smtp.SendMail(host, auth, user, send_to, msg)

	return err
}

//SendEmailWithAttachment : send email with attachment
func SendEmailWithAttachment(user, password, hostAndPort, toMailList, subject, body, mailtype string) error {
	host := strings.Split(hostAndPort, ":")
	auth := smtp.PlainAuth("", user, password, host[0])
	buffer := bytes.NewBuffer(nil)

	boudary := "THIS_IS_BOUNDARY_JUST_MAKE_YOURS"
	header := fmt.Sprintf(
		"To:%s\r\n"+
			"From:%s\r\n"+
			"Subject:%s\r\n"+
			"Content-Type:multipart/mixed;Boundary=\"%s\"\r\n"+
			"Mime-Version:1.0\r\n"+
			"Date:%s\r\n"+
			"\r\n\r\n--%s\r\n",
		toMailList, user, subject, boudary, time.Now().String(), boudary)
	buffer.WriteString(header)

	//内容
	var content_type string
	if mailtype == "html" {
		content_type = "Content-Type: text/" + mailtype + "; charset=UTF-8"
	} else {
		content_type = "Content-Type: text/plain" + "; charset=UTF-8"
	}

	msg1 := content_type + "\r\n\r\n" + body + "\r\n"

	buffer.WriteString(msg1)

	msg2 := fmt.Sprintf(
		"\r\n--%s\r\n"+
			"Content-Transfer-Encoding: base64\r\n"+
			"Content-Disposition: attachment;\r\n"+
			"Content-Type:image/jpg;name=\"test.jpg\"\r\n", boudary)
	buffer.WriteString(msg2)
	fmt.Print(msg2)

	attachmentBytes, err := ioutil.ReadFile("E:\\gopath\\src\\github.com\\jiangmitiao\\go-learn\\mail_learn\\test.jpg")
	if err != nil {
		fmt.Println("ReadFile ./test.jpg Error : " + err.Error())
		return err
	}
	b := make([]byte, base64.StdEncoding.EncodedLen(len(attachmentBytes)))
	base64.StdEncoding.Encode(b, attachmentBytes)
	buffer.WriteString("\r\n")
	fmt.Print("\r\n")
	fmt.Print("图片base64编码")
	for i, l := 0, len(b); i < l; i++ {
		buffer.WriteByte(b[i])
		if (i+1)%76 == 0 {
			buffer.WriteString("\r\n")
		}
	}

	buffer.WriteString("\r\n--" + boudary + "--")
	fmt.Print("\r\n--" + boudary + "--")

	sendto := strings.Split(toMailList, ";")
	err = smtp.SendMail(hostAndPort, auth, user, sendto, buffer.Bytes())

	return err
}
