package main

import (
	"encoding/json"
	"io/ioutil"
	"net/mail"
	"os"
)

type config struct {
	Username string       `json:"username"`
	Password string       `json:"password"`
	Server   string       `json:"server"`
	To       mail.Address `json:"to"`
	From     mail.Address `json:"from"`
	Subject  string       `json:"subject"`
	Body     string       `json:"body"`
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

	Init(c.Username, c.Password, c.Server)

	SendMail(c.To, c.From, c.Subject, c.Body)
}
