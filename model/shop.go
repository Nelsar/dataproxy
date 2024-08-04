package model

import "go.starlark.net/lib/time"

type Shop struct {
	XMLName    struct{}   `xml:"shop"`
	Name       string     `xml:"shop"`
	Company    string     `xml:"company"`
	Url        string     `xml:"url"`
	Platform   string     `xml:"platform"`
	Currencies Currencies `xml:"currencies"`
	Categories Categories `xml:"categories"`
	Offers     Offers     `xml:"offers"`
}

type Currencies struct {
	Currency Currency  `xml:"currency"`
	Valute   []Valute  `xml:"Valute"`
	Date     time.Time `xml:"Date,attr"`
	Name     string    `xml:"name,attr"`
}

type Currency struct {
	Id       string `xml:"id,attr"`
	ID       string `xml:"ID,attr"`
	Currency string
	Code     string  `xml:"CharCode"`
	Rate     float64 `xml:"rate,attr"`
	Name     string  `xml:"Name"`
	Value    float64 `xml:"Value"`
	UnitRate float64 `xml:"VunitRate"`
	Nominal  float64 `xml:"Nominal"`
}

type Valute struct {
	Id       string  `xml:"ID,attr"`
	Code     string  `xml:"CharCode"`
	Name     string  `xml:"Name"`
	Value    float64 `xml:"Value"`
	UnitRate float64 `xml:"VunitRate"`
	Nominal  float64 `xml:"Nominal"`
}
