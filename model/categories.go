package model

import (
	"encoding/xml"
	"time"
)

type Categories struct {
	XMLName  xml.Name   `xml:"categories"`
	Category []Category `xml:"category"`
}

type Category struct {
	XMLName         xml.Name         `xml:"category"`
	Value           string           `xml:",chardata"`
	Id              int              `xml:"id,attr"`
	ParentId        int              `xml:"id_parent"`
	AparentId       int              `xml:"parentId,attr"`
	Active          int              `xml:"active"`
	DefaultShopId   int              `xml:"id_shop_default"`
	IsRootCategory  int              `xml:"is_root_category"`
	Position        int              `xml:"position"`
	DateAdd         time.Time        `xml:"date_add"`
	DateUpd         time.Time        `xml:"date_upd"`
	Name            *Name            `xml:"name"`
	LinkRewrite     *LinkRewrite     `xml:"link_rewrite"`
	Description     *Description     `xml:"description"`
	MetaTitle       *MetaTitle       `xml:"meta_title"`
	MetaDescription *MetaDescription `xml:"meta_description"`
	MetaKeywords    *MetaKeywords    `xml:"meta_keywords"`
	Associations    *Associations    `xml:"associations"`
}

type OzonCategory struct {
	Id       int64          `json:"category_id"`
	Name     string         `json:"category_name"`
	Disabled bool           `json:"disabled"`
	TypeName string         `json:"type_name"`
	TypeId   int64          `json:"type_id"`
	Children []OzonCategory `json:"children"`
}
