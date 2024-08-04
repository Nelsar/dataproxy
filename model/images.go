package model

type Images struct {
	XMLName struct{} `xml:"images"`
	Image   *Image   `xml:"image"`
}

type Image struct {
	XMLName struct{} `xml:"image"`
	Id      int      `xml:"id"`
}
