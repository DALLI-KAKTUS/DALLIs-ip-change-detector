package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/smtp"
	"time"
)

type smtpServer struct {
	host string
	port string
}

func (s *smtpServer) Address() string {
	return s.host + ":" + s.port
}

func main() {
	var ipOld string
	for {
		url := "https://api.ipify.org?format=text"
		fmt.Println(string("IP adresi alınıyor..."))
		resp, err := http.Get(url)
		for err != nil {
			fmt.Println("İnternete Bağlanılamadı Tekrar Deneniyor...")
			time.Sleep(5 * time.Second)
			resp, err = http.Get(url)

		}
		defer resp.Body.Close()
		Ip, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			panic(err)
		}
		ip := string(Ip)
		fmt.Println("IP adresiniz: " + string(ip))
		fmt.Println("IP Adresinde Değişiklik Bekleniyor")

		if ip != ipOld {
			ipOld = ip

			from := ""     //sender address to here
			password := "" //sender E-mail pass to here
			fmt.Println(string("Posta Gönderiliyor..."))
			to := []string{"berked2003@hotmail.com", "awsd2003@hotmail.com"} //reciver E-mails to here (please delete mine)
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
		time.Sleep(5 * time.Second)
	}
}
