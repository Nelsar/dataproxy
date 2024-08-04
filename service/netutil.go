package service

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

func GetDownload(url string) []byte {
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	defer response.Body.Close()

	byteValue, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	return byteValue
}

func GetJsonData[T any](url string) T {

	response, err := http.Get(url)
	if err != nil {
		log.Fatalln(err)

	}

	defer response.Body.Close()

	var out T

	body, err := ioutil.ReadAll(response.Body)

	if err != nil {
		log.Fatalln(err)
	}

	err = json.Unmarshal(body, &out)
	if err != nil {
		log.Fatalln(err)
	}

	return out
}

func GetData(url string) []byte {
	response, err := http.Get(url)
	if err != nil {
		log.Fatalln(err)
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatalln(err)
	}
	return body
}
