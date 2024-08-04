package model

type Addresses struct {
	XMLName struct{} `xml:"addresses"`
	Address Address  `xml:"address"`
}

type Address struct {
	XMLName struct{} `xml:"address"`
	Id      int      `xml:"id"`
}
