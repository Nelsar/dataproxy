package model

import "encoding/xml"

type Offers struct {
	Offer []Offer `xml:"offer"`
}

type Offer struct {
	XMLName         struct{}  `xml:"offer"`
	Id              int       `xml:"id,attr"`
	Price           float64   `xml:"price"`
	PurchasePrice   float64   `xml:"purchase_price"`
	CurrencyId      string    `xml:"currencyId"`
	CategoryId      int       `xml:"categoryId"`
	Url             string    `xml:"url"`
	VendorCode      int       `xml:"vendorCode"`
	Picture         []Picture `xml:"picture"`
	Available       string    `xml:"available"`
	Vendor          string    `xml:"vendor"`
	Name            string    `xml:"name"`
	Description     string    `xml:"description"`
	Param           []Param   `xml:"param"`
	Quantity        string    `xml:"quantity"`
	QuantityInStock string    `xml:"quantity_in_stock"`
}

type Picture struct {
	XMLName xml.Name `xml:"picture"`
	Value   string   `xml:",chardata"`
}

type Param struct {
	XMLName xml.Name `xml:"param"`
	Name    string   `xml:"name,attr"`
	Value   string   `xml:",chardata"`
}
