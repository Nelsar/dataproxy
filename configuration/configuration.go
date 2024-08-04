package configuration

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Configuration struct {
	Shipper     Shipper     `json:"Shipper"`
	Distributor Distributor `json:"Distributor"`
	Database    Database    `json:"Database"`
	Currency    Currency    `json:"Currency"`
}

type Distributor struct {
	Yml string `json:"Yml"`
	Api string `json:"Api"`
}

type Database struct {
	Login    string `json:"Login"`
	Password string `json:"Password"`
	Dsn      string `json:"DSN"`
}

type Shipper struct {
	Url         string `json:"Url"`
	AccessToken string `json:"AccessToken"`
}
type Currency struct {
	Url string `json:"Url"`
}

func Exec() Configuration {
	jsonFile, err := os.Open("api.nmart.me/app.json")

	if err != nil {
		fmt.Println(err)
	}

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)
	config := Configuration{}

	json.Unmarshal([]byte(byteValue), &config)

	return config
}
