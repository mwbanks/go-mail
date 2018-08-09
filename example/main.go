package main

import (
	"encoding/json"
	gm "github.com/mwbanks/goMail"
	"io/ioutil"
	"net/mail"
	"os"
)

type userAuth struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type config struct {
	Auth    userAuth     `json:"auth"`
	Server  string       `json:"server"`
	To      mail.Address `json:"to"`
	From    mail.Address `json:"from"`
	Subject string       `json:"subject"`
	Body    string       `json:"body"`
}

func main() {

	jsonFile, err := os.Open("config.json")
	if err != nil {
		// fmt.Errorf("%s", err)
		panic(err)
	}
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var c config

	json.Unmarshal(byteValue, &c)

	gm.Init(c.Auth.Username, c.Auth.Password, c.Server)

	gm.SendMail(c.To, c.From, c.Subject, c.Body)
}
