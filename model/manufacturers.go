package model

import "time"

type Manufacturer struct {
	XMLName          struct{}          `xml:"manufacturer"`
	Id               int               `xml:"id"`
	Active           int               `xml:"active"`
	Name             string            `xml:"name"`
	DateAdd          time.Time         `xml:"date_add"`
	DateUpd          time.Time         `xml:"time.Time"`
	Description      *Description      `xml:"description"`
	DescriptionShort *DescriptionShort `xml:"short_description"`
	MetaTitle        *MetaTitle        `xml:"meta_title"`
	MetaDescription  *MetaDescription  `xml:"meta_description"`
	MetaKeywords     *MetaKeywords     `xml:"meta_keywords"`
	Associations     *Associations     `xml:"associations"`
}
