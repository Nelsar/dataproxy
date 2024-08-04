package model

type Tags struct {
	XMLName struct{} `xml:"tags"`
	Tag     *Tag     `xml:"tag"`
}

type Tag struct {
	XMLName struct{} `xml:"tag"`
	Id      int      `xml:"id"`
}
