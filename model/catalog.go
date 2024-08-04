package model

type Catalog struct {
	XMLName struct{} `xml:"yml_catalog"`
	Date    string   `xml:"date,attr"`
	Shop    Shop     `xml:"shop"`
}
