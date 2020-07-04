package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/smtp"
)

type smtpServer struct {
	host string
	port string
}

func (s *smtpServer) Address() string {
	return s.host + ":" + s.port
}

func main() {
	url := "https://api.ipify.org?format=text"
	fmt.Println(string("IP adresi alınıyor..."))
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	ip, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println("IP adresiniz: " + string(ip))

	from := "berked20031@gmail.com"
	password := "nebeka565771"
	fmt.Println(string("Posta Gönderiliyor..."))
	to := []string{"berked2003@hotmail.com", "awsd2003@hotmail.com"}
	smtpServer := smtpServer{host: "smtp.gmail.com", port: "587"}
	message := []byte(
		string("Subject: IP Adresiniz Deişti Berke Hazretleri\r\n") +
			"\r\n" +
			"IP Adresiniz: " +
			string(ip) +
			"\r\n")
	auth := smtp.PlainAuth("", from, password, smtpServer.host)
	smtperr := smtp.SendMail(smtpServer.Address(), auth, from, to, message)
	if smtperr != nil {
		fmt.Println(smtperr)
		return
	}
	fmt.Println("Posta Gonderildi")

}
