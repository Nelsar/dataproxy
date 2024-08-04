package model

type Attachments struct {
	XMLName    struct{}    `xml:"attachments"`
	Attachment *Attachment `xml:"attachment"`
}

type Attachment struct {
	XMLName struct{} `xml:"attachment"`
	Id      int      `xml:"id"`
}
